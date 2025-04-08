package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strings"
)

func main() {
	cache, err := NewRedisCache()
	if err != nil {
		fmt.Printf("err: %+v", err)
	}

	cache.SetField("user:1", "name", "Alice")
	cache.SetField("user:1", "age", "11")
	cache.SetField("user:2", "name", "Alice22")
	cache.SetField("user:2", "age", "22")
	// 示例：获取哈希字段的值
	name, err := cache.GetField("user:1", "name")
	if err != nil {
		panic(err)
	}
	fmt.Printf("User name: %s\n", name)

	// 示例：删除哈希字段
	err = cache.DeleteField("user:1", "name")
	if err != nil {
		panic(err)
	}

	// 示例：检查字段是否存在
	exists, err := cache.FieldExists("user:1", "name")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Field exists: %v\n", exists)

	// 示例：获取所有字段和值
	fields, err := cache.GetAllFields("user:2")
	if err != nil {
		panic(err)
	}
	fmt.Printf("All fields: %v\n", fields)
}

func NewRedisCache() (*RedisCache, error) {
	redisHost := "172.22.175.230:7000,172.22.175.230:7001,172.22.175.230:7002"
	redisAuth := "MXFhenhzdzI="
	addrs := strings.Split(redisHost, ",")
	redisPwd, _ := base64.StdEncoding.DecodeString(redisAuth)
	redisPwdStr := strings.TrimSpace(string(redisPwd))

	redisClient := redis.NewClusterClient(
		&redis.ClusterOptions{
			Addrs:    addrs,
			Password: redisPwdStr,
		},
	)
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis cluster: %+v", err)
		return nil, err
	}

	return &RedisCache{
		ClusterClient: redisClient,
		//Client:     rdb,
	}, nil
}

type RedisCache struct {
	TenantId   string
	TenantCode string
	*redis.ClusterClient
	//*redis.Client
}

// SetField 设置哈希键 key 下的字段 field 为 value
func (rc *RedisCache) SetField(key, field, value string) error {
	_, err := rc.HSet(context.Background(), key, field, value).Result()
	return err
}

// GetField 获取哈希键 key 下字段 field 的值
func (rc *RedisCache) GetField(key, field string) (string, error) {
	value, err := rc.HGet(context.Background(), key, field).Result()
	return value, err
}

// DeleteField 删除哈希键 key 下的字段 field
func (rc *RedisCache) DeleteField(key, field string) error {
	_, err := rc.HDel(context.Background(), key, field).Result()
	return err
}

// FieldExists 检查哈希键 key 下的字段 field 是否存在
func (rc *RedisCache) FieldExists(key, field string) (bool, error) {
	exists, err := rc.HExists(context.Background(), key, field).Result()
	return exists, err
}

// GetAllFields 获取哈希键 key 下的所有字段及其值
func (rc *RedisCache) GetAllFields(key string) (map[string]string, error) {
	result, err := rc.HGetAll(context.Background(), key).Result()
	return result, err
}

// SetMultipleFields 批量设置哈希键 key 下的多个字段和值
func (rc *RedisCache) SetMultipleFields(key string, fields map[string]interface{}) error {
	_, err := rc.HMSet(context.Background(), key, fields).Result()
	return err
}

// GetMultipleFields 批量获取哈希键 key 下的多个字段的值
func (rc *RedisCache) GetMultipleFields(key string, fields []string) ([]interface{}, error) {
	values, err := rc.HMGet(context.Background(), key, fields...).Result()
	return values, err
}

// GetAllKeys 获取哈希键 key 下的所有字段
func (rc *RedisCache) GetAllKeys(key string) ([]string, error) {
	keys, err := rc.HKeys(context.Background(), key).Result()
	return keys, err
}

// GetAllValues 获取哈希键 key 下的所有值
func (rc *RedisCache) GetAllValues(key string) ([]string, error) {
	values, err := rc.HVals(context.Background(), key).Result()
	return values, err
}
