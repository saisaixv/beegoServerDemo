package routers

import (
	"androidServer/controllers"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
	"github.com/dchest/captcha"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/userinfo", &controllers.UserInfoController{})

	//正则路由
	beego.Router("/userinfo/:id/name", &controllers.UserInfoController{},"get:GetUserNameById")

	beego.Router("/pic",&controllers.PicController{})
	//自定义路由
	beego.Router("/captcha",&controllers.CaptchaController{},"get:GetCaptcha")
	beego.Router("/captcha",&controllers.CaptchaController{},"post:VerifyCaptcha")

	beego.Router("/news",&controllers.NewsController{})
	
	beego.Handler("/captcha/*.png", captcha.Server(150, 60))
	// beego.Get("/",func(ctx *context.Context){
	// 		ctx.Output.Body([]byte("Hello world"))
	// })

	// beego.Get("/login",func(ctx *context.Context){
	// 		ctx.Output.Body([]byte("Login"))
	// })
}
