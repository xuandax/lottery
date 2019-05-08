package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/xuanxiaox/lottery/comm"
	"github.com/xuanxiaox/lottery/models"
	"github.com/xuanxiaox/lottery/services"
)

type IndexController struct {
	Ctx            iris.Context
	ServiceGift    services.GiftService
	ServiceBlack   services.BlackipService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserday services.UserdayService
	ServiceUser    services.UserService
}

//首页
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "欢迎来到抽奖页面 <a href='/public/page.html'>点击抽奖</a>"
}

//获取礼物列表
func (c *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	dataList := c.ServiceGift.GetAll(true)
	list := make([]models.LtGift, 0)
	for _, data := range dataList {
		list = append(list, data)
	}
	rs["gifts"] = list
	return rs
}

//获取最新的获奖列表
func (c *IndexController) GetNewPrize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	//todo

	return rs
}

//用户登录
func (c *IndexController) GetLogin() {
	uid := comm.Random(1000000)
	loginUser := &models.ObjLoginuser{
		Uid:      uid,
		Username: fmt.Sprintf("login_user_%d", uid),
		Now:      comm.NowUnix(),
		Ip:       comm.ClientIP(c.Ctx.Request()),
	}
	refer := c.Ctx.GetHeader("refer")
	if refer == "" {
		refer = "/public/page.html?from=login"
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), loginUser)
	comm.Redirect(c.Ctx.ResponseWriter(), refer)
}

//登出
func (c *IndexController) GetLoginout() {
	refer := c.Ctx.GetHeader("refer")
	if refer == "" {
		refer = "/public/index.html?from=logout"
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	comm.Redirect(c.Ctx.ResponseWriter(), refer)
}
