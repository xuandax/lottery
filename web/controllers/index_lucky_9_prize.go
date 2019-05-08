package controllers

import (
	"github.com/xuanxiaox/lottery/conf"
	"github.com/xuanxiaox/lottery/models"
)

func (c *IndexController) Prize(prizeCode int, limitBlack bool) *models.ObjGiftPrize {
	var prizeGift *models.ObjGiftPrize
	giftList := c.ServiceGift.GetAllUse(true)
	for _, gift := range giftList {
		if gift.PrizeCodeA <= prizeCode && gift.PrizeCodeB >= prizeCode {
			if !limitBlack && gift.Gtype < conf.GtypeGiftSmall {
				prizeGift = &gift
				break
			}
		}
	}
	return prizeGift
}
