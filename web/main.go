package main

import (
	"fmt"
	"github.com/xuanxiaox/lottery/bootstrap"
	"github.com/xuanxiaox/lottery/web/routes"
)

var port = 8888

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Go抽奖系统", "xuanxiaox")
	app.Bootstrap()
	app.Configure(routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
