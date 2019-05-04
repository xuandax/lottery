package utils

import (
	"fmt"
	"github.com/xuanxiaox/lottery/comm"
	"github.com/xuanxiaox/lottery/datasource"
	"log"
	"math"
	"time"
)

const SplitNum = 2

func init() {
	resetGroupList()
}

//清除IP缓存
func resetGroupList() {
	log.Println("ipDayLucky resetGroupList start")
	for i := 0; i < SplitNum; i++ {
		key := fmt.Sprintf("ip_day_lucky_%d", i)
		cacheObj := datasource.InstanceRedisConn()
		cacheObj.Do("DEL", key)
	}
	log.Println("ipDayLucky resetGroupList end")
	//在零点的时候，定时清理缓存
	duration := comm.NextDayDuration()
	time.AfterFunc(duration, resetGroupList)
}

func CheckIpDayLucky(ipStr string) int64 {
	ip := comm.Ip4toInt(ipStr)
	index := ip % SplitNum
	cacheObj := datasource.InstanceRedisConn()
	key := fmt.Sprintf("ip_day_lucky_%d", index)
	rs, err := cacheObj.Do("HINCRBY", key, ip, 1)
	if err != nil {
		log.Println("CheckIpDayLucky HINCRBY err = ", err)
		return math.MaxInt64
	}
	return rs.(int64)
}
