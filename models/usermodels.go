package models

import(

	// "strings"

	"androidServer/common"
	"androidServer/utils/db"
	"androidServer/utils"
	"androidServer/utils/cache"

	// clog "github.com/cihub/seelog"
)

func Login(req *common.LoginReq,rsp *common.LoginRsp) error {
	
	sql:="select user_id,identify_type,identifier,credential from user_auths where identify_type = ? and identifier = ?;"
	ret,err:=db.DoQuery(common.DBmysqlAndroid,sql,req.Identify_type,req.Identifier)
	if err!=nil{
		return err
	}

	if len(ret)==0{
		rsp.Error_code=common.ErrAccountNotExist
		return nil
	}

	credential:=ret[0][3]

	if req.Credential!=credential{
		rsp.Error_code=common.ErrPwdError
		return nil
	}

	rsp.Error_code=common.OK
	rsp.Token=ret[0][0]+utils.SPLIT_CHAR+utils.GetToken()
	
	return nil
}

func Logout(token string) bool {
	
	keyToken:=common.GetKeyToken(token)
	if keyToken==""{
		return false
	}

	retToken:=cache.DoDel(keyToken)
	if retToken==false{
		return false
	}

	return true
}

func GetUserInfo(user_id string,rsp *common.UserInfoRsp) error {
	
	sql:="select nickname,avatar from users where id = ?;"
	ret,err:=db.DoQuery(common.DBmysqlAndroid,sql,user_id)
	if err!=nil{
		rsp.Error_code=common.ErrParamsError
		return err
	}

	if len(ret)==0{
		rsp.Error_code=common.ErrUserNotExist
		return nil
	}

	rsp.Error_code=common.OK
	rsp.Userinfo.Id=user_id
	rsp.Userinfo.Nickname=ret[0][0]
	rsp.Userinfo.Avatar=ret[0][1]

	// b,str:=cache.DoStrGet("a")
	// if b==false{
	// 	clog.Trace("没有 key")
	// }else{
	// 	clog.Trace(str)
	// }
	// b=cache.DoStrSet("a","aaaaaaaaaa",100)
	// if b==false{
	// 	clog.Trace("set 失败")
	// }else{
	// 	clog.Trace("set 成功")
	// }
	// b,str=cache.DoStrGet("a")
	// if b==false{
	// 	clog.Trace("没有 key")
	// }else{
	// 	clog.Trace(str)
	// }

	// cache.DoGetExpire("a")
	return nil
}


func UpdateUserInfo(req *common.PutUserInfoReq) (bool,error) {
	
	sql:="update users set nickname=?,avatar=? where id=?;"

	return db.DoExec(common.DBmysqlAndroid,sql,req.Nickname,req.Avatar,req.User_id)
}