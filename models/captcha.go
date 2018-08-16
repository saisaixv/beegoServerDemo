package models

import(
	"androidServer/utils/captcha"
)

func CreateCaptcha()(id,url string){
	id=captcha.GenerateCaptcha()
	url=id+".png"
	return id,url
}

func VerifyCaptcha(id string,value string) bool {
	return captcha.VerifyCaptcha(id,value)
}