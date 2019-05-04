package controllers

import (
	"github.com/xuanxiaox/lottery/comm"
	"github.com/xuanxiaox/lottery/conf"
	"github.com/xuanxiaox/lottery/models"
	"github.com/xuanxiaox/lottery/web/utils"
	"log"
)

func (c *IndexController) GetLucky() map[string]interface{} {
	rs := make(map[string]interface{}, 0)
	rs["code"] = 0
	rs["msg"] = ""

	//1 验证用户是否登录
	loginUser := comm.GetLoginUser(c.Ctx.Request())
	if loginUser == nil || loginUser.Uid < 1 {
		rs["code"] = 101
		rs["msg"] = "请先登录，再参与抽奖"
		return rs
	}

	//2 设置用户抽奖分布式锁
	ok := utils.LockUserLucky(loginUser.Uid)
	if ok {
		defer utils.UnlockUserLucky(loginUser.Uid)
	} else {
		rs["code"] = 102
		rs["msg"] = "正在抽奖，请勿重复点击"
		return rs
	}
	//3 验证用户今日参与次数
	ok = c.CheckUserDay(loginUser.Uid)
	if !ok {
		rs["code"] = 103
		rs["msg"] = "今日抽奖次数已用完，请明日再来"
		return rs
	}
	//4 验证IP今日参与次数
	ip := comm.ClientIP(c.Ctx.Request())
	num := utils.CheckIpDayLucky(ip)
	if num > conf.LimitIpMaxNum {
		rs["code"] = 103
		rs["msg"] = "今日抽奖次数已用完，请明日再来"
		return rs
	}
	//是否在黑名单
	limitBlack := false
	if num > conf.IpPrizeMax {
		limitBlack = true
	}
	//5 验证IP黑名单
	var blackIp *models.LtBlackip
	if !limitBlack {
		ok, blackIp = c.CheckBlackIp(ip)
		if ok {
			log.Println("黑名单中的ip：", ip, limitBlack)
			limitBlack = true
		}
	}
	//6 验证用户黑名单
	var blackUser *models.LtUser
	if !limitBlack {
		ok, blackUser = c.CheckBlackUser(loginUser.Uid)
		if ok {
			log.Println("黑名单中的用户：", loginUser.Uid, limitBlack)
			limitBlack = true
		}
	}
	//7 获取抽奖编码
	prizeCode := comm.Random(10000)
	//8 匹配是否中奖
	giftPrize := c.Prize(prizeCode, limitBlack)
	if giftPrize != nil || giftPrize.PrizeNum <= 0 || (giftPrize.PrizeNum > 0 && giftPrize.LeftNum <= 0) {
		rs["code"] = 205
		rs["msg"] = "很遗憾，没有中奖，请下次再来"
		return rs
	}
	//9 有限制的发放奖品
	if giftPrize.PrizeNum > 0 {
		ok := utils.PrizeGift(giftPrize.Id, giftPrize.LeftNum)
		if !ok {
			rs["code"] = 207
			rs["msg"] = "很遗憾，没有中奖，请下次再来"
			return rs
		}
	}
	//10 不同编码的优惠券发放
	if giftPrize.Gtype == conf.GtypeCodeDiff {
		code := utils.PrizeCodeDiff(giftPrize.Id, c.ServiceCode)
		if code == "" {
			rs["code"] = 208
			rs["msg"] = "很遗憾，没有中奖，请下次再来"
			return rs
		}
		giftPrize.Gdata = code
	}
	//11 记录中奖记录
	result := models.LtResult{
		GiftId:     giftPrize.Id,
		GiftName:   giftPrize.Title,
		GiftType:   giftPrize.Gtype,
		Uid:        loginUser.Uid,
		Username:   loginUser.Username,
		PrizeCode:  prizeCode,
		GiftData:   giftPrize.Gdata,
		SysStatus:  0,
		SysCreated: comm.NowUnix(),
		SysIp:      ip,
	}
	err := c.ServiceResult.Create(&result)
	if err != nil {
		log.Println("index_lucky.go GetLucky c.ServiceResult.Create err=", err)
		rs["code"] = 209
		rs["msg"] = "很遗憾，没有中奖，请下次再来"
		return rs
	}
	//中了实物大奖需要关小黑屋,将用户和IP设置为黑名单一段时间
	if giftPrize.Gtype == conf.GtypeGiftLarge {

	}
	//12 返回抽奖结果
	rs["gift"] = giftPrize
	return rs
}
