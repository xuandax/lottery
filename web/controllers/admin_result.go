package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/xuanxiaox/lottery/models"
	"github.com/xuanxiaox/lottery/services"
)

type AdminResultController struct {
	Ctx            iris.Context
	ServiceGift    services.GiftService
	ServiceBlack   services.BlackipService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserday services.UserdayService
	ServiceUser    services.UserService
}

func (c *AdminResultController) Get() mvc.Result {
	giftId := c.Ctx.URLParamIntDefault("gift_id", 0)
	page := c.Ctx.URLParamIntDefault("page", 1)
	size := 20
	pageNext := ""
	pagePrev := ""
	var dataList = make([]models.LtResult, 0)
	if giftId > 0 {
		dataList = c.ServiceResult.GetByGiftId(giftId)
	} else {
		dataList = c.ServiceResult.GetAll(page, size)
	}

	total := (page-1)*size + len(dataList)
	if len(dataList) >= size {
		if giftId > 0 {
			total = int(c.ServiceResult.CountByGiftId(giftId))
		} else {
			total = int(c.ServiceResult.CountAll())
		}
		pageNext = fmt.Sprintf("%d", page+1)
	}

	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}

	return mvc.View{
		Name:   "admin/result.html",
		Layout: "admin/layout.html",
		Data: iris.Map{
			"Title":    "后台管理",
			"Channel":  "result",
			"Datalist": dataList,
			"Total":    total,
			"PagePrev": pagePrev,
			"PageNext": pageNext,
		},
	}
}

//数据删除
func (c *AdminResultController) GetDelete() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	if id > 0 {
		c.ServiceResult.Delete(id)
	}
	return mvc.Response{
		Path: "/admin/result",
	}
}

//数据重置
func (c *AdminResultController) GetReset() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	data := models.LtResult{Id: id, SysStatus: 0}
	if id > 0 {
		c.ServiceResult.Update(&data, []string{"sys_status"})
	}
	return mvc.Response{
		Path: "/admin/result",
	}
}

//数据重置
func (c *AdminResultController) GetCheat() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	data := models.LtResult{Id: id, SysStatus: 2}
	if id > 0 {
		c.ServiceResult.Update(&data, []string{"sys_status"})
	}
	return mvc.Response{
		Path: "/admin/result",
	}
}
