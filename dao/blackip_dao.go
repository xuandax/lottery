package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/xuanxiaox/lottery/models"
)

type BlackipDao struct {
	engine *xorm.Engine
}

//创建BlackipDao
func newBlackipDao(engine *xorm.Engine) *BlackipDao {
	return &BlackipDao{
		engine: engine,
	}
}

//根据ID获取数据
func (d *BlackipDao) Get(id int) *models.LtBlackip {
	data := &models.LtBlackip{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

//获取所有数据
func (d *BlackipDao) GetAll() []*models.LtBlackip {
	dataList := make([]*models.LtBlackip, 0)
	_ := d.engine.Desc("id").Find(&dataList)
	return dataList
}

//获取总数
func (d *BlackipDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtBlackip{})
	if err != nil {
		return 0
	}
	return num
}

//软删除
//func(d *BlackipDao) Delete(id int) error {
//	data := models.LtBlackip{Id:id, SysStatus:1}
//	_, err := d.engine.Id(data.Id).Update(data)
//	return err
//}

//更新
func (d *BlackipDao) Update(data *models.LtBlackip, colums []string) error {
	_, err := d.engine.Id(data.Id).Cols(colums...).Update(data)
	return err
}

//创建
func (d *BlackipDao) Create(data *models.LtBlackip) error {
	_, err := d.engine.Insert(data)
	return err
}

//根据IP查询
func (d *BlackipDao) GetByIp(ip string) []*models.LtBlackip {
	dataList := make([]*models.LtBlackip, 0)
	err := d.engine.Where("ip = ?", ip).Desc("ip").Find(&dataList)
	if err != nil || len(dataList) < 1 {
		return nil
	}
	return dataList
}
