package cache
import("github.com/go-redis/redis/v8")

//具体实现SetURL(ctx context.Context, url repo.Url) error接口

type RedisCache struct {
	client *redis.Client
}


func (c *RedisCache) SetURL(ctx context.Context, url repo.Url) error {
	//因为传入的是结构体，存储的是string，所以需要将结构体转换为string，序列化
	data,err:= json.Marshal(url)
	if err!=nil {
		return err
	}
	if err:= c.client.Set(ctx, url.ShortCode, data, time.Until(url.ExpiredAt)).Err();err!=nil {
		return err
	}

	return nil
}