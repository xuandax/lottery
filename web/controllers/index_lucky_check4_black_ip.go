package controllers

import (
	"github.com/xuanxiaox/lottery/models"
	"time"
)

func (c *IndexController) CheckBlackIp(ip string) (bool, *models.LtBlackip) {
	blackIpInfo := c.ServiceBlack.GetByIp(ip)
	if blackIpInfo == nil || blackIpInfo.Ip == "" {
		return false, nil
	}

	if blackIpInfo.BlackTime > int(time.Now().Unix()) {
		return true, blackIpInfo
	}
}
