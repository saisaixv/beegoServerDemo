package db

import (
	"fmt"
	"database/sql"
	"time"

	"androidServer/utils"

	_ "github.com/go-sql-driver/mysql"

	clog "github.com/cihub/seelog"

)

func OpenDB(url string,maxLT int,maxOC int,maxIC int) (DBmysql *sql.DB) {
	
	var err error

	DBmysql,err=sql.Open("mysql",url)
	if err!=nil{
		clog.Critical(err.Error())
		panic(err)
	}
	utils.CheckErr(err,utils.CHECK_FLAG_EXIT)

	DBmysql.SetConnMaxLifetime(time.Duration(maxLT)*time.Second)
	DBmysql.SetMaxOpenConns(maxOC)
	DBmysql.SetMaxIdleConns(maxIC)

	clog.Info("[db opened]")


	return DBmysql
}

func CloseDB(DBmysql *sql.DB){

	DBmysql.Close()
	clog.Info("[db close mysql]")
}

func DoQuery(DBmysql *sql.DB,sql string,args ...interface{})(results[][]string,err error){

	clog.Trace("[sql]: ",sql+" args:"+utils.Args2Str(args...))

	rows,err:=DBmysql.Query(sql,args...)
	utils.CheckErr(err,utils.CHECK_FLAG_LOGONLY)
	if  err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}

	cols,_:=rows.Columns()
	values:=make([][]byte,len(cols))
	scans:=make([]interface{},len(cols))
	for i:= range values{
		scans[i]=&values[i]
	}

	results=make([][]string,0)

	i:=0

	for rows.Next(){
		err=rows.Scan(scans...)
		utils.CheckErr(err,utils.CHECK_FLAG_LOGONLY)
		row:=make([]string,0)
		for _,v:=range values{
			row=append(row,string(v))
		}
		results=append(results,row)
		i++
	}
	rows.Close()

	return results,nil

}

func DoExec(DBmysql *sql.DB,sql string,args ...interface{})(bool,error)  {
	
	clog.Trace("[sql]: ",sql+" args:"+utils.Args2Str(args...))
	_,err:=DBmysql.Exec(sql,args...)
	utils.CheckErr(err,utils.CHECK_FLAG_LOGONLY)
	if err==nil{
		return true,err
	}else{
		return false,err
	}
}

func DoExecBatch(DBmysql *sql.DB,sqls []string,args [][]interface{}) (bool,error)  {

	tx,errBegin:=DBmysql.Begin()
	utils.CheckErr(errBegin,utils.CHECK_FLAG_LOGONLY)

	if errBegin!=nil{
		return false,errBegin
	}

	var errExec error
	for idx,sql:=range sqls{
		clog.Trace("[sql]: ",sql+" args:"+utils.Args2Str(args[idx]...))
		_,errExec =tx.Exec(sql,args[idx]...)

		if errExec!=nil{
			errRollbak:=tx.Rollback()
			utils.CheckErr(errRollbak,utils.CHECK_FLAG_LOGONLY)
			clog.Error("[sql]:"+sql)
			return false,errExec
		}
	}
	errCommit:=tx.Commit()
	utils.CheckErr(errCommit,utils.CHECK_FLAG_LOGONLY)
	if errCommit!=nil{
		errRollback:=tx.Rollback()
		utils.CheckErr(errRollback,utils.CHECK_FLAG_LOGONLY)
		return false,errCommit
	}

	return true,nil
	
}

func SqlColon(cnt int) string  {
	if cnt==0{
		return ""
	}

	sql:=""
	for i:=0;i<cnt;i++{
		sql=sql+`?,`
	}
	if utils.Substring(sql,len(sql)-1,len(sql))==","{
		sql=utils.Substring(sql,0,len(sql)-1)
	}

	fmt.Println("SqlColon:"+sql)
	return sql
}