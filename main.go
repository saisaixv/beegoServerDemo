package main

import (

	// "strconv"
	
	"androidServer/common"
	// "androidServer/utils"
	"androidServer/utils/db"
	"androidServer/utils/cache"
	
	_ "androidServer/routers"
	"github.com/astaxie/beego"
	// "github.com/go-ini/ini"
	clog "github.com/cihub/seelog"
)

func main() {
	setupDB()
	err:=setupRedis()
	if err!=nil{
		clog.Trace(err.Error())
	}

	setupHttps()

	beego.Run()
}


func setupDB() {
	
	url:=beego.AppConfig.String("db_android::url")
	max_lt,_:=beego.AppConfig.Int("db_android::max_life_time")
	max_oc,_:=beego.AppConfig.Int("db_android::max_open_conns")
	max_ic,_:=beego.AppConfig.Int("db_android::max_idle_conns")

	common.DBmysqlAndroid=db.OpenDB(url,max_lt,max_oc,max_ic)

	// clog.Trace("url: "+url+",max_lt: %d,max_oc: %d,max_ic: ",max_lt,max_oc,max_ic)

}

func setupRedis() error {
	url:=beego.AppConfig.String("redis_android::url")
	auth:=beego.AppConfig.String("redis_android::auth")
	max_idles,_:=beego.AppConfig.Int("redis_android::max_idles")
	idle_timeout,_:=beego.AppConfig.Int("redis_android::idle_timeout")

	cache.Init(url,auth,max_idles,idle_timeout,common.REDIS_DB_ANDROID)
	
	err:=cache.Ping()
	if err!=nil{
		return err
	}

	return nil
}

func setupHttps() error {
	//配置https
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.Graceful = true
	beego.BConfig.Listen.HTTPSAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPSPort = 8090
	beego.BConfig.Listen.HTTPSCertFile="D:\\goWorkSpace\\src\\androidServer\\CARoot1024.crt"
	beego.BConfig.Listen.HTTPSKeyFile ="D:\\goWorkSpace\\src\\androidServer\\CARoot1024.key"
	return nil
}