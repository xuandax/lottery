package controllers

import "imooc.com/lottery/comm"

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

	//3 验证用户今日参与次数

	//4 验证IP今日参与次数

	//5 验证IP黑名单

	//6 验证用户黑名单

	//7 获取抽奖编码

	//8 匹配是否中奖

	//9 有限制的发放奖品

	//10 不同编码的优惠券发放

	//11 记录中奖记录

	//12 返回抽奖结果

	return rs
}
