package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/xuanxiaox/lottery/models"
)

type CodeDao struct {
	engine *xorm.Engine
}

//创建CodeDao
func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{
		engine: engine,
	}
}

//根据ID获取数据
func (d *CodeDao) Get(id int) *models.LtCode {
	data := &models.LtCode{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

//获取所有数据
func (d *CodeDao) GetAll(page, size int) []models.LtCode {
	dataList := make([]models.LtCode, 0)
	start := (page - 1) * size
	_ = d.engine.Asc("id").Limit(size, start).Find(&dataList)
	return dataList
}

//获取总数
func (d *CodeDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtCode{})
	if err != nil {
		return 0
	}
	return num
}

//软删除
func (d *CodeDao) Delete(id int) error {
	data := models.LtCode{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

//更新
func (d *CodeDao) Update(data *models.LtCode, colums []string) error {
	_, err := d.engine.Id(data.Id).Cols(colums...).Update(data)
	return err
}

//创建
func (d *CodeDao) Create(data *models.LtCode) error {
	_, err := d.engine.Insert(data)
	return err
}

//根据gift_id获取数据
func (d *CodeDao) GetByGiftId(giftId int) []models.LtCode {
	dataList := make([]models.LtCode, 0)
	_ = d.engine.Where("gift_id=?", giftId).Desc("id").Find(&dataList)
	return dataList
}

func (d *CodeDao) CountByGiftId(giftId int) int64 {
	count, err := d.engine.Where("gift_id=?", giftId).Count()
	if err != nil {
		return 0
	}
	return count
}
