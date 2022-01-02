// Package redis 工具包
package redis

import (
	"context"
	"errors"
	"gohub/pkg/logger"
	"sync"
	"time"

	redis "github.com/go-redis/redis/v8"
)

// RedisClient Redis 服务
type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// once 确保全局的 Redis 对象只实例一次
var once sync.Once

// Redis 全局 Redis，使用 db 1
var Redis *RedisClient

// ConnectRedis 连接 redis 数据库，设置全局的 Redis 对象
func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

// NewClient 创建一个新的 redis 连接
func NewClient(address string, password string, username string, db int) *RedisClient {

	// 初始化自定的 RedisClient 实例
	rds := &RedisClient{}
	// 使用默认的 context
	rds.Context = context.Background()

	// 使用 redis 库里的 NewClient 初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	// 测试一下连接
	err := rds.Ping()
	logger.LogIf(err)

	return rds
}

// Ping 用以测试 redis 连接是否正常
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return rds.Client.Set(rds.Context, key, value, expiration).Err()
}

// Get 获取 key 对应的 value
func (rds RedisClient) Get(key string) (string, error) {
	return rds.Client.Get(rds.Context, key).Result()
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds RedisClient) Del(keys ...string) error {
	return rds.Client.Del(rds.Context, keys...).Err()
}

// FlushDB 清空当前 redis db 里的所有数据
func (rds RedisClient) FlushDB(keys ...string) error {
	return rds.Client.FlushDB(rds.Context).Err()
}

// Increment 当参数只有 1 个时，为 key，其值增加 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要增加的值 int64 类型。
func (rds RedisClient) Increment(parameters ...interface{}) error {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		return rds.Client.Incr(rds.Context, key).Err()
	case 2:
		key := parameters[0].(string)
		value := parameters[0].(int64)
		return rds.Client.IncrBy(rds.Context, key, value).Err()
	default:
		return errors.New("请提供正确的参数")
	}
}

// Decrement 当参数只有 1 个时，为 key，其值减去 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要减去的值 int64 类型。
func (rds RedisClient) Decrement(parameters ...interface{}) error {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		return rds.Client.Decr(rds.Context, key).Err()
	case 2:
		key := parameters[0].(string)
		value := parameters[0].(int64)
		return rds.Client.DecrBy(rds.Context, key, value).Err()
	default:
		return errors.New("请提供正确的参数")
	}
}
