package models

type LtUser struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Username   string `xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	Realname   string `xorm:"not null default '' comment('真实姓名') VARCHAR(50)"`
	Mobile     string `xorm:"not null default '' comment('手机号码') VARCHAR(20)"`
	BlackTime  int    `xorm:"not null default 0 comment('黑名单限制到期时间') INT(11)"`
	Address    string `xorm:"not null default '' comment('联系地址') VARCHAR(255)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	SysUpdated int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	SysIp      string `xorm:"not null default '' comment('IP地址') VARCHAR(50)"`
}
