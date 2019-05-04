package services

import (
	"fmt"
	"github.com/xuanxiaox/lottery/dao"
	"github.com/xuanxiaox/lottery/datasource"
	"github.com/xuanxiaox/lottery/models"
	"strconv"
	"time"
)

type UserdayService interface {
	Get(id int) *models.LtUserday
	GetAll() []models.LtUserday
	CountAll() int64
	//Delete(id int) error
	Create(data *models.LtUserday) error
	Update(data *models.LtUserday, columns []string) error
	GetUserToday(uid int) *models.LtUserday
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

func (s *userdayService) GetUserToday(uid int) *models.LtUserday {
	y, m, d := time.Now().Date()
	dayStr := fmt.Sprintf("%d%02d%02d", y, m, d)
	day, _ := strconv.Atoi(dayStr)
	userDay := s.dao.Search(uid, day)
	if len(userDay) > 0 {
		return &userDay[0]
	}
	return nil
}
