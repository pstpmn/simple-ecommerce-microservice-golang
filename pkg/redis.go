package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type IRedis interface {
	Cache(data any, key string, ttl time.Duration) error
	Find(key string) (map[string]interface{}, error)
	Delete(key string) error
}

type r struct {
	client *redis.Client
}

var ctx = context.Background()

func NewRedis(user string, pass string, host string, port string) (IRedis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Username: user,
		Password: pass,
		DB:       0,
	})
	_, err := client.Ping(ctx).Result()
	return &r{
		client: client,
	}, err
}

// cache implements IRedis.
func (r *r) Cache(data any, key string, ttl time.Duration) error {
	str, _ := json.Marshal(data)
	_, err := r.client.SetNX(ctx, key, string(str), time.Duration(ttl)).Result()
	return err
}

// delete implements IRedis.
func (r *r) Delete(key string) error {
	r.client.Del(ctx, key)
	return nil
}

// find implements IRedis.
func (r *r) Find(key string) (map[string]interface{}, error) {
	var cache, err = r.client.Get(ctx, key).Result()
	result := make(map[string]interface{})
	json.Unmarshal([]byte(cache), &result)
	errStr := fmt.Sprintf("%s", err)
	if errStr == "redis: nil" {
		err = nil
	}
	return result, err
}
