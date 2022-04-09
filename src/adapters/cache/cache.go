package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack"
)

type CacheAdapter interface {
	GetInt64(ctx context.Context, key string) (int64, error)
	Get(ctx context.Context, key string, v interface{}) error
	Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HGet(ctx context.Context, key string, field string) (string, error)
	HSet(ctx context.Context, key string, field string, value interface{}) error
	HIncrBy(ctx context.Context, key string, field string, incr int64) (int64, error)
}

type redisAdapter struct {
	client *redis.Client
}

func (a *redisAdapter) Get(ctx context.Context, key string, v interface{}) error {
	data, err := a.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	return msgpack.Unmarshal(data, v)
}

func (a *redisAdapter) Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error {
	data, err := msgpack.Marshal(v)
	if err != nil {
		return nil
	}

	_, err = a.client.Set(ctx, key, data, expiration).Result()
	return err
}

func (r *redisAdapter) GetInt64(ctx context.Context, key string) (int64, error) {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return 0, err
	}
	var rs int64
	msgpack.Unmarshal(data, &rs)
	return rs, nil
}

func (a *redisAdapter) HSet(ctx context.Context, key string, field string, v interface{}) error {
	data, err := msgpack.Marshal(v)
	if err != nil {
		return nil
	}

	_, err = a.client.HSet(ctx, key, field, string(data)).Result()
	return err
}

func (r *redisAdapter) HGet(ctx context.Context, key string, field string) (string, error) {
	return r.client.HGet(ctx, key, field).Result()
}

func (r *redisAdapter) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.client.HGetAll(ctx, key).Result()
}

func (a *redisAdapter) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	return a.client.HIncrBy(ctx, key, field, incr).Result()
}

func NewRedisAdapter(client *redis.Client) CacheAdapter {
	return &redisAdapter{
		client: client,
	}
}
