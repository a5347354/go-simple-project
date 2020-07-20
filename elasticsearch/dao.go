package elasticsearch

import(
  "github.com/elastic/go-elasticsearch/v6/esapi"
  "github.com/elastic/go-elasticsearch/v6"

  "bytes"
  "context"
  "encoding/json"
  "log"
  "strconv"
  "strings"
  "sync"
)

type Dao interface{
  Insert()
}

func Insert(client ){
    req := esapi.IndexRequest{
          Index:      "test",
          DocumentID: strconv.Itoa(i + 1),
          Body:       strings.NewReader(b.String()),
          Refresh:    "true",
        }

    // Perform the request with the client.
    res, err := req.Do(context.Background(), client)
    if err != nil {
      log.Fatalf("Error getting response: %s", err)
    }
    defer res.Body.Close()
}