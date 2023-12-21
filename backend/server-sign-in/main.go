// 登录授权服务
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"mysql-module"
	"net/http"
	"time"
	"token-module"
)

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

	//数据库初始化
	initAdminSql(db)
	initUserSql(db)

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
		//登录失败 代表用户名或者密码错误 登录错误 代表 程序出现了问题

		//获取前端数据
		var user User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(200, "登录错误")
			return
		}

		//查询账户是否存在
		var userNameCount int
		userNameCountSql := `select count(*) from user where username = ?`
		err = db.QueryRow(userNameCountSql, user.Username).Scan(&userNameCount)
		if err != nil {
			c.JSON(200, "登录错误")
			return
		}

		if userNameCount != 1 {
			//账号过多或者不存在账号
			c.JSON(200, "登录失败")
			return
		}

		//将密码转成 哈希值
		hashpassword, err := hashPassword(user.Password)
		if err != nil {
			c.JSON(200, "登录错误")
			return
		}

		//与数据库 比较
		var password string
		passwordSql := `select password from user where username = user.username`
		err = db.QueryRow(passwordSql).Scan(&password)
		if hashpassword != password {
			//密码错误
			c.JSON(200, "登录失败")
			return
		}

		//到这里登录成功

		//更改上次登录时间
		updateLoginTimeSql := `UPDATE user set last_login_time = now() where username = ?`
		_, err = db.Exec(updateLoginTimeSql, user.Username)
		if err != nil {
			c.JSON(200, "登录错误")
			return
		}

		//分发token
		token := token_module.GetToken(user.Username)
		expireTime := time.Now().Add(24 * time.Hour) // 计算一天后的时间
		expirationSeconds := int(expireTime.Unix())  // 转换为 Unix 时间戳的秒数
		c.SetCookie("token", token, expirationSeconds, "/admin", "", false, true)

		c.JSON(200, "登录成功")
	})

	//注册数据 储存传入 json
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
			c.JSON(200, "注册失败")
			return
		}

		//用于邮箱验证
		if false {
			c.JSON(200, "邮箱验证码错误")
			return
		}

		//验证邮箱是否被注册过
		var mailCount int
		mailCountSql := `select COUNT(*) from user where mail = ?`
		err = db.QueryRow(mailCountSql, info.Email).Scan(&mailCount)
		if err != nil {
			return
		}
		if mailCount != 0 {
			c.JSON(200, "邮箱已被注册")
			return
		}

		//验证用户名是否被注册过
		var userCount int
		userCountSql := `select count(*) from user where username = ?`
		err = db.QueryRow(userCountSql, info.Username).Scan(&userCount)
		if userCount != 0 {
			c.JSON(200, "用户名已被注册")
			return
		}

		//创建哈希值用来储存密码
		hashaedPassword, err := hashPassword(info.Password)
		if err != nil {
			c.JSON(200, "注册失败")
			return
		}

		insertUser := `INSERT INTO user (username,password,status,balance,if_merchant,phone,mail,invite_code,inviter_id,registration_time,last_login_time) values (?,?,?,?,?,?,?,?,?,?,?)`
		//因为邀请码有唯一性 所以 用户名就是邀请码
		_, err = db.Exec(insertUser, info.Username, hashaedPassword, true, 0, false, info.Phone, info.Email, info.Username, info.InvitationCode, time.Now(), time.Now())
		if err != nil {
			c.JSON(200, "注册失败")
			return
		}
		//到此代表注册成功

		//分发token
		token := token_module.GetToken(info.Username)
		expireTime := time.Now().Add(24 * time.Hour) // 计算一天后的时间
		expirationSeconds := int(expireTime.Unix())  // 转换为 Unix 时间戳的秒数
		c.SetCookie("token", token, expirationSeconds, "/admin", "", false, true)

		c.JSON(200, "注册成功")
	})

	//运行
	runErr := sign.Run(":8080")
	if runErr != nil {
		return
	}
}

// 哈希值转换
func hashPassword(password string) (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(password))
	if err != nil {
		return "", err
	}

	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return hashedPassword, nil
}
