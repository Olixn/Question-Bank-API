/**
 * @Author: Ne-21
 * @Description:
 * @File:  redis
 * @Version: 1.0.0
 * @Date: 2022/1/13 12:21
 */

package database

import (
	"github.com/Olixn/Question-Bank-API/config"
	"github.com/gomodule/redigo/redis"
	"time"
)

var AppKvDbPool *redis.Pool

func InitRedisPool() {
	var AppConfig = config.AppConfig
	AppKvDbPool = &redis.Pool{
		MaxIdle:     AppConfig.Redis.MaxIdle,
		MaxActive:   AppConfig.Redis.MaxActive,
		IdleTimeout: time.Second * time.Duration(AppConfig.Redis.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", AppConfig.Redis.Address+":"+AppConfig.Redis.Port)
		},
	}
}
