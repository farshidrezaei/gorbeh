package redis

import (
	"context"
	"farshidrezaei/vms_mongo_log/config"
	"farshidrezaei/vms_mongo_log/library/utils"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var ctx = context.Background()

func newConnection() (context.Context, *redis.Client) {
	host := config.Database("connections.redis.host", nil)
	port := config.Database("connections.redis.port", nil)
	password := config.Database("connections.redis.password", nil)
	db := config.Database("connections.redis.db", nil)
	dbNumber, _ := strconv.Atoi(fmt.Sprintf("%s", db))
	redisConnection := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: fmt.Sprintf("%s", password), // no password set
		DB:       dbNumber,                    // use default DB
	})

	return ctx, redisConnection

}

func Set(key string, value string, ttl int) string {
	ctx, connection := newConnection()
	result, err := connection.Set(ctx, key, value, time.Duration(ttl*int(time.Second))).Result()
	if err != nil {
		utils.Dd(err)
	}
	return result
}

func Get(key string) (string, bool) {
	ctx, connection := newConnection()
	result, err := connection.Get(ctx, key).Result()
	if err != nil {
		return "", false
	}
	return result, true
}
