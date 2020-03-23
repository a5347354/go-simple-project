package helper

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test_bsonToStringOfJSON(t *testing.T) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	collection := client.Database("baz").Collection("qux")

	singleResult := collection.FindOne(ctx, bson.D{})

	str, err := bsonToStringOfJSON(singleResult)
	if err != nil {
		log.Fatal(err)
	}
	So(str, ShouldEqual, `{"abc":"abc"}`)
}
