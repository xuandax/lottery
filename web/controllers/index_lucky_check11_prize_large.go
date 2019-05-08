package controllers

import (
	"github.com/xuanxiaox/lottery/comm"
	"github.com/xuanxiaox/lottery/models"
)

func (c *IndexController) PrizeLarge(loginUser *models.ObjLoginuser,
	userInfo *models.LtUser, blackIpInfo *models.LtBlackip) {
	now := comm.NowUnix()
	blackTime := 7 * 86400
	if userInfo == nil || userInfo.Id < 0 {
		userInfo = &models.LtUser{
			Id:         loginUser.Uid,
			Username:   loginUser.Username,
			BlackTime:  now + blackTime,
			SysCreated: now,
			SysIp:      loginUser.Ip,
		}
		c.ServiceUser.Create(userInfo)
	} else {
		userInfo.BlackTime = now + blackTime
		userInfo.SysUpdated = now
		userInfo.SysIp = loginUser.Ip
		c.ServiceUser.Update(userInfo, nil)
	}

	if blackIpInfo == nil || blackIpInfo.Id < 0 {
		blackIpInfo = &models.LtBlackip{
			Ip:         loginUser.Ip,
			BlackTime:  now + blackTime,
			SysCreated: now,
		}
		c.ServiceBlack.Create(blackIpInfo)
	} else {
		blackIpInfo.Ip = loginUser.Ip
		blackIpInfo.BlackTime = now + blackTime
		blackIpInfo.SysUpdated = now
		c.ServiceBlack.Update(blackIpInfo, nil)
	}
}
