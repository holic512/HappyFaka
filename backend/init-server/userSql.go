// 用于初始化 用户的 sql
package main

import "database/sql"

func initUserSql(db *sql.DB) {
	createUserSql := `CREATE TABLE IF NOT EXISTS user( 
    id int unsigned auto_increment primary key ,
    username varchar(255),
    password varchar(255),
    status bool,
    balance float,
    if_merchant bool,
    merId int,
    phone varchar(255),
    mail varchar(255),
    invite_code varchar(64),
    registration_time datetime,
    last_login_time datetime   
)`
	//执行 创建数据表
	_, err := db.Exec(createUserSql)
	if err != nil {
		return
	}

}
