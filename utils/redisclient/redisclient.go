package redisclient

import (
	"context"
	"fmt"

	"github.com/omniful/go_commons/redis"
	cache "github.com/omniful/go_commons/redis_cache"
)

var Redis_Client *redis.Client
var Redis_Cache *cache.RedisCache

func Connect(config *redis.Config) (*redis.Client, error) {
	client := redis.NewClient(config)

	// Check connection
	if err := client.Ping(context.TODO()).Err(); err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return nil, err
	}

	Redis_Client = client
	fmt.Println("Connected to Redis successfully!")

	return client, nil
}

// func InitCache(client *redis.Client, nameSpace string) {
// 	Redis_Cache = cache.NewRedisCacheClient(client, &cache.JSONSerializer{}, nameSpace)
// }
