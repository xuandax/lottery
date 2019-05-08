package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/xuanxiaox/lottery/comm"
	"github.com/xuanxiaox/lottery/models"
	"github.com/xuanxiaox/lottery/services"
)

type AdminUserController struct {
	Ctx            iris.Context
	ServiceGift    services.GiftService
	ServiceBlack   services.BlackipService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserday services.UserdayService
	ServiceUser    services.UserService
}

func (c *AdminUserController) Get() mvc.Result {
	page := c.Ctx.URLParamIntDefault("page", 1)
	size := 20
	pageNext := ""
	pagePrev := ""
	var dataList = make([]models.LtUser, 0)
	dataList = c.ServiceUser.GetAll(page, size)

	total := (page-1)*size + len(dataList)
	if len(dataList) >= size {
		total = int(c.ServiceUser.CountAll())
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
			"Channel":  "user",
			"Datalist": dataList,
			"Total":    total,
			"PagePrev": pagePrev,
			"PageNext": pageNext,
			"Now":      comm.NowUnix(),
		},
	}
}

//拉黑
func (c *AdminUserController) GetBlack() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	setTime := c.Ctx.URLParamIntDefault("time", 0)
	if id > 0 && setTime >= 0 {
		setTime = setTime*86400 + comm.NowUnix()
		data := models.LtUser{Id: id, BlackTime: setTime, SysUpdated: comm.NowUnix()}
		c.ServiceUser.Update(&data, []string{"black_time", "sys_updated"})
	}
	return mvc.Response{
		Path: "/admin/user",
	}
}
