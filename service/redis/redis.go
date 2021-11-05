package redis

import (
	"context"
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/utils"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     utils.SingleBackInt(conf.RedisInfo.IP+":", conf.RedisInfo.Port),
		Password: conf.RedisInfo.Password,
		DB:       15,
		PoolSize: 100,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func SetAdd(key, value string) (err error) {
	ctx := context.Background()
	if initErr := initClient(); err != nil {
		return initErr
	}

	addErr := rdb.SAdd(ctx, key, value).Err()
	if addErr != nil {
		return addErr
	} else {
		return nil
	}
}

func SetDelete(key, member string) (err error) {

	ctx := context.Background()
	if err = initClient(); err != nil {
		return err
	}

	data, err := rdb.SMembers(ctx, key).Result()

	for _, item := range data {
		if strings.Contains(item, member) {
			member = item
			break
		}
	}

	return rdb.SRem(ctx, key, member).Err()
}

func SetGet(key string) (data []string, err error) {
	ctx := context.Background()
	if err = initClient(); err != nil {
		return nil, err
	}

	data, err = rdb.SMembers(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
