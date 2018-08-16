package models

import(
	"strings"

	"androidServer/common"
	"androidServer/utils"
	"androidServer/utils/cache"
)

func Auth(authType int,urlToken string) (ret bool,user_id string) {
	
	ret,user_id=DoAuth(authType,urlToken)
	return
}

func DoAuth(authType int,urlToken string) (ret bool, user_id string) {

	if urlToken==""{
		return false,""
	}

	retSplit:=strings.Split(urlToken,utils.SPLIT_CHAR)
	if len(retSplit)!=2{
		return false,""
	}

	rediKey:=common.GetKeyToken(urlToken)
	retGet,token:=cache.DoStrGet(rediKey)
	if retGet == false || token ==""{
		return false,""
	}


	//redis中取出的字符串带双引号
	if token != urlToken{
		return false,""
	}else{
		retExpire:=cache.DoExpire(rediKey,common.KEY_TOKEN_EX)
		if retExpire==false{
			return false,""
		}
		return true,retSplit[0]
	}
	
	
}