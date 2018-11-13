package db

import (
	"auditIntegral/_goSql"
	"auditIntegral/_public/config"
)

//var (
//	db *gosql.DbMysql
//)

func Init() *gosql.DbMysql {
	conf := gosql.Config{
		Master: config.MySqlMaster,
		Slave:  config.MySqlSlave,
	}
	db := conf.NewDbMysql()
	return db
}
