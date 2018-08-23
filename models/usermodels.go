package models

import(

	"fmt"
	// "strings"
	"strconv"

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
	
	sql:="select nickname,avatar,createtime from users where id = ?;"
	ret1,err:=db.DoQuery(common.DBmysqlAndroid,sql,user_id)
	if err!=nil{
		rsp.Error_code=common.ErrParamsError
		return err
	}

	if len(ret1)==0{
		rsp.Error_code=common.ErrUserNotExist
		return nil
	}

	
	ret2,err:=db.DoQuery(common.DBmysqlAndroid,
		"select identify_type,identifier from user_auths where user_id = ?;",user_id)

	if err!=nil{
		rsp.Error_code=common.ErrParamsError
		return err
	}

	rsp.Error_code=common.OK
	rsp.UserInfo.Id=user_id
	rsp.UserInfo.Nickname=ret1[0][0]
	rsp.UserInfo.Avatar=ret1[0][1]
	rsp.UserInfo.CreateTime=ret1[0][2]

	for _,auth:= range ret2{
		if auth[0]=="email"{
			rsp.UserInfo.Email=auth[1]
		}else if auth[0]=="phone"{
			rsp.UserInfo.Phone=auth[1]
		}
	}

	return nil
}


func UpdateUserInfo(req *common.PutUserInfoReq) (bool,error) {
	
	sql:="update users set nickname=?,avatar=? where id=?;"

	return db.DoExec(common.DBmysqlAndroid,sql,req.Nickname,req.Avatar,req.User_id)
}

func GetUserInfoList(pageNum int,pageSize int,sex int,rsp *common.UserInfoLisRsp) error {


	// var userList2 []common.User

	fmt.Println(sex)

	if pageNum!=0{
	
		if sex==-1{
			sql:="select id,nickname,avatar,sex,createtime from users limit ?,?;"

			ret,err:=db.DoQuery(common.DBmysqlAndroid,sql,(pageNum-1)*pageSize,pageSize)
			if err!=nil{
				return err
			}
			// clog.Trace(ret)
	
			for _,u := range ret{
	
				userInfo:=common.UserInfo{}
				userInfo.Id=u[0]
				userInfo.Nickname=u[1]
				userInfo.Avatar=u[2]
				userInfo.Sex,_=strconv.Atoi(u[3])
				userInfo.CreateTime=u[4]

				sql="select identify_type,identifier from user_auths where user_id=?;"
				ret2,err:=db.DoQuery(common.DBmysqlAndroid,sql,userInfo.Id)
				if err!=nil{
					return err
				}
				for _,u2 := range ret2{
				
					if u2[0]=="phone"{
						userInfo.Phone=u2[1]
					}else if u2[0]=="email"{
						userInfo.Email=u2[1]
					}
				}


				rsp.UserInfoList=append(rsp.UserInfoList,userInfo)
			}
	
		}else{
			sql:="select id,nickname,avatar,sex,createtime from users where sex = ? limit ?,?;"

			ret,err:=db.DoQuery(common.DBmysqlAndroid,sql,sex,(pageNum-1)*pageSize,pageSize)
			if err!=nil{
				return err
			}
			// clog.Trace(ret)
	
			for _,u := range ret{
	
				userInfo:=common.UserInfo{}
				userInfo.Id=u[0]
				userInfo.Nickname=u[1]
				userInfo.Avatar=u[2]
				userInfo.Sex,_=strconv.Atoi(u[3])
				userInfo.CreateTime=u[4]

				sql="select identify_type,identifier from user_auths where user_id=?;"
				ret2,err:=db.DoQuery(common.DBmysqlAndroid,sql,userInfo.Id)
				if err!=nil{
					return err
				}
				for _,u2 := range ret2{
				
					if u2[0]=="phone"{
						userInfo.Phone=u2[1]
					}else if u2[0]=="email"{
						userInfo.Email=u2[1]
					}
				}


				rsp.UserInfoList=append(rsp.UserInfoList,userInfo)
			}
	
		}
		
		// clog.Trace(rsp.List)
		
	}else{

		if sex==-1{
			sql:="select id,nickname,avatar,sex,createtime from users;"

			ret,err:=db.DoQuery(common.DBmysqlAndroid,sql)
			if err!=nil{
				return err
			}
			// clog.Trace(ret)
	
			for _,u := range ret{
	
				userInfo:=common.UserInfo{}
				userInfo.Id=u[0]
				userInfo.Nickname=u[1]
				userInfo.Avatar=u[2]
				userInfo.Sex,_=strconv.Atoi(u[3])
				userInfo.CreateTime=u[4]

				sql="select identify_type,identifier from user_auths where user_id=?;"
				ret2,err:=db.DoQuery(common.DBmysqlAndroid,sql,userInfo.Id)
				if err!=nil{
					return err
				}
				for _,u2 := range ret2{
				
					if u2[0]=="phone"{
						userInfo.Phone=u2[1]
					}else if u2[0]=="email"{
						userInfo.Email=u2[1]
					}
				}
				rsp.UserInfoList=append(rsp.UserInfoList,userInfo)
	
			}
		}else{
			sql:="select id,nickname,avatar,sex,createtime from users where sex = ?;"

			ret,err:=db.DoQuery(common.DBmysqlAndroid,sql,sex)
			if err!=nil{
				return err
			}
			// clog.Trace(ret)
	
			for _,u := range ret{
	
				userInfo:=common.UserInfo{}
				userInfo.Id=u[0]
				userInfo.Nickname=u[1]
				userInfo.Avatar=u[2]
				userInfo.Sex,_=strconv.Atoi(u[3])
				userInfo.CreateTime=u[4]

				sql="select identify_type,identifier from user_auths where user_id=?;"
				ret2,err:=db.DoQuery(common.DBmysqlAndroid,sql,userInfo.Id)
				if err!=nil{
					return err
				}
				for _,u2 := range ret2{
				
					if u2[0]=="phone"{
						userInfo.Phone=u2[1]
					}else if u2[0]=="email"{
						userInfo.Email=u2[1]
					}
				}


				rsp.UserInfoList=append(rsp.UserInfoList,userInfo)
	
			}
		}
		
		// clog.Trace(rsp.List)
	}

	return nil
}


func ChangePwd(user_id string,req *common.ChangePwdReq,rsp *common.ChangePwdRsp) (bool,error) {

	if req.NewPwd =="" || req.OldPwd =="" || req.VerifyPwd==""{
		rsp.Error_code=common.ErrParamsError
		return false,nil
	}

	if req.NewPwd != req.VerifyPwd{
		rsp.Error_code=common.ErrParamsError
		return false,nil
	}

	if req.NewPwd == req.OldPwd{
		rsp.Error_code=common.ErrParamsError
		return false,nil
	}

	sql:="select identify_type,identifier,credential from user_auths where user_id = ?;"
	ret,err:=db.DoQuery(common.DBmysqlAndroid,sql,user_id)
	if err!=nil{
		rsp.Error_code=common.ErrChangePwdError
		return false,err
	}

	if len(ret)==0{
		rsp.Error_code=common.ErrUserNotExist
		return false,nil
	}

	for _,pwd := range ret{
		if pwd[2]!=req.OldPwd{
			rsp.Error_code=common.ErrOldPwdError
			return false,nil
		}

		pwd[2]=req.NewPwd
	}

	sql="update user_auths set credential = ? where user_id = ?;"
	b,err:=db.DoExec(common.DBmysqlAndroid,sql,req.NewPwd,user_id)

	if err!=nil{
		rsp.Error_code=common.ErrChangePwdError
		return false,err
	}

	if !b{
		rsp.Error_code=common.ErrChangePwdError
		return false,err
	}

	return true,nil
	
}