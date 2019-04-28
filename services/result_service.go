package services

import (
	"github.com/xuanxiaox/lottery/dao"
	"github.com/xuanxiaox/lottery/datasource"
	"github.com/xuanxiaox/lottery/models"
)

type ResultService interface {
	Get(id int) *models.LtResult
	GetAll(page, size int) []models.LtResult
	CountAll() int64
	Delete(id int) error
	Create(data *models.LtResult) error
	Update(data *models.LtResult, columns []string) error
	GetByGiftId(giftId int) []models.LtResult
	CountByGiftId(giftId int) int64
}

type resultService struct {
	dao *dao.ResultDao
}

func NewResultService() ResultService {
	return &resultService{
		dao: dao.NewResultDao(datasource.InstanceMaster()),
	}
}

func (s *resultService) Get(id int) *models.LtResult {
	return s.dao.Get(id)
}
func (s *resultService) GetAll(page, size int) []models.LtResult {
	return s.dao.GetAll(page, size)
}
func (s *resultService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *resultService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *resultService) Create(data *models.LtResult) error {
	return s.dao.Create(data)
}
func (s *resultService) Update(data *models.LtResult, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *resultService) GetByGiftId(giftId int) []models.LtResult {
	return s.dao.GetByGiftId(giftId)
}

func (s *resultService) CountByGiftId(giftId int) int64 {
	return s.dao.CountByGiftId(giftId)
}
