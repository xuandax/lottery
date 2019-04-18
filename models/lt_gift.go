package models

type LtGift struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Title      string `xorm:"not null default '' comment('奖品名称') VARCHAR(255)"`
	PrizeNum   int    `xorm:"not null default 0 comment('奖品数量') INT(10)"`
	LeftNum    int    `xorm:"not null default 0 comment('剩余奖品数量') INT(10)"`
	PrizeCode  string `xorm:"not null default '0' comment('0-9999标识100%中奖概率，0-0标识万分之一的中奖概率') VARCHAR(50)"`
	PrizeTime  int    `xorm:"not null default 0 comment('发放周期，以天为单位：D') INT(10)"`
	Img        string `xorm:"not null default '' comment('奖品图片') VARCHAR(255)"`
	PrizeOrder int    `xorm:"not null default 0 comment('奖品排序序号，小的排在前面') INT(10)"`
	Gtype      int    `xorm:"not null default 0 comment('奖品类型，0：虚拟币 1：实物小奖 2：实物大奖') INT(10)"`
	Gdata      string `xorm:"not null default '' comment('奖品扩展数据，比如:虚拟币数量') VARCHAR(255)"`
	TimeBegin  int    `xorm:"not null default 0 comment('抽奖活动开始时间') INT(10)"`
	TimeEnd    int    `xorm:"not null default 0 comment('抽奖活动结束时间') INT(10)"`
	PrizeData  string `xorm:"comment('发奖计划， [[时间1,数量1],[时间2,数量2]...]') MEDIUMTEXT"`
	PrizeBegin int    `xorm:"not null default 0 comment('发奖计划周期的开始') INT(10)"`
	PrizeEnd   int    `xorm:"not null default 0 comment('发奖计划周期的结束') INT(10)"`
	SysStatus  int    `xorm:"not null default 0 comment('状态 0：正常 1：删除') TINYINT(3)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	SysUpdated int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	SysIp      string `xorm:"not null default '' comment('操作人IP') VARCHAR(50)"`
}
