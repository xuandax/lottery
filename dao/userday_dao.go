package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/xuanxiaox/lottery/models"
)

type UserdayDao struct {
	engine *xorm.Engine
}

//创建UserdayDao
func NewUserdayDao(engine *xorm.Engine) *UserdayDao {
	return &UserdayDao{
		engine: engine,
	}
}

//根据ID获取数据
func (d *UserdayDao) Get(id int) *models.LtUserday {
	data := &models.LtUserday{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

//获取所有数据
func (d *UserdayDao) GetAll() []models.LtUserday {
	dataList := make([]models.LtUserday, 0)
	_ = d.engine.Desc("id").Find(&dataList)
	return dataList
}

//获取总数
func (d *UserdayDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtUserday{})
	if err != nil {
		return 0
	}
	return num
}

//软删除
//func(d *UserdayDao) Delete(id int) error {
//	data := models.LtUserday{Id:id, SysStatus:1}
//	_, err := d.engine.Id(data.Id).Update(data)
//	return err
//}

//更新
func (d *UserdayDao) Update(data *models.LtUserday, colums []string) error {
	_, err := d.engine.Id(data.Id).Cols(colums...).Update(data)
	return err
}

//创建
func (d *UserdayDao) Create(data *models.LtUserday) error {
	_, err := d.engine.Insert(data)
	return err
}

//根据IP查询
func (d *UserdayDao) GetByIp(ip string) []*models.LtUserday {
	dataList := make([]*models.LtUserday, 0)
	err := d.engine.Where("ip = ?", ip).Desc("ip").Find(&dataList)
	if err != nil || len(dataList) < 1 {
		return nil
	}
	return dataList
}

func (d *UserdayDao) Search(uid int, day int) []models.LtUserday {
	userDay := make([]models.LtUserday, 0)
	_ = d.engine.Where("uid=?", uid).Where("day=?", day).Desc("id").Find(&userDay)
	return userDay
}
