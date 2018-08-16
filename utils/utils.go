package utils

import(
	"runtime"
	"strconv"
	"fmt"
	"crypto/sha1"

	clog "github.com/cihub/seelog"
	"gopkg.in/mgo.v2/bson"
	"github.com/satori/go.uuid"

)

const(
	CHECK_FLAG_EXIT=1
	CHECK_FLAG_LOGONLY=2

	// 分隔符
	SPLIT_CHAR = "_"

	TIME_MINUTE_ONE=60
	TIME_MINUTE_FIVE=5*60
	TIME_MINUTE_TEN=10*60
	TIME_HOUR_ONE=60*60
	TIME_HOUR_TWO=2*60*60
	TIME_DAY_ONE=24*60*60
	
)

func CheckErr(err error,flag int)  {

	var path string

	if err!=nil{
		_,file,line,_:=runtime.Caller(1)
		path=" -- "+file+":"+strconv.Itoa(line)

		switch flag {
		case CHECK_FLAG_EXIT:
			clog.Critical(err.Error()+path)
			clog.Critical(StackTrace(false))
			panic(err)
		case CHECK_FLAG_LOGONLY:
			clog.Critical(err.Error()+path)
			clog.Critical(StackTrace(false))
		default:
			clog.Info(err.Error()+path)
		}
	}
	
}

func StackTrace(all bool) string {
	buf:=make([]byte,10240)

	for{
		size:=runtime.Stack(buf,all)

		if size==len(buf){
			buf=make([]byte,len(buf)<<1)
			continue
		}
		break
	}
	return string(buf)
}

func Args2Str(args ...interface{})(ret string){

	split:=","
	for _,v:=range args{
		switch v.(type) {
		case string:
			ret=ret+v.(string)+split
		case int:
			ret=ret+strconv.Itoa(v.(int))+split
		case int64:
			ret =ret+strconv.FormatInt(v.(int64),10)+split
		default:
			ret=ret+fmt.Sprintf("%T",v)+split
			
		}
	}
	return
}


func Substring(source string,start int,end int) string  {
	
	var r=[]rune(source)
	length:=len(r)

	if start<0||end>length||start>end{

		return ""
	}

	if start ==0&& end ==length{
		return source

	}

	return string(r[start:end])
}

func GetMongoObjectId() string {
	ret:=bson.NewObjectId()
	return ret.Hex()
}

func GetToken() string {
	ul,_:=uuid.NewV4()
	return Sha1(ul.String())
}

func Sha1(str string) string {
	sum:=sha1.Sum([]byte(str))
	return fmt.Sprintf("%x",sum)
}