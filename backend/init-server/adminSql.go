// 此包用于 初始化 管理后台数据库
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func initAdminSql(db *sql.DB) {
	createAdminInfoSql := `CREATE TABLE IF NOT EXISTS adminInfo(
			id int unsigned auto_increment primary key ,
			name varchar(255),
			password varchar(255),
			
# 			密保问题和答案
            securityProblem varchar(255),
            securityAnswer varchar(255)          
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
	INSERT INTO admininfo (name,password) values  ("admin","123456")
`
	_, err = db.Exec(insertAdminSql)
	if err != nil {
		return
	}
}
