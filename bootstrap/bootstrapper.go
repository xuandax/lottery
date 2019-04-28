package bootstrap

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/xuanxiaox/lottery/conf"
	"time"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName     string //应用名称
	AppOwner    string //应用拥有者
	AppSpawTime time.Time
}

//创建Bootstrapper
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application: iris.New(),
		AppName:     appName,
		AppOwner:    appOwner,
		AppSpawTime: time.Now(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}

//设置视图模板
func (b *Bootstrapper) SetViews(viewDir string) {
	viewEngine := iris.HTML(viewDir, ".html").Layout("share/layout.html")
	//设置是否每次都加载模板
	viewEngine.Reload(true)
	//增加时间格式化方法
	viewEngine.AddFunc("FormatUnixTime", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SystimeFormat)
	})
	viewEngine.AddFunc("FormatUnixTimeShort", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SystimeFormatShort)
	})
	b.RegisterView(viewEngine)
}

//设置异常处理
func (b *Bootstrapper) SetErrorHandlers() {
	b.OnAnyErrorCode(func(context iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  context.GetStatusCode(),
			"message": context.Values().GetString("message"),
		}

		if jsonOutput := context.URLParamExists("json"); jsonOutput {
			context.JSON(err)
			return
		}

		context.ViewData("Err", err)
		context.ViewData("Title", "Error")
		context.View("share/error.html")
	})
}

//配置初始化方法
func (b *Bootstrapper) Configure(cfgs ...Configurator) {
	for _, cfg := range cfgs {
		cfg(b)
	}
}

//设置定时任务
func (b *Bootstrapper) SetCron() {

}

const (
	//设置根目录
	StaticAssets = "./public/"
	//设置图片
	Favicon = "favicon.ico"
)

//启动Bootstrap,初始化应用
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	//设置模板目录
	b.SetViews("./views")
	b.SetErrorHandlers()
	b.SetCron()

	b.Favicon(StaticAssets + Favicon)
	b.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)
	b.Use(recover.New())
	b.Use(logger.New())
	return b
}

//启动监听
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
