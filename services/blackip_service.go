package services

import (
	"github.com/xuanxiaox/lottery/dao"
	"github.com/xuanxiaox/lottery/datasource"
	"github.com/xuanxiaox/lottery/models"
)

type BlackipService interface {
	Get(id int) *models.LtBlackip
	GetAll(page, size int) []models.LtBlackip
	CountAll() int64
	//Delete(id int) error
	Create(data *models.LtBlackip) error
	Update(data *models.LtBlackip, columns []string) error
}

type blackipService struct {
	dao *dao.BlackipDao
}

func NewBlackipService() BlackipService {
	return &blackipService{
		dao: dao.NewBlackipDao(datasource.InstanceMaster()),
	}
}

func (s *blackipService) Get(id int) *models.LtBlackip {
	return s.dao.Get(id)
}
func (s *blackipService) GetAll(page, size int) []models.LtBlackip {
	return s.dao.GetAll(page, size)
}
func (s *blackipService) CountAll() int64 {
	return s.dao.CountAll()
}

//func (s *blackipService) Delete(id int) error {
//	return s.dao.Delete(id)
//}
func (s *blackipService) Create(data *models.LtBlackip) error {
	return s.dao.Create(data)
}
func (s *blackipService) Update(data *models.LtBlackip, columns []string) error {
	return s.dao.Update(data, columns)
}
