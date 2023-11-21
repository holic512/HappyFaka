// 设置服务
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
	//管理员后台网站配置 端口8081

	setup := gin.Default()

	// 配置CORS中间件，允许所有来源跨域访问
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	setup.Use(cors.New(config))

	//创建数据库对象
	db, err := mysqlConnect()
	if err != nil {
		return
	}

	//进行数据库内容初始化
	basicSetupMysql(db)

	//检索基础设置
	type site_info struct {
		Name        string `json:"name"`
		Title       string `json:"title"`
		Seo         string `json:"seo"`
		Description string `json:"description"`
		AutoMessage string `json:"autoMessage"`
	}
	var nowSiteInfo site_info
	selectInfo := `
	select name,title,seo,description,autoMessage from site_info
`
	row := db.QueryRow(selectInfo)
	if err := row.Scan(&nowSiteInfo.Name, &nowSiteInfo.Title, &nowSiteInfo.Seo, &nowSiteInfo.Description, &nowSiteInfo.AutoMessage); err != nil {
		println("检索基础设置出现问题")
		return
	}
	//检索基础设置

	//向前端发送基础数据内容
	setup.GET("/api/admin/getInfo", func(c *gin.Context) {

		// 构建JSON响应
		response := gin.H{
			"name":        nowSiteInfo.Name,
			"title":       nowSiteInfo.Title,
			"seo":         nowSiteInfo.Seo,
			"description": nowSiteInfo.Description,
			"autoMessage": nowSiteInfo.AutoMessage,
		}

		// 返回JSON响应
		c.JSON(http.StatusOK, response)

	})

	//接受前端数据并进行数据库内容修改
	setup.POST("/api/admin/editInfo", func(c *gin.Context) {

		//将前端数据接受并存储到newSiteInfo
		var newSiteInfo site_info
		if err := c.BindJSON(&newSiteInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 JSON 数据"})
			println("123")
			return
		}

		query := "UPDATE site_info SET name=?, title=?, seo=?, description=?, autoMessage=? WHERE id=?"

		_, err := db.Exec(query, newSiteInfo.Name, newSiteInfo.Title, newSiteInfo.Seo, newSiteInfo.Description, newSiteInfo.AutoMessage, 1)
		if err != nil {
			c.JSON(http.StatusNotFound, "修改失败")
			log.Fatal(err)
		} else {
			c.JSON(http.StatusOK, "修改成功")
			nowSiteInfo = newSiteInfo
		}

	})

	setup.Run(":8081")
}

func basicSetupMysql(db *sql.DB) {
	//用于初始化构建 基础设置数据库
	Message := "Admin-setup: "

	createAdminSql := `
	create table if not exists site_info(
	    id int unsigned auto_increment primary key ,
	    name varchar(255) ,
	    title varchar(255) ,
	    seo varchar(255) ,
	    description varchar(255),
	    autoMessage varchar(255)
	)
`
	_, err := db.Exec(createAdminSql)
	if err != nil {
		return
	}

	var adminCount int
	noSelectAdmin := `
	select COUNT(*) from site_info
`
	err = db.QueryRow(noSelectAdmin).Scan(&adminCount)
	if err != nil {
		println("查询是否存在基础设置失败")
		return
	}
	if adminCount == 0 {
		insertAdminSql := `
	INSERT INTO site_info (name,title,seo,description,autoMessage) values  ("开心发卡网","开心发卡网","发卡网,专业发卡网,高端发卡网","一款功能齐全的发卡系统","欢迎您使用开心发卡网.\n开心发卡系统官网：\n如果您在使用中发现bug或有好的建议清河我们联系\n客服qq:")
`
		_, err = db.Exec(insertAdminSql)
		if err != nil {
			println(Message, "初始化系统基础设置失败")
			return
		} else {
			println(Message, "初始化系统基础设置成功")
		}
	}
}
