package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/xuanxiaox/lottery/comm"
	"github.com/xuanxiaox/lottery/models"
	"github.com/xuanxiaox/lottery/services"
)

type AdminBlackipController struct {
	Ctx               iris.Context
	ServiceGift       services.GiftService
	ServiceBlackip    services.BlackipService
	ServiceCode       services.CodeService
	ServiceResult     services.ResultService
	ServiceBlackipday services.UserdayService
	UserService       services.UserService
}

func (c *AdminBlackipController) Get() mvc.Result {
	page := c.Ctx.URLParamIntDefault("page", 1)
	size := 20
	pageNext := ""
	pagePrev := ""
	var dataList = make([]models.LtBlackip, 0)
	dataList = c.ServiceBlackip.GetAll(page, size)

	total := (page-1)*size + len(dataList)
	if len(dataList) >= size {
		total = int(c.ServiceBlackip.CountAll())
		pageNext = fmt.Sprintf("%d", page+1)
	}

	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}

	return mvc.View{
		Name:   "admin/user.html",
		Layout: "admin/layout.html",
		Data: iris.Map{
			"Title":    "后台管理",
			"Channel":  "blackip",
			"Datalist": dataList,
			"Total":    total,
			"PagePrev": pagePrev,
			"PageNext": pageNext,
		},
	}
}

//拉黑
func (c *AdminBlackipController) GetBlack() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	setTime := c.Ctx.URLParamIntDefault("time", 0)
	data := models.LtBlackip{Id: id, BlackTime: setTime, SysUpdated: comm.NowUnix()}
	if id > 0 && setTime > 0 {
		c.ServiceBlackip.Update(&data, []string{"black_time", "sys_updated"})
	}
	return mvc.Response{
		Path: "/admin/user",
	}
}
