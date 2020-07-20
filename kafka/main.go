package kafka

import (

	"context"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"go.elastic.co/apm"
	"go.uber.org/fx"
)

// Producer for produce msg to kafka
type Producer interface {
	SendMsg(context.Context, sarama.Encoder) (int32, int64, error)
}

type producer struct {
	sarama.SyncProducer
	topic string
}

// NewProducer new a Kafka Producer
func NewProducer(lc fx.Lifecycle, c *conf.Kafka) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Flush.Frequency = 100 * time.Millisecond
	config.Producer.Return.Successes = true

	p, err := sarama.NewSyncProducer(c.Broker, config)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return p.Close()
		},
	})

	return &producer{p, c.Topic}, nil
}

// NewProducerWithEnv factory methord of Producer with env struct
func NewProducerWithEnv(lc fx.Lifecycle, env *conf.Env) (Producer, error) {
	return NewProducer(lc, env.Kafka)
}

// NewLogProducerWithEnv factory methord of Producer with env struct use log topic
func NewLogProducerWithEnv(lc fx.Lifecycle, env *conf.Env) (Producer, error) {
	return NewProducer(lc, &conf.Kafka{
		Broker: env.Kafka.Broker,
		Topic:  env.Kafka.TopicLog,
	})
}

// SendMsg2Kafka send message to kafka
func (p *producer) SendMsg(ctx context.Context, message sarama.Encoder) (int32, int64, error) {
	span, ctx := apm.StartSpan(ctx, "SendMessage", "queue.kafka.produce")
	defer span.End()
	return p.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Value: message,
	})
}

// MessageHandleFunc for consumer handle message
type MessageHandleFunc func(*sarama.ConsumerMessage) error

// Consumer kafka consumer
type Consumer struct {
	ready         chan bool
	broker        []string
	topic         string
	consumerGroup string
	handleFunc    MessageHandleFunc
}

// NewConsumer factory of consumer
func NewConsumer(c *conf.Kafka, h MessageHandleFunc) *Consumer {
	return &Consumer{
		ready:         make(chan bool),
		broker:        c.Broker,
		topic:         c.Topic,
		consumerGroup: c.ConsumerGroup,
		handleFunc:    h,
	}
}

// NewConsumerWithEnv factory of consumer with env
func NewConsumerWithEnv(c *conf.Env, h MessageHandleFunc) *Consumer {
	return NewConsumer(c.Kafka, h)
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	defer func() {
		if v := recover(); v != nil {
			e := apm.DefaultTracer.Recovered(v)
			e.Send()
		}
	}()
	for message := range claim.Messages() {
		if err := c.handleFunc(message); err != nil {
			continue
		}
		session.MarkMessage(message, "success")
	}
	return nil
}

// StartConsumerGroup consumer runtime
func StartConsumerGroup(lc fx.Lifecycle, c *Consumer) error {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(c.broker, c.consumerGroup, config)
	if err != nil {
		cancel()
		return err
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				defer wg.Done()
				for {
					if err := client.Consume(ctx, []string{c.topic}, c); err != nil {
						panic(err)
					}
					if ctx.Err() != nil {
						return
					}
					c.ready = make(chan bool)
				}
			}()
			<-c.ready
			return nil
		},
		OnStop: func(context.Context) error {
			cancel()
			wg.Wait()
			return client.Close()
		},
	})
	return nil
}
