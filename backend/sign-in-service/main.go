// 登录授权服务
package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

import "mysql-module"

import "token-module"

func main() {
	sign := gin.Default()

	// 配置CORS中间件，允许所有来源跨域访问
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true // 允许发送身份验证信息（例如 Cookie）
	sign.Use(cors.New(config))

	//创建数据库对象
	db, err := mysql_module.GetMysqlConnect()
	if err != nil {
		return
	}

	type User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	sign.POST("/admin/login", func(c *gin.Context) {

		var loginInfo User
		if err := c.BindJSON(&loginInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		selectPassword := `
		select name,password from admininfo
`
		row := db.QueryRow(selectPassword)
		var dbname, dbPassword string
		if err := row.Scan(&dbname, &dbPassword); err != nil {
			println("检索管理员密码出现问题")
			return
		}

		//账号密码匹配
		if loginInfo.Username == dbname && loginInfo.Password == dbPassword {
			token := token_module.GetToken(loginInfo.Username)
			expireTime := time.Now().Add(24 * time.Hour) // 计算一天后的时间
			expirationSeconds := int(expireTime.Unix())  // 转换为 Unix 时间戳的秒数
			c.SetCookie("token", token, expirationSeconds, "/admin", "", false, true)

			c.JSON(200, "登录成功")

		} else {
			c.JSON(http.StatusUnauthorized, "登录失败")
		}

	})

	sign.POST("/admin/api", func(c *gin.Context) {
		token, err := c.Request.Cookie("token")
		if err != nil {
			return
		}
		println(token.Value)
	})

	runErr := sign.Run(":8080")
	if runErr != nil {
		return
	}
}
