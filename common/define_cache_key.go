package common

import(

	"androidServer/utils"

	// clog "github.com/cihub/seelog"
)

//缓存的key定义
const(
	PREFIX="android_"

	KEY_TOKEN=PREFIX+"tk_"//token
	KEY_LOGIN_ERR_CNT=PREFIX+"lec_"//登录失败次数
)

//缓存超时时间设置
const(

	KEY_TOKEN_EX=utils.TIME_MINUTE_TEN

	KEY_LOGIN_ERR_CNT_EX=utils.TIME_MINUTE_FIVE
)

func GetKeyToken(uid string) string {
	key:=KEY_TOKEN+uid
	return key
}

func GetKeyLoginErrCount(account string) string {
	
	return KEY_LOGIN_ERR_CNT+account
}