package redis_client

import (
	"fmt"
	goredis "github.com/go-redis/redis/v7"
	"time"
)

var (
	RedisCache *goredis.Client // bot 先存入 cache中
)

type Duration time.Duration

type RedisConfig struct {
	Name         string // redis name
	Addr         string
	Active       int // pool
	Idle         int // pool
	DialTimeout  Duration
	ReadTimeout  Duration
	WriteTimeout Duration
	IdleTimeout  Duration

	DBNum    int    // db号
	Password string // 密码
}

func InstallRedisClientManager(configs RedisConfig) {
	pool := NewRedis(&configs)
	if pool == nil {
		err := fmt.Errorf("InstallRedisClient - NewRedis {%v} error!", configs)
		panic(err)
	}
	pool_redigo := NewRedisPool(&configs)
	redisPoolMap[configs.Name] = pool_redigo
	RedisCache = pool
}

func NewRedis(c *RedisConfig) (client *goredis.Client) {
	client = goredis.NewClient(&goredis.Options{
		Addr:         c.Addr,
		Password:     c.Password, // no password set
		DB:           c.DBNum,    // use default DB	c.DBNum
		DialTimeout:  time.Duration(c.DialTimeout),
		ReadTimeout:  time.Duration(c.ReadTimeout),
		WriteTimeout: time.Duration(c.WriteTimeout),
		PoolSize:     c.Active,                     // 连接池最大socket连接数, 默认为10倍CPU数， 10 * runtime.NumCPU
		MinIdleConns: c.Idle,                       // 在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量
		PoolTimeout:  time.Duration(c.IdleTimeout), // 当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
	})
	return
}
