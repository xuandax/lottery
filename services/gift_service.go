package services

import (
	"github.com/xuanxiaox/lottery/dao"
	"github.com/xuanxiaox/lottery/datasource"
	"github.com/xuanxiaox/lottery/models"
)

type GiftService interface {
	Get(id int) *models.LtGift
	GetAll() []models.LtGift
	CountAll() int64
	Delete(id int) error
	Create(data *models.LtGift) error
	Update(data *models.LtGift, columns []string) error
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
