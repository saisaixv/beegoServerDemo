package controllers

import(
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"androidServer/utils"
	"androidServer/common"
	"androidServer/models"
	
	"github.com/astaxie/beego"
	clog "github.com/cihub/seelog"
	
)

type BaseController struct {
	beego.Controller
	Token string 
	Language string
	TimeZone int 
	AuthType int
	User_id string

}

func (this * BaseController)fetchAuthType()  {
	ul:=this.Ctx.Input.Header("x-us-authtype")
	if ul!=""{
		i,err:=strconv.Atoi(ul)
		if err==nil{
			this.AuthType=i
		}
	}
}

func (this *BaseController)fetchToken()  {
	token:=this.Ctx.Input.Header("x-us-token")
	this.Token=token
	if token!=""{
		ret:=strings.Split(token,utils.SPLIT_CHAR)
		if len(ret)==2{
			this.User_id=ret[0]
		}
	}
}

func (this *BaseController)fetchLanguage()  {
	this.Language=this.Ctx.Input.Header("accept-language")
}

func (this *BaseController)fetchTimeZone()  {
	timeZone:=this.Ctx.Input.Header("time-zone")
	if timeZone!=""{
		i,err:=strconv.Atoi(timeZone)
		if err==nil{
			this.TimeZone=i
		}else{
			this.TimeZone=common.CST_TIME_ZONE_ERROR
		}
	}else{
		this.TimeZone=common.CST_TIME_ZONE_ERROR
	}
}

func (this *BaseController)Prepare()  {
	this.fetchAuthType()
	this.fetchLanguage()
	this.fetchTimeZone()
	this.fetchToken()

	ret:=checkHttpHead(this.Language,this.TimeZone)

	if !ret{
		
		rsp:=new(common.BaseRsp)
		rsp.Error_code=common.ErrParamsError
		this.SetRspCode(common.RSP_CODE_BAD_REQUEST)
		this.Data["json"]=rsp
		this.ServeJSON()
		this.StopRun()
	}

	if this.Token!=""{
		ret,user_id:=models.Auth(this.AuthType,this.Token)

		if !ret{
			rsp:=new(common.BaseRsp)
			rsp.Error_code=common.ErrAuthFailed
			this.Data["json"]=rsp
			this.SetRspCode(common.RSP_CODE_UNAUTHORIZED)
			this.ServeJSON()
			this.StopRun()
		}else{
			this.User_id=user_id
		}
	}
	
}


func (this *BaseController)FetchBodyJsonToOBJ(v interface{}) error {
	r:=this.Ctx.Request
	defer r.Body.Close()

	body,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		clog.Error(err)
		return err
	}

	if len(body)==0{

		return nil
	}

	clog.Trace("[http req]: ",string(body)) 

	if err:=json.Unmarshal(body,v);err!=nil{
		clog.Error(err)
		return err
	}

	return nil

}

func checkHttpHead(language string,timeZone int) bool {
	if language=="" || timeZone==common.CST_TIME_ZONE_ERROR{
		return false
	}

	if strings.ToUpper(language) !="ZH_CN" &&
		strings.ToUpper(language) !="ZH_TW" &&
		strings.ToUpper(language) !="EN"{
		return false
	}
	
	return true
}

func (this *BaseController)SetRspCode(code int)  {
	this.Ctx.ResponseWriter.WriteHeader(code)
}
