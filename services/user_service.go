package services

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/xuanxiaox/lottery/comm"
	"github.com/xuanxiaox/lottery/dao"
	"github.com/xuanxiaox/lottery/datasource"
	"github.com/xuanxiaox/lottery/models"
	"log"
)

type UserService interface {
	Get(id int) *models.LtUser
	GetAll(page, size int) []models.LtUser
	CountAll() int64
	//Delete(id int) error
	Create(data *models.LtUser) error
	Update(data *models.LtUser, columns []string) error
	GetByCache(uid int) *models.LtUser
	SetByCache(user *models.LtUser)
	UpdateByCache(user *models.LtUser, columns []string)
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
	user := s.GetByCache(id)
	if user == nil || user.Id < 0 {
		user = s.dao.Get(id)
		if user == nil || user.Id < 0 {
			user = &models.LtUser{
				Id: id,
			}
		}
		s.SetByCache(user)
	}
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
	s.UpdateByCache(data, columns)
	return s.dao.Update(data, columns)
}

func (s *userService) GetByCache(uid int) *models.LtUser {
	key := fmt.Sprintf("lt_user_info_%d", uid)
	cache := datasource.InstanceRedisConn()
	userMap, err := redis.StringMap(cache.Do("HGETALL", key))
	if err != nil {
		log.Println("user_service.go GetByCache redis.StringMap(cache.Do(HGETALL, key)), err =", err)
		return nil
	}
	dataId := comm.GetInt64FromStringMap(userMap, "Id", 0)
	if dataId < 1 {
		return nil
	}
	data := &models.LtUser{
		Id:         int(dataId),
		Username:   comm.GetStringFromStringMap(userMap, "Username", ""),
		BlackTime:  int(comm.GetInt64FromStringMap(userMap, "BlackTime", 0)),
		Realname:   comm.GetStringFromStringMap(userMap, "Realname", ""),
		Mobile:     comm.GetStringFromStringMap(userMap, "Mobile", ""),
		Address:    comm.GetStringFromStringMap(userMap, "Address", ""),
		SysCreated: int(comm.GetInt64FromStringMap(userMap, "SysCreated", 0)),
		SysUpdated: int(comm.GetInt64FromStringMap(userMap, "SysUpdated", 0)),
		SysIp:      comm.GetStringFromStringMap(userMap, "SysIp", ""),
	}
	return data
}

func (s *userService) SetByCache(user *models.LtUser) {
	if user == nil || user.Id < 0 {
		return
	}
	id := user.Id
	key := fmt.Sprintf("lt_user_info_%d", id)
	cache := datasource.InstanceRedisConn()
	userMap := map[string]interface{}{}
	if user.Username != "" {
		userMap = map[string]interface{}{
			"Id":         id,
			"Username":   user.Username,
			"BlackTime":  user.BlackTime,
			"Realname":   user.Realname,
			"Mobile":     user.Mobile,
			"Address":    user.Address,
			"SysCreated": user.SysCreated,
			"SysUpdated": user.SysUpdated,
			"SysIp":      user.SysIp,
		}
	}
	_, err := cache.Do("HMSET", redis.Args{}.Add(key).AddFlat(userMap)...)
	if err != nil {
		log.Println("user_service.go SetByCache HMSET err = ", err)
	}
}

func (s *userService) UpdateByCache(user *models.LtUser, columns []string) {
	if user == nil || user.Id < 0 {
		return
	}
	key := fmt.Sprintf("lt_user_info_%d", user.Id)
	cache := datasource.InstanceRedisConn()
	cache.Do("DEL", key)
}
