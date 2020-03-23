package helper

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func bsonToStringOfJSON(rawData *mongo.SingleResult) (string, error) {
	rawBytes, err := rawData.DecodeBytes()
	if err != nil {
		return "", err
	}
	var result map[string]interface{}
	bson.Unmarshal(rawBytes, &result)
	// Remove ID Field
	if _, exist := result["_id"]; exist {
		delete(result, "_id")
	}
	str, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(str), nil
}
