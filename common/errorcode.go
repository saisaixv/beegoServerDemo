package common

const(
	VERSION="1.0.0"
)

const(
	RSP_CODE_OK=200
	RSP_CODE_BAD_REQUEST=400
	RSP_CODE_UNAUTHORIZED=401
	RSP_CODE_FORBIDDEN=403
	RSP_CODE_NOT_FOUND=404
	RSP_CODE_INTERNAL_SERVER_ERROR=500


	OK=0
	ErrNotDefine=100	//未定义
	ErrAuthFailed=101	//授权失败
	ErrNickNameExist=102	//昵称已经存在
	ErrAccountNotExist=103	//账号不存在
	ErrPwdError=104	//密码错误
	ErrParamsError=105	//参数错误
	ErrUserNotExist=106	//用户不存在
	ErrLoginFailed=107	//登录失败
	ErrLogoutFailed=108	//登出失败
	ErrCaptchaVerifyError=109	//验证码确认失败
	ErrTooManyErrorOfLogin=110	//登录错误次数超过10次
	ErrChangePwdError=111	//修改密码失败
	ErrOldPwdError=112	//原始密码不正确
	

)