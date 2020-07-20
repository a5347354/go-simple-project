package elasticsearch

import(
	"github.com/elastic/go-elasticsearch/v6/elasticsearch"

	"net/http"
	"time"
)

func NewClient(host []string){
	cfg := elasticsearch.Config{
		Addresses: host,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
			  MinVersion: tls.VersionTLS11,
			},
		  },
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil{
		panic(err)
	}
}