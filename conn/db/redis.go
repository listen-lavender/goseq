package db

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

func NewPool(network string,
	dsn string,
	db int,
	password string,
	readTimeout time.Duration,
	writeTimeout time.Duration,
	connectTimeout time.Duration,
	maxIdle int,
	maxActive int,
	idleTimeout time.Duration) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(network,
				dsn,
				redis.DialDatabase(db),
				redis.DialConnectTimeout(connectTimeout),
				redis.DialReadTimeout(readTimeout),
				redis.DialWriteTimeout(writeTimeout),
				redis.DialPassword(password),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}
