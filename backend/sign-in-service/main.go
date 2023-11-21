// 登录授权服务
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
	login := gin.Default()

	// 配置CORS中间件，允许所有来源跨域访问
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	login.Use(cors.New(config))

	//创建数据库对象
	db, err := mysqlConnect()
	if err != nil {
		return
	}

	createAdminSql := `		CREATE TABLE IF NOT EXISTS admin(
			id int unsigned auto_increment primary key ,
			name varchar(255),
			password varchar(255)
			)`

	_, err = db.Exec(createAdminSql)
	if err != nil {
		return
	}

	var adminCount int
	noSelectAdmin := `
	select COUNT(*) from admin
`
	err = db.QueryRow(noSelectAdmin).Scan(&adminCount)
	if err != nil {
		println("查询是否存在管理员账户失败")
		return
	}
	if adminCount == 0 {
		insertAdminSql := `
	INSERT INTO admin (name,password) values  ("admin","123456")
`
		_, err = db.Exec(insertAdminSql)
		if err != nil {
			println("初始化admin密码失败")
			return
		} else {
			println("初始化admin密码成功")
		}
	}

	type User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	login.POST("/admin/login", func(c *gin.Context) {
		var loginInfo User
		if err := c.BindJSON(&loginInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		selectPassword := `
		select name,password from admin
`
		row := db.QueryRow(selectPassword)
		var dbname, dbPassword string
		if err := row.Scan(&dbname, &dbPassword); err != nil {
			println("检索管理员密码出现问题")
			return
		}

		if loginInfo.Username == dbname && loginInfo.Password == dbPassword {
			c.JSON(http.StatusOK, "登录成功")
		} else {
			c.JSON(http.StatusUnauthorized, "登录失败")
		}

	})

	runErr := login.Run(":8080")
	if runErr != nil {
		return
	}
}
