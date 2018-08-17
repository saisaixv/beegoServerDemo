package common

type BaseRsp struct {
	Error_code int `json:"error_code"`
}

type User struct {
	Id string `json:"id"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	
}

type UserInfo struct {
	User
	CreateTime string `json:"createtime"`
	UpdateTime string `json:"updatetime"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type RegisterReq struct {
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Credential string `json:"credential"`
}

type RegisterRsp struct {
	BaseRsp
	UserInfo UserInfo `json:"userinfo"`
}

type LoginReq struct {
	Identify_type string `json:"identify_type"`
	Identifier string `json:"identifier"`
	Credential string `json:"credential"`
	CaptchaId string `json:"captcha_id"`
	Value string `json:"value"`
}

type LoginRsp struct {
	BaseRsp
	Token string `json:"token"`
	ErrCount int `json:"err_count"`
	CaptchaId string `json:"captcha_id"`
	CaptchaUrl string `json:"captcha_url"`
}

type LogoutReq struct {
	
}

type LogoutRsp struct {
	BaseRsp
}

type UserInfoReq struct {
	
}

type UserInfoRsp struct {
	BaseRsp
	UserInfo UserInfo `json:"userinfo"`
}

type PutUserInfoReq struct {
	User_id string `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type PutUserInfoRsp struct {
	BaseRsp
}

type UserInfoLisReq struct {
	
}

type UserInfoLisRsp struct {
	BaseRsp
	UserList []User `json:"userlist"`
}

type UploadPicReq struct {
	
}

type UploadPicRsp struct {
	BaseRsp
	Url string `json:"url"`
}

type GetCaptchaReq struct {
	
}

type GetCaptchaRsp struct {
	BaseRsp
	Id string `json:"id"`
	Url string `json:"url"`
}

type VerifyCaptchaReq struct {
	Id string `json:"id"`
	Value string `json:"value"`
}

type VerifyCaptchaRsp struct {
	BaseRsp
}

type ChangePwdReq struct {
	OldPwd string `json:"old_password"`
	NewPwd string 	`json:"new_password"`
	VerifyPwd string `json:"verify_password"`
}

type ChangePwdRsp struct {
	BaseRsp
}

