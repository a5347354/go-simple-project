package main

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

func main() {
	c := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    strings.Split("IP:PORT,IP:PORT,IP:PORT", ","),
		Password: "PASSWORD",
	})

	if err := c.Ping().Err(); err != nil {
		panic(err)
	}
	// LPush
	key := "KEY"
	snList := make([]string, 0)
	for i := 0; i < 1; i++ {
		snList = append(snList, fmt.Sprintf("gcard_sn_%d", i))
	}
	c.LPush(key, snList)

	// Check Redeem Code is sold out or not
	fmt.Println(c.LRange("12", 0, -1).Result())

}
