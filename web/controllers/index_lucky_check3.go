package controllers

import (
	"fmt"
	"github.com/xuanxiaox/lottery/conf"
	"github.com/xuanxiaox/lottery/models"
	"log"
	"strconv"
	"time"
)

func (c *IndexController) CheckUserDay(uid int) bool {
	var userDay *models.LtUserday
	userDay = c.ServiceUserday.GetUserToday(uid)
	//用户已经参与过
	if userDay != nil && userDay.Uid == uid {
		if userDay.Num >= conf.UserPrizeMax {
			return false
		}
		userDay.Num++
		err103 := c.ServiceUserday.Update(userDay, nil)
		if err103 != nil {
			log.Println("index_lucky_check3.go CheckUserDay ServiceUserday.Update err:", err103)
		}
	} else {
		y, m, d := time.Now().Date()
		dayStr := fmt.Sprintf("%d%02d%02d", y, m, d)
		day, _ := strconv.Atoi(dayStr)
		userDay = &models.LtUserday{
			Uid:        uid,
			Day:        day,
			Num:        1,
			SysCreated: int(time.Now().Unix()),
		}
		err103 := c.ServiceUserday.Create(userDay)
		if err103 != nil {
			log.Println("index_lucky_check3.go CheckUserDay ServiceUserday.Create err:", err103)
		}
	}
	return true
}
