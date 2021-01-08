package util

import (
	"encoding/json"
	"fmt"
)

func NestedStruct() {
	tests := []struct {
		Host   string `json:"host"`
		Insert bool   `json:"insert"`
		Update bool   `json:"update"`
		Sync   []struct {
			ConnectionID string
		}
	}{
		{
			Host:   "192.0.2.0:8000",
			Insert: true,
			Update: true,
			Sync: []struct {
				ConnectionID string
			}{
				{
					ConnectionID: "192.0.2.0:8000",
				},
			},
		},
		{
			Host:   "192.0.2.0:80",
			Insert: true,
			Update: true,
			Sync: []struct {
				ConnectionID string
			}{
				{
					ConnectionID: "192.0.2.0:80",
				},
			},
		},
	}
	for _, tt := range tests {
		fmt.Println(tt.Host)
		foo_marshalled, _ := json.Marshal(tt)
		fmt.Println(string(foo_marshalled))
	}
}
