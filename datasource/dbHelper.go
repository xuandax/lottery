package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/xuanxiaox/lottery/conf"
	"log"
	"sync"
)

var instanceDb *xorm.Engine
var mu sync.Mutex

func InstanceMaster() *xorm.Engine {
	if instanceDb != nil {
		return instanceDb
	}
	//第一个进来的上锁，创建实例
	mu.Lock()
	defer mu.Unlock()
	//防止大家排队的时候，解锁了再创建
	if instanceDb != nil {
		return instanceDb
	}

	return newInstanceDb()
}

func newInstanceDb() *xorm.Engine {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Pwd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database)
	engine, err := xorm.NewEngine(conf.DriverName, dataSourceName)
	if err != nil {
		log.Fatal("newInstanceDb err = ", err)
		return nil
	}
	return engine
}
