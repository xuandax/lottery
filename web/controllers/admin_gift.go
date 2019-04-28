package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/xuanxiaox/lottery/comm"
	"github.com/xuanxiaox/lottery/models"
	"github.com/xuanxiaox/lottery/services"
	"github.com/xuanxiaox/lottery/web/viewmodels"
	"time"
)

type AdminGiftController struct {
	Ctx            iris.Context
	ServiceGift    services.GiftService
	ServiceBlack   services.BlackipService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserday services.UserdayService
	ServiceUser    services.UserService
}

//获取奖品列表
func (c *AdminGiftController) Get() mvc.Result {
	dataList := c.ServiceGift.GetAll()
	for i, giftInfo := range dataList {
		prizeData := make([][2]int, 0)
		//json反解析
		err := json.Unmarshal([]byte(giftInfo.PrizeData), &prizeData)
		if err != nil || prizeData == nil || len(prizeData) < 1 {
			dataList[i].PrizeData = ""
		} else {
			newPrizeData := make([]string, 0)
			//循环处理时间格式
			for index, pd := range prizeData {
				ct := comm.FormatFromUnixTime(int64(pd[0]))
				newPrizeData[index] = fmt.Sprintf("[%s]:%d", ct, pd[1])
			}
			//json解析重新赋值
			str, err := json.Marshal(newPrizeData)
			if err == nil && len(str) > 0 {
				dataList[i].PrizeData = string(str)
			} else {
				dataList[i].PrizeData = "[]"
			}
		}
	}
	total := len(dataList)
	return mvc.View{
		Name:   "admin/gift.html",
		Layout: "admin/layout.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "gift",
			"Datalist": dataList,
			"Total":    total,
		},
		Code: 0,
		Err:  nil,
	}
}

//编辑
func (c *AdminGiftController) GetEdit() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	giftInfo := viewmodels.ViewGift{}
	if id > 0 {
		data := c.ServiceGift.Get(id)
		giftInfo.Id = data.Id
		giftInfo.Title = data.Title
		giftInfo.PrizeNum = data.PrizeNum
		giftInfo.PrizeCode = data.PrizeCode
		giftInfo.PrizeTime = data.PrizeTime
		giftInfo.Img = data.Img
		giftInfo.PrizeOrder = data.PrizeOrder
		giftInfo.Gtype = data.Gtype
		giftInfo.Gdata = data.Gdata
		giftInfo.TimeBegin = comm.FormatFromUnixTime(int64(data.TimeBegin))
		giftInfo.TimeEnd = comm.FormatFromUnixTime(int64(data.TimeEnd))
	}
	return mvc.View{
		Name:   "admin/giftEdit.html",
		Layout: "admin/layout.html",
		Data: iris.Map{
			"Title":   "后台管理",
			"Channel": "gift",
			"info":    giftInfo,
		},
	}
}

//数据提交
func (c *AdminGiftController) PostSave() mvc.Result {
	data := viewmodels.ViewGift{}
	err := c.Ctx.ReadForm(&data)
	if err != nil {
		return mvc.Response{
			Text: fmt.Sprintf("admin_gift PostSave err = %s", err),
		}
	}
	giftInfo := models.LtGift{}
	giftInfo.Id = data.Id
	giftInfo.Title = data.Title
	giftInfo.PrizeNum = data.PrizeNum
	giftInfo.PrizeCode = data.PrizeCode
	giftInfo.PrizeTime = data.PrizeTime
	giftInfo.Img = data.Img
	giftInfo.PrizeOrder = data.PrizeOrder
	giftInfo.Gtype = data.Gtype
	giftInfo.Gdata = data.Gdata
	t1, err1 := comm.ParseTime(data.TimeBegin)
	t2, err2 := comm.ParseTime(data.TimeEnd)
	if err1 != nil || err2 != nil {
		return mvc.Response{
			Text: fmt.Sprintf("admin_gift PostSave parseTime err1 = %s, err2 = %s", err1, err2),
		}
	}
	giftInfo.TimeBegin = int(t1.Unix())
	giftInfo.TimeEnd = int(t2.Unix())

	if giftInfo.Id > 0 {
		dataInfo := c.ServiceGift.Get(giftInfo.Id)
		if dataInfo != nil {
			giftInfo.SysIp = comm.ClientIP(c.Ctx.Request())
			giftInfo.SysUpdated = int(time.Now().Unix())
			//奖品总数发生变化
			if dataInfo.PrizeNum != giftInfo.PrizeNum {
				giftInfo.LeftNum = giftInfo.LeftNum - (dataInfo.PrizeNum - giftInfo.PrizeNum)
				if giftInfo.LeftNum < 0 || giftInfo.PrizeNum <= 0 {
					giftInfo.LeftNum = 0
				}
			} else {
				giftInfo.LeftNum = giftInfo.PrizeNum
			}
			c.ServiceGift.Update(&giftInfo, []string{"title", "prize_num", "prize_code", "prize_time",
				"left_num", "prize_num", "img", "prize_order", "gtype", "gdata", "time_begin", "time_end", "sys_updated"})
		}
	}
	if giftInfo.Id == 0 {
		giftInfo.LeftNum = giftInfo.PrizeNum
		giftInfo.SysCreated = int(time.Now().Unix())
		giftInfo.SysIp = comm.ClientIP(c.Ctx.Request())
		c.ServiceGift.Create(&giftInfo)
	}

	return mvc.Response{
		Path: "/admin/gift",
	}
}

//数据删除
func (c *AdminGiftController) GetDelete() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	if id > 0 {
		c.ServiceGift.Delete(id)
	}
	return mvc.Response{
		Path: "/admin/gift",
	}
}

//数据重置
func (c *AdminGiftController) GetReset() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	data := models.LtGift{Id: id, SysStatus: 0}
	if id > 0 {
		c.ServiceGift.Update(&data, []string{"sys_status"})
	}
	return mvc.Response{
		Path: "/admin/gift",
	}
}
