package cache

import(

	"time"
	"encoding/json"

	"androidServer/utils"

	clog "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
)

var(
	pool *redis.Pool
)

func newPool(url string,max_idle,idle_timeout int,db int)  *redis.Pool{
	timeout:=time.Duration(idle_timeout)*time.Second
	return &redis.Pool{
		MaxIdle:max_idle,
		IdleTimeout:timeout,
		Dial:func () (redis.Conn,error) {
			c,err:=redis.DialURL(url)
			if err!=nil{
				return nil,err
			}
			c.Do("SELECT",db)
			return c,err
		},
		TestOnBorrow:func (c redis.Conn,t time.Time) error {
			if time.Since(t)<time.Minute{
				return nil
			}
			_,err:=c.Do("PING")
			return err
		},
	}
}

func Init(url string,max_idle,idle_timeout int,db int)  {
	pool=newPool(url,max_idle,idle_timeout,db)
}

func  Get() redis.Conn {
	if pool == nil{
		clog.Critical("Please set cache pool first!")	
		return nil
	}
	return pool.Get()
}

func Ping() error {
	conn:=Get()
	defer conn.Close()
	_,err:=conn.Do("PING")
	return err
}

func Close(){
	clog.Info("[redis pool close]")
	pool.Close()
}

func DoExpire(key string,expire int) bool {
	redisConn:=Get()
	defer redisConn.Close()

	ret,err:=redisConn.Do("EXPIRE",key,expire)
	if err!=nil{
		return false
	}

	value,_:=redis.Int(ret,err)
	if value==1{
		return true
	}else {
		return false
	}

}

func DoGetExpire(key string) (error) {
	
	redisConn:=Get()
	defer redisConn.Close()

	t,err:=redisConn.Do("TTL",key)
	clog.Trace(t)
	return err
}

func DoStrSet(key string,obj string,expire int) bool {
	redisConn:=Get()
	defer redisConn.Close()

	//存入redis
	ret,err:=redisConn.Do("SETEX",key,expire,obj)
	utils.CheckErr(err,utils.CHECK_FLAG_LOGONLY)
	if ret==false{
		return false
	}else{
		return true
	}
}

func DoStrGet(key string) (ret bool,obj string) {
	redisConn:=Get()
	defer redisConn.Close()

	retGet,err1:=redisConn.Do("get",key)
	if retGet==nil{
		return false,""
	}

	if err1==nil{
		value,err2:=redis.String(retGet,err1)
		utils.CheckErr(err2,utils.CHECK_FLAG_LOGONLY)
		if err2==nil{
			return true,value
		}else{
			return false,""
		}
	}else{
		return false,""
	}
}

func DoSet(key string,obj interface{},expire int) bool {
	redisConn:=Get()
	defer redisConn.Close()

	value,_:=json.Marshal(obj)
	_,err:=redisConn.Do("SETEX",key,expire,value)
	utils.CheckErr(err,utils.CHECK_FLAG_LOGONLY)
	if err!=nil{
		return false
	}
	return true
}

func DoGet(key string,obj interface{}) bool {
	redisConn:=Get()
	defer redisConn.Close()

	ret,err1:=redisConn.Do("GET",key)
	if ret==nil{
		return false
	}

	value,err2:=redis.Bytes(ret,err1)
	utils.CheckErr(err2,utils.CHECK_FLAG_LOGONLY)
	if err2!=nil{
		return false
	}

	err3:=json.Unmarshal(value,obj)
	utils.CheckErr(err3,utils.CHECK_FLAG_LOGONLY)
	if err3!=nil{
		return false
	}
	return true
}


func DoDel(key string) bool {
	redisConn:=Get()
	defer redisConn.Close()

	_,err :=redisConn.Do("DEL",key)
	utils.CheckErr(err,utils.CHECK_FLAG_LOGONLY)
	if err!=nil{
		return false
	}
	return true
}
