// 此包用于 初始化 管理后台数据库
package main

import (
	"database/sql"
)

func initAdminSql(db *sql.DB) {
	createAdminInfoSql := `CREATE TABLE IF NOT EXISTS adminInfo(
			id int unsigned auto_increment primary key ,
			name varchar(255),
			password varchar(255),
			
# 			密保问题和答案
            securityProblem varchar(255),
            securityAnswer varchar(255) ,
            
            mail varchar(255),
            if_mail bool,
            phone varchar(255),
            if_phone  bool     
                                    )
			`

	//执行创建数据表
	_, err := db.Exec(createAdminInfoSql)
	if err != nil {
		return
	}

	//查询数据表 是否有内容 但不存在内容 执行初始化默认密码 admin 123456
	noSelectAdmin := `
	select COUNT(*) from admininfo
`
	var adminCount int
	err = db.QueryRow(noSelectAdmin).Scan(&adminCount)
	if err != nil {
		println("查询是否存在管理员账户失败")
		return
	}

	//当 计数器 不为 0 时 则直接推出 证明 有初始数据
	if adminCount != 0 {
		return
	}

	insertAdminSql := `
	INSERT INTO admininfo (name,password,if_mail,if_phone) values  ("admin","123456",false,false)
`
	_, err = db.Exec(insertAdminSql)
	if err != nil {
		return
	}
}
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
	inviter_id varchar(64),
    registration_time datetime,
    last_login_time datetime   
)`
	//执行 创建数据表
	_, err := db.Exec(createUserSql)
	if err != nil {
		return
	}

}
