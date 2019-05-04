package utils

import (
	"github.com/xuanxiaox/lottery/comm"
	"github.com/xuanxiaox/lottery/services"
	"log"
)

func PrizeGift(giftId int, leftNum int) bool {
	giftService := services.NewGiftService()
	rows, err := giftService.DecrLeftNum(giftId, 1)
	if rows < 1 || err != nil {
		log.Println("prizedata.go PrizeGift err=", err, ", rows=", rows)
		return false
	}
	return true
}

func PrizeCodeDiff(giftId int, service services.CodeService) string {
	lockUid := 0 - giftId - 100000
	LockUserLucky(lockUid)
	defer UnlockUserLucky(lockUid)

	//获取抽奖码
	codeId := 0
	codeInfo := service.NextUsingCode(giftId, codeId)
	if codeInfo != nil || codeInfo.Id > 0 {
		codeInfo.SysStatus = 2
		codeInfo.SysUpdated = comm.NowUnix()
		service.Update(codeInfo, nil)
	} else {
		log.Println("prizedata.go PrizeCodeDiff err, gift_id=", giftId)
		return ""
	}
	return codeInfo.Code
}
