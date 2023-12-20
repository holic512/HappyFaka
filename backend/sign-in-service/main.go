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

	//管理后台的登录
	sign.POST("/admin/login", func(c *gin.Context) {

		var loginInfo User

		//获取前端数据
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

	//调用token演示
	sign.POST("/admin/api", func(c *gin.Context) {
		token, err := c.Request.Cookie("token")
		if err != nil {
			return
		}
		println(token.Value)
	})

	//用户面 登录
	sign.POST("/user/login", func(c *gin.Context) {

	})

	//储存传入 json
	type registration struct {
		Username              string `json:"username"`
		Password              string `json:"password"`
		Phone                 string `json:"phone"`
		Email                 string `json:"mail"`
		EmailVerificationCode string `json:"email_verification_code"`
		InvitationCode        string `json:"invitation_code"`
	}
	//用户面 注册
	sign.POST("/user/register", func(c *gin.Context) {
		var info registration

		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(404, "注册失败")
			return
		}

		//用于邮箱验证
		if false {
			c.JSON(404, "邮箱验证码错误")
			return
		}

		insertUser := `INSERT INTO user (username,password,mail,phone,invite_code) values (?,?,?,?,?)`

		_, err = db.Exec(insertUser, info.Username, info.Password, info.Email, info.Phone, info.InvitationCode)
		if err != nil {
			c.JSON(404, "注册失败")
			return
		}
		//到此代表注册成功

		//分发token
		token := token_module.GetToken(info.Username)
		expireTime := time.Now().Add(24 * time.Hour) // 计算一天后的时间
		expirationSeconds := int(expireTime.Unix())  // 转换为 Unix 时间戳的秒数
		c.SetCookie("token", token, expirationSeconds, "/admin", "", false, true)

		c.JSON(200, "登录成功")
	})

	//运行
	runErr := sign.Run(":8080")
	if runErr != nil {
		return
	}
}
