package routers

import (
	"androidServer/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/dchest/captcha"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowAllOrigins:  true,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin","x-us-authtype", "time-zone","accept-language","x-us-token","Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        AllowCredentials: true,
    }))


	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/userinfo", &controllers.UserInfoController{})
	beego.Router("/userinfolist", &controllers.UserInfoController{},"get:GetUserinfoList")
	beego.Router("/changepwd",&controllers.ChangePwdController{},"post:ChangePwd")

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
