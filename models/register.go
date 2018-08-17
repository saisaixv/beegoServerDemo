package models

import (

	"fmt"
	"time"
	// "strconv"

	"androidServer/utils/db"
	"androidServer/common"
	"androidServer/utils"

	

)

func checkErr(err error)  {
	if err!=nil{
		panic(err)
	}
}


func Register( req *common.RegisterReq,rsp *common.RegisterRsp) (error,) {
	
	sql:="select * from users where nickname = ?;"

	//查询昵称是否存在
	ret,err:=db.DoQuery(common.DBmysqlAndroid,sql,req.Nickname)
	
	if(err!=nil){
		rsp.Error_code=-1
		return err
	}

	//该nickname已经存在
	if len(ret)!=0{
		rsp.Error_code=common.ErrNickNameExist

		return nil
	}

	//创建新的 用户插入表中
	sql="insert into users (id,nickname,avatar,createtime) values (?,?,?,?);"
	user_id:=utils.GetMongoObjectId()

	now:=time.Now()

	createtime:=fmt.Sprintf("%d-%d-%d %d:%d:%d",
		now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	

	_,err=db.DoExec(common.DBmysqlAndroid,
		sql,user_id,req.Nickname,req.Avatar,createtime)

	if(err!=nil){
		rsp.Error_code=-1
		return err
	}

	
	sql="insert into user_auths (user_id,identify_type,identifier,credential,registertime) values (?,?,?,?,?);"

	//用email注册
	if req.Email!=""{
		_,err=db.DoExec(common.DBmysqlAndroid,
			sql,user_id,"email",req.Email,req.Credential,createtime)
	
		if(err!=nil){
			rsp.Error_code=-1
			//注册失败，删除users表中的数据
			db.DoExec(common.DBmysqlAndroid,"delete from users where id=?",user_id)
			return err
		}
	}

	//用phone注册
	if req.Phone!=""{
		_,err=db.DoExec(common.DBmysqlAndroid,
			sql,user_id,"phone",req.Phone,req.Credential,createtime)
	
		if(err!=nil){
			rsp.Error_code=-1
			//注册失败，删除users表中的数据
			db.DoExec(common.DBmysqlAndroid,"delete from users where id=?",user_id)
			return err
		}
	}
	

	rsp.Error_code=common.OK
	rsp.UserInfo.Id=user_id
	rsp.UserInfo.Nickname=req.Nickname
	rsp.UserInfo.Avatar=req.Avatar
	rsp.UserInfo.CreateTime=createtime
	rsp.UserInfo.Phone=req.Phone
	rsp.UserInfo.Email=req.Email

	return nil
}

func Query() error {
	sql:="select * from User_t;"
	results,err:=db.DoQuery(common.DBmysqlAndroid,sql)
	if err!=nil{
		return err
	}

	fmt.Println(results)
	return nil
}