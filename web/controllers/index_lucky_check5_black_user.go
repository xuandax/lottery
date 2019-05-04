package controllers

import (
	"github.com/xuanxiaox/lottery/models"
	"time"
)

func (c *IndexController) CheckBlackUser(id int) (bool, *models.LtUser) {
	userInfo := c.ServiceUser.Get(id)
	if userInfo != nil && userInfo.BlackTime > int(time.Now().Unix()) {
		return true, userInfo
	}
	return false, nil
}
