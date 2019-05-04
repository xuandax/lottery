package services

import (
	"github.com/xuanxiaox/lottery/dao"
	"github.com/xuanxiaox/lottery/datasource"
	"github.com/xuanxiaox/lottery/models"
	"strconv"
	"strings"
)

type GiftService interface {
	Get(id int) *models.LtGift
	GetAll() []models.LtGift
	CountAll() int64
	Delete(id int) error
	Create(data *models.LtGift) error
	Update(data *models.LtGift, columns []string) error
	GetAllUse() []models.ObjGiftPrize
	DecrLeftNum(giftId, num int) (int64, error)
	IncrLeftNum(giftId, num int) (int64, error)
}

type giftService struct {
	dao *dao.GiftDao
}

func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(datasource.InstanceMaster()),
	}
}

func (s *giftService) Get(id int) *models.LtGift {
	return s.dao.Get(id)
}
func (s *giftService) GetAll() []models.LtGift {
	return s.dao.GetAll()
}
func (s *giftService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *giftService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *giftService) Create(data *models.LtGift) error {
	return s.dao.Create(data)
}
func (s *giftService) Update(data *models.LtGift, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *giftService) GetAllUse() []models.ObjGiftPrize {
	var ltGiftList []models.LtGift
	ltGiftList = s.dao.GetAllUse()
	var objGiftList []models.ObjGiftPrize
	for _, gift := range ltGiftList {
		codes := strings.Split(gift.PrizeCode, "-")
		if len(codes) == 2 {
			codeA, errA := strconv.Atoi(codes[0])
			codeB, errB := strconv.Atoi(codes[1])
			if errA == nil && errB == nil && codeA > 0 && codeA < codeB && codeB < 10000 {
				data := models.ObjGiftPrize{
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

func (s *giftService) DecrLeftNum(giftId, num int) (int64, error) {
	return s.dao.DecrLeftNum(giftId, num)
}

func (s *giftService) IncrLeftNum(giftId, num int) (int64, error) {
	return s.dao.IncrLeftNum(giftId, num)
}
