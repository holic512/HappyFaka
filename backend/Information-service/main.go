package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

func mysqlConnect() (*sql.DB, error) {
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
func main() {
	//用于用户页面 端口8082
	info := gin.Default()

	// 配置CORS中间件，允许所有来源跨域访问
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	info.Use(cors.New(config))

	//创建数据库对象
	db, err := mysqlConnect()
	if err != nil {
		return
	}

	//基础信息查询
	var name, message string
	selectTitleMessage := `select name,autoMessage from site_info where id = 1`
	info.GET("/api/user/message", func(c *gin.Context) {
		err := db.QueryRow(selectTitleMessage).Scan(&name, &message)
		if err != nil {
			c.JSON(http.StatusNotFound, "错误")
		}

		response := gin.H{
			"name":    name,
			"message": message,
		}
		c.JSON(http.StatusOK, response)
	})

	//标题查询
	var title string
	selectTitle := `select title from site_info where id = 1`
	info.GET("/api/user/title", func(c *gin.Context) {
		err := db.QueryRow(selectTitle).Scan(&title)
		if err != nil {
			c.JSON(http.StatusNotFound, "错误")
		}

		response := gin.H{
			"title": title,
		}
		c.JSON(http.StatusOK, response)
	})

	//启动服务
	info.Run(":8082")
}
