package handler

import "auditIntegral/goSql"

func main() {
	conf := gosql.Config{
		Master: &gosql.ConfigModel{Host: "127.0.0.1", Port: 3306, User: "root", Password: "root", DBName: 100},
		Slave: []*gosql.ConfigModel{
			&gosql.ConfigModel{Host: "127.0.0.1", Port: 3306, User: "root", Password: "root", DBName: 100},
			&gosql.ConfigModel{Host: "127.0.0.1", Port: 3306, User: "root", Password: "root", DBName: 100},
		},
	}
}
