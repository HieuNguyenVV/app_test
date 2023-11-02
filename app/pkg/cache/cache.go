package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

const (
	Expired = 3600 * time.Second
)

type Caching struct {
	cache     *redis.Client
	batchSize int
}

func NewCaching(cache *redis.Client) *Caching {
	return &Caching{
		cache:     cache,
		batchSize: 500,
	}
}

func (c *Caching) Write(ctx context.Context, key string, value interface{}, args redis.SetArgs) error {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(value)
	if err != nil {
		return err
	}
	err = c.cache.SetArgs(context.Background(), key, buf.Bytes(), args).Err()
	if err != nil {
		return err
	}
	return c.cache.Expire(context.Background(), key, Expired).Err()
}

func (c *Caching) Read(ctx context.Context, key string) ([]byte, error) {
	return c.cache.Get(context.Background(), key).Bytes()
}

func (c *Caching) Delete(ctx context.Context, keys ...string) ([]string, error) {
	if len(keys) < c.batchSize {
		return nil, fmt.Errorf("max length of key is %v", c.batchSize)
	}

	pipe := c.cache.Pipeline()
	for _, key := range keys {
		pipe.Del(context.Background(), key)
		if _, err := pipe.Exec(context.Background()); err != nil {
			return keys, nil
		}
	}
	if _, err := pipe.Exec(context.Background()); err != nil {
		return nil, err
	}
	return keys, nil
}
