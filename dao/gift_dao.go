package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/xuanxiaox/lottery/comm"
	"github.com/xuanxiaox/lottery/models"
)

type GiftDao struct {
	engine *xorm.Engine
}

//创建GiftDao
func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{
		engine: engine,
	}
}

//根据ID获取数据
func (d *GiftDao) Get(id int) *models.LtGift {
	data := &models.LtGift{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

//获取所有数据
func (d *GiftDao) GetAll() []models.LtGift {
	dataList := make([]models.LtGift, 0)
	_ = d.engine.Asc("sys_status").Asc("prize_order").Find(&dataList)
	return dataList
}

//获取总数
func (d *GiftDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtGift{})
	if err != nil {
		return 0
	}
	return num
}

//软删除
func (d *GiftDao) Delete(id int) error {
	data := models.LtGift{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

//更新
func (d *GiftDao) Update(data *models.LtGift, colums []string) error {
	_, err := d.engine.Id(data.Id).Cols(colums...).Update(data)
	return err
}

//创建
func (d *GiftDao) Create(data *models.LtGift) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *GiftDao) GetAllUse() []models.LtGift {
	now := comm.NowUnix()
	dataList := make([]models.LtGift, 0)
	_ = d.engine.
		Where("sys_status=?", 0).
		Where("time_begin<=?", now).
		Where("time_end>?", now).
		Where("prize_num>?", 0).
		Desc("gtype").
		Asc("prize_order").
		Find(&dataList)
	return dataList
}

func (d *GiftDao) DecrLeftNum(giftId, num int) (int64, error) {
	row, err := d.engine.Id(giftId).
		Where("left_num>=?", num).
		Decr("left_num", num).
		Update(&models.LtGift{Id: giftId})
	return row, err
}

func (d *GiftDao) IncrLeftNum(giftId, num int) (int64, error) {
	row, err := d.engine.Id(giftId).
		Incr("left_num", num).
		Update(&models.LtGift{Id: giftId})
	return row, err
}
