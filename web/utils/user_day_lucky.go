package utils

import (
	"fmt"
	"github.com/xuanxiaox/lottery/comm"
	"github.com/xuanxiaox/lottery/datasource"
	"log"
	"math"
	"time"
)

const UserLimitNum = 2

func init() {
	resetGroupUserLuckyList()
}

func resetGroupUserLuckyList() {
	log.Println("user_day_lucky.go resetGroupUserLuckyList start")
	cache := datasource.InstanceRedisConn()
	for i := 0; i < UserLimitNum; i++ {
		cache.Do("DEL", fmt.Sprintf("user_day_lucky_%d", i))
	}
	log.Println("user_day_lucky.go resetGroupUserLuckyList end")
	duration := comm.NextDayDuration()
	time.AfterFunc(duration, resetGroupUserLuckyList)
}

func IncrUserLuckyNum(uid int) int64 {
	sNum := uid % UserLimitNum
	key := fmt.Sprintf("user_day_lucky_%d", sNum)
	cache := datasource.InstanceRedisConn()
	rs, err := cache.Do("HINCRBY", key, uid, 1)
	if err != nil {
		log.Panicln("user_day_lucky.go IncrUserLuckyNum HINCRBY key=", key, ",uid=", uid, ", err =", err)
		return math.MaxInt64
	}
	return rs.(int64)
}

func InitUserLuckyNum(uid int, num int) {
	if num < 1 {
		return
	}
	sNum := uid % UserLimitNum
	key := fmt.Sprintf("user_day_lucky_%d", sNum)
	cache := datasource.InstanceRedisConn()
	_, err := cache.Do("HSET", key, uid, num)
	if err != nil {
		log.Panicln("user_day_lucky.go InitUserLuckyNum HSET key=", key, ",uid=", uid, ", err =", err)
	}
}
