package redis_client

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"pop-api/baselib/logger"
	"time"
)

type RedisPool struct {
	*redis.Pool
	env string
}

var redisPoolMap = make(map[string]*RedisPool)

func GetRedisPoolClient(redisName string) (conn redis.Conn) {
	pool, ok := redisPoolMap[redisName]
	if !ok {
		logger.LogSugar.Errorf("GetRedisClient - Not found client: %s", redisName)
	}
	conn = pool.Pool.Get()
	return
}

func NewRedisPool(c *RedisConfig) (pool *RedisPool) {
	pool = &RedisPool{env: fmt.Sprintf("[%s]tcp@%s", c.Name, c.Addr)}
	cnop := redis.DialConnectTimeout(time.Duration(c.DialTimeout))
	rdop := redis.DialReadTimeout(time.Duration(c.ReadTimeout))
	wrop := redis.DialWriteTimeout(time.Duration(c.WriteTimeout))

	dialFunc := func() (rconn redis.Conn, err error) {
		rconn, err = redis.Dial("tcp", c.Addr, cnop, rdop, wrop)
		if err != nil {
			logger.LogSugar.Errorf("Redis connect %s error: %v", pool.env, err)
			return
		}

		if c.Password != "" {
			if _, err = rconn.Do("AUTH", c.Password); err != nil {
				logger.LogSugar.Errorf("Redis %s AUTH(password: %s) error: %v", pool.env, c.Password, err)
				rconn.Close()
				rconn = nil
				return
			}
		}

		// TODO(@work):  检查c.DBNum，必须是数字
		_, err = rconn.Do("SELECT", c.DBNum)
		if err != nil {
			logger.LogSugar.Errorf("Redis %s SELECT %s error: %v", pool.env, c.DBNum, err)
			rconn.Close()
			rconn = nil
		}
		return
	}

	pool.Pool = &redis.Pool{
		MaxActive:   c.Active,
		MaxIdle:     c.Idle,
		IdleTimeout: time.Duration(c.IdleTimeout),
		Dial:        dialFunc,
		Wait:        true,
	}
	return
}
