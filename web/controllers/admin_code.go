package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/xuanxiaox/lottery/models"
	"github.com/xuanxiaox/lottery/services"
)

type AdminCodeController struct {
	Ctx            iris.Context
	ServiceGift    services.GiftService
	ServiceBlack   services.BlackipService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserday services.UserdayService
	ServiceUser    services.UserService
}

func (c *AdminCodeController) Get() mvc.Result {
	giftId := c.Ctx.URLParamIntDefault("gift_id", 0)
	page := c.Ctx.URLParamIntDefault("page", 1)
	size := 3
	pageNext := ""
	pagePrev := ""

	//获取数据列表
	var dataList []models.LtCode
	if giftId > 0 {
		dataList = c.ServiceCode.GetByGiftId(giftId)
	} else {
		dataList = c.ServiceCode.GetAll(page, size)
	}
	total := (page-1)*size + len(dataList)
	if len(dataList) >= size {
		if giftId > 0 {
			total = int(c.ServiceCode.CountByGiftId(giftId))
		} else {
			total = int(c.ServiceCode.CountAll())
		}
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}
	return mvc.View{
		Name:   "admin/code.html",
		Layout: "admin/layout",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "code",
			"GiftId":   giftId,
			"Datalist": dataList,
			"Total":    total,
			"PagePrev": pagePrev,
			"PageNext": pageNext,
		},
	}
}

//数据删除
func (c *AdminCodeController) GetDelete() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	if id > 0 {
		c.ServiceCode.Delete(id)
	}
	return mvc.Response{
		Path: "/admin/code",
	}
}

//数据重置
func (c *AdminCodeController) GetReset() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	data := models.LtCode{Id: id, SysStatus: 0}
	if id > 0 {
		c.ServiceCode.Update(&data, []string{"sys_status"})
	}
	return mvc.Response{
		Path: "/admin/code",
	}
}
