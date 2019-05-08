package services

import (
	"encoding/json"
	"github.com/xuanxiaox/lottery/comm"
	"github.com/xuanxiaox/lottery/dao"
	"github.com/xuanxiaox/lottery/datasource"
	"github.com/xuanxiaox/lottery/models"
	"log"
	"strconv"
	"strings"
)

type GiftService interface {
	Get(id int, useCache bool) *models.LtGift
	GetAll(useCache bool) []models.LtGift
	CountAll() int64
	Delete(id int) error
	Create(data *models.LtGift) error
	Update(data *models.LtGift, columns []string) error
	GetAllUse(useCache bool) []models.ObjGiftPrize
	DecrLeftNum(giftId, num int) (int64, error)
	IncrLeftNum(giftId, num int) (int64, error)
	GetAllByCache() []models.LtGift
	SetAllByCache(gifts []models.LtGift)
	UpdateByCache(data *models.LtGift, columns []string)
}

type giftService struct {
	dao *dao.GiftDao
}

func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(datasource.InstanceMaster()),
	}
}

func (s *giftService) Get(id int, useCache bool) *models.LtGift {
	if !useCache {
		return s.dao.Get(id)
	}
	gifts := s.GetAll(true)
	for _, gift := range gifts {
		if gift.Id == id {
			return &gift
		}
	}
	return nil
}
func (s *giftService) GetAll(useCache bool) []models.LtGift {
	if !useCache {
		return s.dao.GetAll()
	}
	gifts := s.GetAllByCache()
	if len(gifts) < 1 {
		gifts = s.dao.GetAll()
		s.SetAllByCache(gifts)
	}
	return gifts
}
func (s *giftService) CountAll() int64 {
	return int64(len(s.GetAll(true)))
}
func (s *giftService) Delete(id int) error {
	data := &models.LtGift{Id: id}
	s.UpdateByCache(data, nil)
	return s.dao.Delete(id)
}
func (s *giftService) Create(data *models.LtGift) error {
	s.UpdateByCache(data, nil)
	return s.dao.Create(data)
}
func (s *giftService) Update(data *models.LtGift, columns []string) error {
	s.UpdateByCache(data, columns)
	return s.dao.Update(data, columns)
}

func (s *giftService) GetAllUse(useCache bool) []models.ObjGiftPrize {
	var ltGiftList []models.LtGift
	if !useCache {
		ltGiftList = s.dao.GetAllUse()
	} else {
		now := comm.NowUnix()
		gifts := s.GetAll(true)
		for _, gift := range gifts {
			if gift.Id > 0 && gift.SysStatus == 0 &&
				gift.PrizeNum > 0 && gift.LeftNum > 0 &&
				gift.TimeBegin <= now && gift.TimeEnd > now {
				ltGiftList = append(ltGiftList, gift)
			}
		}
	}

	if ltGiftList != nil {
		var objGiftList []models.ObjGiftPrize
		for _, gift := range ltGiftList {
			codes := strings.Split(gift.PrizeCode, "-")
			if len(codes) == 2 {
				codeA, errA := strconv.Atoi(codes[0])
				codeB, errB := strconv.Atoi(codes[1])
				if errA == nil && errB == nil && codeA > 0 && codeA < codeB && codeB < 10000 {
					data := models.ObjGiftPrize{
						Id:         gift.Id,
						Title:      gift.Title,
						PrizeNum:   gift.PrizeNum,
						LeftNum:    gift.LeftNum,
						PrizeCodeA: codeA,
						PrizeCodeB: codeB,
						Img:        gift.Img,
						PrizeOrder: gift.PrizeOrder,
						Gtype:      gift.Gtype,
						Gdata:      gift.Gdata,
					}
					objGiftList = append(objGiftList, data)
				}
			}
		}
		return objGiftList
	}

	return []models.ObjGiftPrize{}
}

func (s *giftService) DecrLeftNum(giftId, num int) (int64, error) {
	return s.dao.DecrLeftNum(giftId, num)
}

func (s *giftService) IncrLeftNum(giftId, num int) (int64, error) {
	return s.dao.IncrLeftNum(giftId, num)
}

func (s *giftService) GetAllByCache() []models.LtGift {
	key := "gift_all"
	cache := datasource.InstanceRedisConn()
	rs, err := cache.Do("GET", key)
	if err != nil {
		log.Println("gift_service.go getAllByCache cache GET key = ", key, ",err = ", err)
		return nil
	}
	rsStr := comm.GetString(rs, "")
	if rsStr == "" {
		return nil
	}
	dataList := []map[string]interface{}{}
	err = json.Unmarshal([]byte(rsStr), &dataList)
	if err != nil {
		log.Println("gift_service.go getAllByCache json.Unmarshal err = ", err)
		return nil
	}
	//将map转换
	gifts := make([]models.LtGift, len(dataList))
	for i := 0; i < len(dataList); i++ {
		data := dataList[i]
		Id := comm.GetInt64FromMap(data, "Id", 0)
		if Id <= 0 {
			gifts[i] = models.LtGift{}
		} else {
			gifts[i] = models.LtGift{
				Id:         int(Id),
				Title:      comm.GetStringFromMap(data, "Title", ""),
				PrizeNum:   int(comm.GetInt64FromMap(data, "PrizeNum", 0)),
				LeftNum:    int(comm.GetInt64FromMap(data, "LeftNum", 0)),
				PrizeCode:  comm.GetStringFromMap(data, "PrizeCode", ""),
				PrizeTime:  int(comm.GetInt64FromMap(data, "PrizeTime", 0)),
				Img:        comm.GetStringFromMap(data, "Img", ""),
				PrizeOrder: int(comm.GetInt64FromMap(data, "PrizeOrder", 0)),
				Gtype:      int(comm.GetInt64FromMap(data, "Gtype", 0)),
				Gdata:      comm.GetStringFromMap(data, "Gdata", ""),
				TimeBegin:  int(comm.GetInt64FromMap(data, "TimeBegin", 0)),
				TimeEnd:    int(comm.GetInt64FromMap(data, "TimeEnd", 0)),
				//PrizeData:    comm.GetStringFromMap(data, "PrizeData", ""),
				PrizeBegin: int(comm.GetInt64FromMap(data, "PrizeBegin", 0)),
				PrizeEnd:   int(comm.GetInt64FromMap(data, "PrizeEnd", 0)),
				SysStatus:  int(comm.GetInt64FromMap(data, "SysStatus", 0)),
				SysCreated: int(comm.GetInt64FromMap(data, "SysCreated", 0)),
				SysUpdated: int(comm.GetInt64FromMap(data, "SysUpdated", 0)),
				SysIp:      comm.GetStringFromMap(data, "SysIp", ""),
			}
		}
	}
	return gifts
}

func (s *giftService) SetAllByCache(gifts []models.LtGift) {
	giftsStr := ""
	if len(gifts) > 0 {
		dataList := make([]map[string]interface{}, len(gifts))
		for i := 0; i < len(gifts); i++ {
			gift := gifts[i]
			data := make(map[string]interface{})
			data["Id"] = gift.Id
			data["Title"] = gift.Title
			data["PrizeNum"] = gift.PrizeNum
			data["LeftNum"] = gift.LeftNum
			data["PrizeCode"] = gift.PrizeCode
			data["PrizeTime"] = gift.PrizeTime
			data["Img"] = gift.Img
			data["PrizeOrder"] = gift.PrizeOrder
			data["Gtype"] = gift.Gtype
			data["Gdata"] = gift.Gdata
			data["TimeBegin"] = gift.TimeBegin
			data["TimeEnd"] = gift.TimeEnd
			//data["PrizeData"] = gift.PrizeData
			data["PrizeBegin"] = gift.PrizeBegin
			data["PrizeEnd"] = gift.PrizeEnd
			data["SysStatus"] = gift.SysStatus
			data["SysCreated"] = gift.SysCreated
			data["SysUpdated"] = gift.SysUpdated
			data["SysIp"] = gift.SysIp
			dataList[i] = data
		}
		giftsByte, err := json.Marshal(dataList)
		if err != nil {
			log.Println("gift_service.go SetAllByCache  json.Marshal err = ", err)
		}
		giftsStr = string(giftsByte)
	}
	key := "gift_all"
	cache := datasource.InstanceRedisConn()
	_, err := cache.Do("SET", key, giftsStr)
	if err != nil {
		log.Println("gift_service.go cache.Do SET  key", key, ",giftsStr", giftsStr, ", err = ", err)
	}
}

func (s *giftService) UpdateByCache(data *models.LtGift, columns []string) {
	if data == nil || data.Id <= 0 {
		return
	}
	key := "gift_all"
	cache := datasource.InstanceRedisConn()
	cache.Do("DEL", key)
}
