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

type BlackipService interface {
	Get(id int) *models.LtBlackip
	GetAll(page, size int) []models.LtBlackip
	CountAll() int64
	//Delete(id int) error
	Create(data *models.LtBlackip) error
	Update(data *models.LtBlackip, columns []string) error
	GetByIp(ip string) *models.LtBlackip
	GetByCache(ip string) *models.LtBlackip
	SetByCache(blackip *models.LtBlackip)
	UpdateByCache(blackip *models.LtBlackip, columns []string)
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
	s.UpdateByCache(data, columns)
	return s.dao.Update(data, columns)
}

func (s *blackipService) GetByIp(ip string) *models.LtBlackip {
	blackip := s.GetByCache(ip)
	if blackip == nil || blackip.Ip == "" {
		blackip = s.dao.GetByIp(ip)
		if blackip == nil || blackip.Ip == "" {
			blackip = &models.LtBlackip{Ip: ip}
		}
		s.SetByCache(blackip)
	}
	return blackip
}

func (s *blackipService) GetByCache(ip string) *models.LtBlackip {
	if ip == "" {
		return nil
	}
	key := fmt.Sprintf("lt_blackip_info_%s", ip)
	cache := datasource.InstanceRedisConn()
	rs, err := redis.StringMap(cache.Do("HGETALL", key))
	if err != nil {
		log.Println("blackip_service.go GetByCache HGETALL ,err = ", err)
		return nil
	}
	dataIp := comm.GetStringFromStringMap(rs, "Ip", "")
	if dataIp != "" {
		blackip := &models.LtBlackip{
			Id:         int(comm.GetInt64FromStringMap(rs, "Id", 0)),
			Ip:         dataIp,
			BlackTime:  int(comm.GetInt64FromStringMap(rs, "BlackTime", 0)),
			SysCreated: int(comm.GetInt64FromStringMap(rs, "SysCreated", 0)),
			SysUpdated: int(comm.GetInt64FromStringMap(rs, "SysUpdated", 0)),
		}
		return blackip
	}
	return nil
}

func (s *blackipService) SetByCache(blackip *models.LtBlackip) {
	if blackip == nil || blackip.Ip == "" {
		return
	}
	key := fmt.Sprintf("lt_blackip_info_%s", blackip.Ip)
	cache := datasource.InstanceRedisConn()
	blackipMap := map[string]interface{}{}
	if blackip.Id > 0 {
		blackipMap = map[string]interface{}{
			"Id":         blackip.Id,
			"Ip":         blackip.Ip,
			"BlackTime":  blackip.BlackTime,
			"SysCreated": blackip.SysCreated,
			"SysUpdated": blackip.SysUpdated,
		}
	}
	_, err := cache.Do("HMSET", redis.Args{}.Add(key).AddFlat(blackipMap)...)
	if err != nil {
		log.Println("blackip_service.go SetByCache HMSET ,err = ", err)
	}
}

func (s *blackipService) UpdateByCache(blackip *models.LtBlackip, columns []string) {
	if blackip == nil || blackip.Ip == "" {
		return
	}
	key := fmt.Sprintf("lt_blackip_info_%s", blackip.Ip)
	cache := datasource.InstanceRedisConn()
	cache.Do("DEL", key)
}
