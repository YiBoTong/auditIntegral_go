package config

import "auditIntegral/_goSql"

// 数据库名称
const DBName = "auditIntegral"

// mysql主配置
var MySqlMaster = &gosql.ConfigModel{
	Host:          "127.0.0.1",
	Port:          3310,
	User:          "root",
	Password:      "root",
	DBName:        DBName,
	AutoCloseTime: 30,
	MaxOpenConns:  10,
	MaxIdleConns:  100,
}

// mysql从配置
var MySqlSlave = []*gosql.ConfigModel{
	&gosql.ConfigModel{
		Host:          "127.0.0.1",
		Port:          3311,
		User:          "root",
		Password:      "root",
		DBName:        DBName,
		AutoCloseTime: 30,
		MaxOpenConns:  10,
		MaxIdleConns:  100,
	},
}
