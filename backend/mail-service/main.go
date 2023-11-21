// 邮件服务
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
	"net/smtp"
	"os"

	myemail "github.com/jordan-wright/email"
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
	e := gin.Default()

	// 配置CORS中间件，允许所有来源跨域访问
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	e.Use(cors.New(config))

	//创建数据库对象
	db, err := mysqlConnect()
	if err != nil {
		return
	}

	createEmailSql := `
	create table if not exists email_info(
	 id int unsigned auto_increment primary key ,
	 smtp_server varchar(255),
	 username varchar(255),
	 authorization_code varchar(255)
)

`
	//创建数据库
	_, err = db.Exec(createEmailSql)
	if err != nil {
		return
	}

	//email数据库初始化项目
	var emailCount int
	noSelectEmail := `
	select COUNT(*) from email_info
`
	err = db.QueryRow(noSelectEmail).Scan(&emailCount)
	if err != nil {
		println("查询是否存在邮箱设置失败")
		return
	}
	if emailCount == 0 {
		insertEmailSql := `
	INSERT INTO email_info (smtp_server,username,authorization_code) values  ("QQ","#","#")
`
		_, err = db.Exec(insertEmailSql)
		if err != nil {
			println("初始化数据库设置失败")
			return
		} else {
			println("初始化数据库设置成功")
		}
	}

	//设置email结构体
	type email struct {
		SmtpServer        string `json:"server"`
		Username          string `json:"name"`
		AuthorizationCode string `json:"authcode"`
	}

	var nowEmail email
	selectInfo := `
	select smtp_server,username,authorization_code from email_info
`
	row := db.QueryRow(selectInfo)
	if err := row.Scan(&nowEmail.SmtpServer, &nowEmail.Username, &nowEmail.AuthorizationCode); err != nil {
		println("检索邮箱设置出现问题")
		return
	}

	//获取email的info
	e.GET("/api/email/getInfo", func(c *gin.Context) {
		// 构建JSON响应
		response := gin.H{
			"server":   nowEmail.SmtpServer,
			"name":     nowEmail.Username,
			"authcode": nowEmail.AuthorizationCode,
		}

		// 返回JSON响应
		c.JSON(http.StatusOK, response)
	})

	//接受前端数据并进行数据库内容修改
	e.POST("/api/email/editInfo", func(c *gin.Context) {

		//将前端数据接受并存储到newSiteInfo
		var newEmail email
		if err := c.BindJSON(&newEmail); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 JSON 数据"})
			return
		}

		query := "UPDATE email_info SET smtp_server=?, username=?, authorization_code=? WHERE id=?"

		_, err := db.Exec(query, newEmail.SmtpServer, newEmail.Username, newEmail.AuthorizationCode, 1)
		if err != nil {
			c.JSON(http.StatusNotFound, "修改失败")
			log.Fatal(err)
		} else {
			c.JSON(http.StatusOK, "修改成功")
			nowEmail = newEmail
		}

	})

	//创建邮箱发送账号结构体
	type senEmail struct {
		Email string `json:"sendEmail"`
	}

	//测试邮件发送
	e.POST("/api/email/sendEmailIf", func(c *gin.Context) {

		var a senEmail
		//将前端数据接受并存储到newSiteInfo
		if err := c.BindJSON(&a); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 JSON 数据"})
			return
		}

		if nowEmail.SmtpServer == "QQ" {
			auth := smtp.PlainAuth("", nowEmail.Username, nowEmail.AuthorizationCode, "smtp.qq.com")
			e := &myemail.Email{
				From:    nowEmail.Username,
				To:      []string{a.Email},
				Subject: "测试邮件",
				Text:    []byte("如果你能看到这个邮件，说明邮箱配置成功"),
			}

			err := e.Send("smtp.qq.com:587", auth)
			if err != nil {
				c.JSON(http.StatusNotFound, "发送成功")
				log.Fatal(err)
			} else {
				c.JSON(http.StatusOK, "发送失败")
			}

		}

	})

	err = e.Run(":8083")
	if err != nil {
		return
	}
}
