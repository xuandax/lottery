package datasource

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/xuanxiaox/lottery/conf"
	"log"
	"sync"
	"time"
)

var instanceRedis *RedisConn
var redisMux sync.Mutex

type RedisConn struct {
	pool      *redis.Pool
	showDebug bool
}

func (rds *RedisConn) Do(connamdName string, args ...interface{}) (reply interface{}, err error) {
	conn := rds.pool.Get()
	defer conn.Close()

	t1 := time.Now().UnixNano()
	reply, err = conn.Do(connamdName, args...)
	if err != nil {
		e := conn.Err()
		if e != nil {
			log.Println("redisHelper Do err:", err, e)
		}
	}
	t2 := time.Now().UnixNano()
	if rds.showDebug {
		fmt.Printf("[redis] [info] [time:%dus]cmd=%s,err=%s,args=%s,reply=%s\n", (t2-t1)/1000,
			connamdName, err, args, reply)
	}
	return reply, err
}

func (rds *RedisConn) ShowDebug(b bool) {
	rds.showDebug = b
}

func InstanceRedisConn() *RedisConn {
	if instanceRedis != nil {
		return instanceRedis
	}
	redisMux.Lock()
	defer redisMux.Unlock()

	if instanceRedis != nil {
		return instanceRedis
	}
	return newRedisConn()
}

//创建RedisConn实例
func newRedisConn() *RedisConn {
	//创建redis连接池
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%d", conf.RedisCache.Host, conf.RedisCache.Port))
			if err != nil {
				log.Fatal("redisHelper newRedisConn err = ", err)
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
		MaxIdle:         10000, //最大连接数
		MaxActive:       10000, //最大活跃数
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}
	instanceRedis = &RedisConn{
		pool:      &pool,
		showDebug: false,
	}
	instanceRedis.ShowDebug(true)
	return instanceRedis
}
