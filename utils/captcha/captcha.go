package captcha

import(

	"androidServer/utils"
	"androidServer/utils/cache"

	"github.com/dchest/captcha"
)

var(
	SR *StoreRedis
)

type StoreRedis struct {
	// captcha.Store
}

type ImgBytes struct {
	Img []byte
}

func (s *StoreRedis) Set(id string, digits []byte) {
	obj := new(ImgBytes)
	obj.Img = digits
	cache.DoSet(id, obj, utils.TIME_MINUTE_FIVE)
}

func (s *StoreRedis) Get(id string, clear bool) (digits []byte) {
	obj := new(ImgBytes)
	_ = cache.DoGet(id, obj)
	return obj.Img
}


func InitCaptcha() {
	SR = new(StoreRedis)
	captcha.SetCustomStore(SR)
}

func GenerateCaptcha() string{
	return captcha.NewLen(6)
}

func VerifyCaptcha(id string,value string) bool {
	return captcha.VerifyString(id,value)
}