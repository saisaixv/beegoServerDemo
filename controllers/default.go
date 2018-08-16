package controllers

import (
	// "fmt"
	// "io/ioutil"
	// "encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	clog "github.com/cihub/seelog"


	"androidServer/utils/captcha"
	"androidServer/common"
	"androidServer/models"
	"androidServer/utils"
	"androidServer/utils/cache"

	"strings"
)



type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

}


type RegisterController struct {
	BaseController
}


type LoginController struct {
	BaseController
}

type LogoutController struct {
	BaseController
}

type UserInfoController struct {
	BaseController
}

type PicController struct {
	BaseController
}

type CaptchaController struct {
	BaseController
}


func (this *RegisterController)Post()  {
	
	req:=new(common.RegisterReq)
	rsp:=new(common.RegisterRsp)
	
	defer func ()  {
		this.Data["json"]=rsp
		this.ServeJSON()
	}()

	rsp.Error_code=common.OK

	err:=this.FetchBodyJsonToOBJ(req)
	if err!=nil{

		rsp.Error_code=common.ErrParamsError
		return
	}

	err=models.Register(req,rsp)
	// err=models.Query()
	if err!=nil{
		clog.Error(err)
		return
	}




}


func (this *LoginController)Post()  {
	
	req:=new(common.LoginReq)
	rsp:=new(common.LoginRsp)

	defer func ()  {
		this.Data["json"]=rsp
		this.ServeJSON()
	}()


	err:=this.FetchBodyJsonToOBJ(req)

	if err!=nil{
		rsp.Error_code=common.ErrParamsError
		return
	}

	//先判断当前登录失败的次数
	lecKey:=common.GetKeyLoginErrCount(req.Identifier)
	ret,errCnt:=cache.DoStrGet(lecKey)
	count,err:=strconv.Atoi(errCnt)
	if err!=nil{
		count=0
	}
	if ret{
		
		//登录错误超过十次 ，限制 五分钟后允许再次登录
		if count >=10{
			rsp.Error_code=common.ErrTooManyErrorOfLogin
			return
		}

		//登录次数超过三次，必须携带验证码登录
		if count>=3{
			if req.CaptchaId=="" || req.Value==""{
				rsp.Error_code=common.ErrCaptchaVerifyError
				rsp.ErrCount=count
				return
			}

			b:=captcha.VerifyCaptcha(req.CaptchaId,req.Value)
			if !b{
				rsp.Error_code=common.ErrCaptchaVerifyError
				captcha_id,url:=models.CreateCaptcha()
				rsp.CaptchaId=captcha_id
				rsp.CaptchaUrl=url
				rsp.ErrCount=count
				return
			}
		}
	}

	err=models.Login(req,rsp)
	if err!=nil{
		rsp.Error_code=common.ErrLoginFailed
		return
	}

	//如果密码错误
	if rsp.Error_code==common.ErrPwdError{
		//设置当前登录次数，并更新过期时间
		count=count+1
		cache.DoStrSet(lecKey,strconv.Itoa(count),common.KEY_LOGIN_ERR_CNT_EX)

		rsp.Error_code=common.ErrPwdError
		rsp.ErrCount=count
		if count>=3{
			captcha_id,captcha_url:=models.CreateCaptcha()
			rsp.CaptchaId=captcha_id
			rsp.CaptchaUrl=captcha_url
		}
		return
	}

	cache.DoDel(lecKey)

	keyToken:=common.GetKeyToken(rsp.Token)

	//缓存token
	retToken:=cache.DoStrSet(keyToken,rsp.Token,common.KEY_TOKEN_EX)
	if retToken==false{
		rsp.Error_code=common.ErrLoginFailed
	}
	
}

func (this *UserInfoController)Get()  {
	// req:=new(common.UserInfoReq)
	rsp:=new(common.UserInfoRsp)

	defer func ()  {
		this.Data["json"]=rsp
		this.ServeJSON()
	}()

	err:=models.GetUserInfo(this.User_id,rsp)
	if err!=nil{
		clog.Trace(err.Error())
	}
}

func (this *UserInfoController)Put()  {
	req:=new(common.PutUserInfoReq)
	rsp:=new(common.PutUserInfoRsp)

	defer func ()  {
		this.Data["json"]=rsp
		this.ServeJSON()
	}()

	err:=this.FetchBodyJsonToOBJ(req)
	if err!=nil{
		rsp.Error_code=common.ErrParamsError
		return
	}

	ret,err:=models.UpdateUserInfo(req)
	if err!=nil{
		rsp.Error_code=common.ErrParamsError
		return
	}
	if !ret{
		rsp.Error_code=common.ErrParamsError
		return
	}
	rsp.Error_code=common.OK
}

func (this *LogoutController)Post()  {
	
	// req:=new(common.LogoutReq)
	rsp:=new(common.LogoutRsp)

	defer func ()  {
		this.Data["json"]=rsp
		this.ServeJSON()
	}()

	rsp.Error_code=common.OK

	ret:=models.Logout(this.Token)

	if !ret{
		rsp.Error_code=common.ErrLogoutFailed
	}


}

func (this *PicController)Post()  {

	rsp:=new(common.UploadPicRsp)

	defer func ()  {
		this.Data["json"]=rsp
		this.ServeJSON()
	}()
	//获取上传的文件
	f,h,err:=this.GetFile("myfile")
	if err!=nil{
		rsp.Error_code=common.ErrParamsError
		return
	}

	//文件目录
	split:=strings.Split(h.Filename,".")
	t:=time.Now()
	
	filepath:=common.LOCAL_DIR_PIC_PATH+split[0]+utils.SPLIT_CHAR+strconv.FormatInt(t.UnixNano(),10)+"."+split[len(split)-1]

	f.Close()
	this.SaveToFile("myfile",filepath)
	
	url:=filepath

	rsp.Url=url
	
}

func (this * CaptchaController)GetCaptcha(){

	rsp:=new(common.GetCaptchaRsp)

	defer func ()  {
		this.Data["json"]=rsp
		this.ServeJSON()
	}()

	rsp.Error_code=common.OK

	//获取验证码图片
	captcha_id,url:=models.CreateCaptcha()
	rsp.Id=captcha_id
	rsp.Url=url
}

func (this *CaptchaController)VerifyCaptcha()  {
	req:=new(common.VerifyCaptchaReq)
	rsp:=new(common.VerifyCaptchaRsp)

	defer func ()  {
		this.Data["json"]=rsp
		this.ServeJSON()
	}()

	rsp.Error_code=common.OK

	err:=this.FetchBodyJsonToOBJ(req)

	if err!=nil{
		rsp.Error_code=common.ErrParamsError
		return
	}

	b:=models.VerifyCaptcha(req.Id,req.Value)

	if !b{
		rsp.Error_code=common.ErrCaptchaVerifyError
		return
	}
}




