package services

import (
	"github.com/xuanxiaox/lottery/dao"
	"github.com/xuanxiaox/lottery/datasource"
	"github.com/xuanxiaox/lottery/models"
)

type UserService interface {
	Get(id int) *models.LtUser
	GetAll(page, size int) []models.LtUser
	CountAll() int64
	//Delete(id int) error
	Create(data *models.LtUser) error
	Update(data *models.LtUser, columns []string) error
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.InstanceMaster()),
	}
}

func (s *userService) Get(id int) *models.LtUser {
	return s.dao.Get(id)
}
func (s *userService) GetAll(page, size int) []models.LtUser {
	return s.dao.GetAll(page, size)
}
func (s *userService) CountAll() int64 {
	return s.dao.CountAll()
}

//func (s *userService) Delete(id int) error {
//	return s.dao.Delete(id)
//}
func (s *userService) Create(data *models.LtUser) error {
	return s.dao.Create(data)
}
func (s *userService) Update(data *models.LtUser, columns []string) error {
	return s.dao.Update(data, columns)
}
