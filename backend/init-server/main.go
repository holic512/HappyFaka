package main

import "mysql-module"

func main() {
	db, err := mysql_module.GetMysqlConnect()
	if err != nil {
		println(err.Error())
		return
	}
	//初始化 管理后台 数据表
	initAdminSql(db)

}
