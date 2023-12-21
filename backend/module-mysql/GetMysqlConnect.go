package mysql_module

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// GetMysqlConnect 此代码用于获取 数据库对象
func GetMysqlConnect() (*sql.DB, error) {
	file, err := os.ReadFile("../config/mysql-config/mysql.json")
	if err != nil {
		fmt.Println("无法读取mysql.json", err)
		return nil, err
	}

	var mysql map[string]string
	err = json.Unmarshal(file, &mysql)
	if err != nil {
		println("解码mysql-json失败", err)
		return nil, err
	}

	//数据库操作
	db, sqlErr := sql.Open("mysql", mysql["mysql"])
	if sqlErr != nil {
		panic(sqlErr.Error()) // 无法确定数据库连接状态，仅初始化了数据库连接池
	} else {
		fmt.Println("数据库连接池已初始化")
	}

	// 尝试连接数据库
	if err := db.Ping(); err != nil {
		return nil, err
	} else {
		fmt.Println("成功连接到数据库")
	}
	return db, nil
}
