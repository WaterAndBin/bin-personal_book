package data

import (
	"bin-personal-book/internal/biz"
	"bin-personal-book/internal/conf"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type RedisData struct {
	redisClient *redis.Client
	log         *log.Helper
}

func NewRedisData(redis *redis.Client, logger log.Logger) biz.RedisZip {
	return &RedisData{
		redisClient: redis,
		log:         log.NewHelper(logger),
	}
}

func NewRedisClient(c *conf.Bootstrap, logger log.Logger) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: c.Data.Redis.Addr,
	})

	// 测试连接
	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.NewHelper(logger).Fatal("redis连接失败")
	}

	log.NewHelper(logger).Info("redis连接成功")

	return rdb
}

// 设置 key-value
func (r *RedisData) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.redisClient.Set(
		ctx,
		key,
		value,
		expiration,
	).Err()
}

// 获取 value
func (r *RedisData) Get(ctx context.Context, key string) (string, error) {
	return r.redisClient.Get(
		ctx,
		key,
	).Result()
}

// 删除 key
func (r *RedisData) Delete(ctx context.Context, key string) error {
	return r.redisClient.Del(
		ctx,
		key,
	).Err()
}
