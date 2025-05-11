package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/fzr365/urlshortener/internal/repo"
	"github.com/go-redis/redis/v8"
)

//具体实现SetURL(ctx context.Context, url repo.Url) error接口

type RedisCache struct {
	client *redis.Client
}

//新建对外函数
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}	
}




func (c *RedisCache) SetURL(ctx context.Context, url repo.Url) error {
	//因为传入的是结构体，存储的是string，所以需要将结构体转换为string，序列化
	data, err := json.Marshal(url)
	if err != nil {
		return err
	}
	if err := c.client.Set(ctx, url.ShortCode, data, time.Until(url.ExpiredAt)).Err(); err != nil {
		return err
	}

	return nil
}

//实现GetURL(ctx context.Context, shortCode string) (*repo.Url, error)

func (c *RedisCache) GetURL(ctx context.Context, shortCode string) (*repo.Url, error) {
	data, err := c.client.Get(ctx, shortCode).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	//将data反序列化为结构体
	var url repo.Url
	if err := json.Unmarshal(data, &url); err != nil {
		return nil, err
	}

	return &url, nil
}
