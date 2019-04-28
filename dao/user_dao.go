package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/xuanxiaox/lottery/models"
)

type UserDao struct {
	engine *xorm.Engine
}

//创建UserDao
func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

//根据ID获取数据
func (d *UserDao) Get(id int) *models.LtUser {
	data := &models.LtUser{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

//获取所有数据
func (d *UserDao) GetAll(page, size int) []models.LtUser {
	dataList := make([]models.LtUser, 0)
	offset := (page - 1) * size
	_ = d.engine.Desc("id").Limit(size, offset).Find(&dataList)
	return dataList
}

//获取总数
func (d *UserDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtUser{})
	if err != nil {
		return 0
	}
	return num
}

//软删除
//func(d *UserDao) Delete(id int) error {
//	data := models.LtUser{Id:id, SysStatus:1}
//	_, err := d.engine.Id(data.Id).Update(data)
//	return err
//}

//更新
func (d *UserDao) Update(data *models.LtUser, colums []string) error {
	_, err := d.engine.Id(data.Id).Cols(colums...).Update(data)
	return err
}

//创建
func (d *UserDao) Create(data *models.LtUser) error {
	_, err := d.engine.Insert(data)
	return err
}

//根据IP查询
func (d *UserDao) GetByIp(ip string) []*models.LtUser {
	dataList := make([]*models.LtUser, 0)
	err := d.engine.Where("ip = ?", ip).Desc("ip").Find(&dataList)
	if err != nil || len(dataList) < 1 {
		return nil
	}
	return dataList
}
