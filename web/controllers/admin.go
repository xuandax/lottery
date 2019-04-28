package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/xuanxiaox/lottery/services"
)

type AdminController struct {
	Ctx            iris.Context
	ServiceGift    services.GiftService
	ServiceBlack   services.BlackipService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserday services.UserdayService
	ServiceUser    services.UserService
}

func (c *AdminController) Get() mvc.Result {
	dataList := c.ServiceGift.Get(1)
	return mvc.View{
		Name:   "admin/index.html",
		Layout: "admin/layout",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "",
			"data":    dataList,
		},
		Code: 0,
		Err:  nil,
	}
}
