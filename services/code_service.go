package services

import (
	"github.com/xuanxiaox/lottery/dao"
	"github.com/xuanxiaox/lottery/datasource"
	"github.com/xuanxiaox/lottery/models"
)

type CodeService interface {
	Get(id int) *models.LtCode
	GetAll(page, size int) []models.LtCode
	CountAll() int64
	Delete(id int) error
	Create(data *models.LtCode) error
	Update(data *models.LtCode, columns []string) error
	GetByGiftId(giftId int) []models.LtCode
	CountByGiftId(giftId int) int64
	NextUsingCode(giftId, codeId int) *models.LtCode
}

type codeService struct {
	dao *dao.CodeDao
}

func NewCodeService() CodeService {
	return &codeService{
		dao: dao.NewCodeDao(datasource.InstanceMaster()),
	}
}

func (s *codeService) Get(id int) *models.LtCode {
	return s.dao.Get(id)
}
func (s *codeService) GetAll(page, size int) []models.LtCode {
	return s.dao.GetAll(page, size)
}
func (s *codeService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *codeService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *codeService) Create(data *models.LtCode) error {
	return s.dao.Create(data)
}
func (s *codeService) Update(data *models.LtCode, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *codeService) GetByGiftId(giftId int) []models.LtCode {
	return s.dao.GetByGiftId(giftId)
}
func (s *codeService) CountByGiftId(giftId int) int64 {
	return s.dao.CountByGiftId(giftId)
}

func (s *codeService) NextUsingCode(giftId, codeId int) *models.LtCode {
	return s.dao.NextUsingCode(giftId, codeId)
}
