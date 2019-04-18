package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/xuanxiaox/lottery/models"
)

type ResultDao struct {
	engine *xorm.Engine
}

//创建ResultDao
func newResultDao(engine *xorm.Engine) *ResultDao {
	return &ResultDao{
		engine: engine,
	}
}

//根据ID获取数据
func (d *ResultDao) Get(id int) *models.LtResult {
	data := &models.LtResult{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

//获取所有数据
func (d *ResultDao) GetAll() []*models.LtResult {
	dataList := make([]*models.LtResult, 0)
	_ := d.engine.Desc("id").Find(&dataList)
	return dataList
}

//获取总数
func (d *ResultDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtResult{})
	if err != nil {
		return 0
	}
	return num
}

//软删除
//func(d *ResultDao) Delete(id int) error {
//	data := models.LtResult{Id:id, SysStatus:1}
//	_, err := d.engine.Id(data.Id).Update(data)
//	return err
//}

//更新
func (d *ResultDao) Update(data *models.LtResult, colums []string) error {
	_, err := d.engine.Id(data.Id).Cols(colums...).Update(data)
	return err
}

//创建
func (d *ResultDao) Create(data *models.LtResult) error {
	_, err := d.engine.Insert(data)
	return err
}

//根据IP查询
func (d *ResultDao) GetByIp(ip string) []*models.LtResult {
	dataList := make([]*models.LtResult, 0)
	err := d.engine.Where("ip = ?", ip).Desc("ip").Find(&dataList)
	if err != nil || len(dataList) < 1 {
		return nil
	}
	return dataList
}
