package services

import (
	"github.com/xuanxiaox/lottery/dao"
	"github.com/xuanxiaox/lottery/datasource"
	"github.com/xuanxiaox/lottery/models"
)

type UserdayService interface {
	Get(id int) *models.LtUserday
	GetAll() []models.LtUserday
	CountAll() int64
	//Delete(id int) error
	Create(data *models.LtUserday) error
	Update(data *models.LtUserday, columns []string) error
}

type userdayService struct {
	dao *dao.UserdayDao
}

func NewUserdayService() UserdayService {
	return &userdayService{
		dao: dao.NewUserdayDao(datasource.InstanceMaster()),
	}
}

func (s *userdayService) Get(id int) *models.LtUserday {
	return s.dao.Get(id)
}
func (s *userdayService) GetAll() []models.LtUserday {
	return s.dao.GetAll()
}
func (s *userdayService) CountAll() int64 {
	return s.dao.CountAll()
}

//func (s *userdayService) Delete(id int) error {
//	return s.dao.Delete(id)
//}
func (s *userdayService) Create(data *models.LtUserday) error {
	return s.dao.Create(data)
}
func (s *userdayService) Update(data *models.LtUserday, columns []string) error {
	return s.dao.Update(data, columns)
}
