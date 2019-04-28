package routes

import (
	"github.com/kataras/iris/mvc"
	"github.com/xuanxiaox/lottery/bootstrap"
	"github.com/xuanxiaox/lottery/services"
	"github.com/xuanxiaox/lottery/web/controllers"
	"github.com/xuanxiaox/lottery/web/middleware"
)

func Configure(b *bootstrap.Bootstrapper) {
	blackipService := services.NewBlackipService()
	codeService := services.NewCodeService()
	giftService := services.NewGiftService()
	resultService := services.NewResultService()
	userdayService := services.NewUserdayService()
	userService := services.NewUserService()

	index := mvc.New(b.Party("/"))
	index.Register(userService,
		blackipService,
		codeService,
		giftService,
		resultService,
		userdayService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(userService,
		blackipService,
		codeService,
		giftService,
		resultService,
		userdayService)
	admin.Handle(new(controllers.AdminController))

	adminGift := admin.Party("/gift")
	adminGift.Register(giftService)
	adminGift.Handle(new(controllers.AdminGiftController))

	adminCode := admin.Party("/code")
	adminCode.Register(codeService)
	adminCode.Handle(new(controllers.AdminCodeController))

	adminResult := admin.Party("/result")
	adminResult.Register(resultService)
	adminResult.Handle(new(controllers.AdminResultController))

	adminUser := admin.Party("/user")
	adminUser.Register(userService)
	adminUser.Handle(new(controllers.AdminUserController))

	blackipUser := admin.Party("/blackip")
	blackipUser.Register(blackipService)
	blackipUser.Handle(new(controllers.AdminBlackipController))
}
