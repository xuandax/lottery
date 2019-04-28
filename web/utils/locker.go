package utils

import (
	"fmt"
	"github.com/xuanxiaox/lottery/datasource"
)

func LockUserLucky(uid int) bool {
	key := getLuckUserLuckyKey(uid)
	cacheObj := datasource.InstanceRedisConn()
	rs, _ := cacheObj.Do("SET", key, 1, "EX", 3, "NX")
	if rs == "OK" {
		return true
	}
	return false
}

func UnlockUserLucky(uid int) bool {
	key := getLuckUserLuckyKey(uid)
	cacheObj := datasource.InstanceRedisConn()
	rs, _ := cacheObj.Do("DEL", key)
	if rs == "OK" {
		return true
	}
	return false
}

func getLuckUserLuckyKey(uid int) string {
	return fmt.Sprintf("luck_user_lucky_%d", uid)
}
