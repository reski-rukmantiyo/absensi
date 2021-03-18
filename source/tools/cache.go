package tools

import (
	"absensi/source/config"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Cache struct {
	config *config.Config
	client *redis.Client
}

func NewCache() *Cache {
	config := config.NewConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.CacheServer,
		Password: "",             // no password set
		DB:       config.CacheDB, // use default DB
	})
	cache := &Cache{
		config: config,
		client: rdb,
	}
	return cache
}

func (cache *Cache) Get(key string) (string, error) {
	val, err := cache.client.Get(ctx, key).Result()
	if err != nil {
		err := fmt.Errorf("empty result for cache %s:%s", key, err.Error())
		return "", err
	}
	return val, nil
}

func (cache *Cache) Set(key, value string, params ...int) error {
	err := cache.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		err := fmt.Errorf("Cache Set %s-%s:%s", key, value, err.Error())
		return err
	}
	return nil
}
