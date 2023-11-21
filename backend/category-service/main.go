// 分类服务
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
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

	createClassSql := `
	create table if not exists class(
	 id int unsigned auto_increment primary key ,
	 name varchar(255),
	 info varchar(255),
	 sort int,
	 tag varchar(255)
);
`

	_, err = db.Exec(createClassSql)
	if err != nil {
		return
	}

	type class struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info string `json:"info"`
		Sort int    `json:"sort"`
		Tag  string `json:"tag"`
	}

	//定义一个 2维切片 用来存储分类  一维的大小 pagesize 为 10
	var classes [][]class
	//创建一个新的一维切片，用来存储每页的数据
	var currentPage []class
	//定义 n  为 有多少页数
	n := 0

	var classCount int
	noClass := `select COUNT(*) from class`

	err = db.QueryRow(noClass).Scan(&classCount)
	if err != nil {
		return
	}
	if classCount != 0 {
		row, err := db.Query("select * from class ORDER BY sort ASC")
		if err != nil {
			return
		}

		for row.Next() {
			//定义c作为临时储存
			var c class
			if err := row.Scan(&c.Id, &c.Name, &c.Info, &c.Sort, &c.Tag); err != nil {
				return
			}
			currentPage = append(currentPage, c)
			//每10条为1页
			if len(currentPage) == 10 {
				classes = append(classes, currentPage)
				currentPage = []class{}
			}
		}
		//上述函数 仍会把低于五条的 数据保留在 currentPage中
		if len(currentPage) > 0 {
			classes = append(classes, currentPage)
		}
		n = len(classes)
	}

	e.GET("/api/admin/classPageNum", func(c *gin.Context) {
		c.JSON(200, n)
	})

	//传输列表数据
	e.GET("/api/admin/allClass/page=:page", func(c *gin.Context) {
		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			c.JSON(404, "页数错误")
			return
		}
		if n == 0 {
			c.JSON(204, "空数据")
			return
		}
		if page > n || page <= 0 {
			c.JSON(404, "页数错误")
			return
		}
		c.JSON(200, classes[page-1])
	})

	e.POST("/api/admin/addClass", func(c *gin.Context) {
		var newClass class
		err := c.BindJSON(&newClass)
		if err != nil {
			c.JSON(404, "添加错误")
			return

		}
		query := "INSERT INTO class (name, info, sort, tag) values (?,?,?,?)"
		_, err = db.Exec(query, newClass.Name, newClass.Info, newClass.Sort, newClass.Tag)
		if err != nil {
			c.JSON(404, "插入错误")
			log.Fatal(err)
		}
		c.JSON(200, "成功")

		row, err := db.Query("select * from class ORDER BY sort ASC")
		if err != nil {
			return
		}

		//初始化classes,currentPage
		classes = make([][]class, 0)
		currentPage = make([]class, 0)

		for row.Next() {
			//定义c作为临时储存
			var c class
			if err := row.Scan(&c.Id, &c.Name, &c.Info, &c.Sort, &c.Tag); err != nil {
				return
			}
			currentPage = append(currentPage, c)
			//每10条为1页
			if len(currentPage) == 10 {
				classes = append(classes, currentPage)
				currentPage = []class{}
			}
		}
		//上述函数 仍会把低于五条的 数据保留在 currentPage中
		if len(currentPage) > 0 {
			classes = append(classes, currentPage)
		}
		//定义 n  为 有多少页数
		n = len(classes)

	})

	e.POST("api/admin/editClass", func(c *gin.Context) {
		var newClass class
		err := c.BindJSON(&newClass)
		if err != nil {
			c.JSON(404, "解析json错误")
			return
		}

		//用于更新数据库的语句
		editClassSql := `
	UPDATE class SET name = ?, info = ?, sort = ?, tag = ? WHERE id = ?;
`
		_, err = db.Exec(editClassSql, newClass.Name, newClass.Info, newClass.Sort, newClass.Tag, newClass.Id)
		if err != nil {
			c.JSON(404, "修改错误")
			return
		}

		c.JSON(200, "成功")

		row, err := db.Query("select * from class ORDER BY sort ASC")
		if err != nil {
			return
		}

		//初始化classes,currentPage
		classes = make([][]class, 0)
		currentPage = make([]class, 0)

		for row.Next() {
			//定义c作为临时储存
			var c class
			if err := row.Scan(&c.Id, &c.Name, &c.Info, &c.Sort, &c.Tag); err != nil {
				return
			}
			currentPage = append(currentPage, c)
			//每10条为1页
			if len(currentPage) == 10 {
				classes = append(classes, currentPage)
				currentPage = []class{}
			}
		}
		//上述函数 仍会把低于五条的 数据保留在 currentPage中
		if len(currentPage) > 0 {
			classes = append(classes, currentPage)
		}
		//定义 n  为 有多少页数
		n = len(classes)

	})

	e.POST("/api/admin/deleteClass", func(c *gin.Context) {
		var id int
		err := c.BindJSON(&id)
		if err != nil {
			return
		}
		deleteClassSql := `
delete from class where id = ?
`
		_, err = db.Exec(deleteClassSql, id)
		if err != nil {
			return
		}
		c.JSON(200, "删除成功")

		row, err := db.Query("select * from class ORDER BY sort ASC")
		if err != nil {
			return
		}

		//初始化classes,currentPage
		classes = make([][]class, 0)
		currentPage = make([]class, 0)

		for row.Next() {
			//定义c作为临时储存
			var c class
			if err := row.Scan(&c.Id, &c.Name, &c.Info, &c.Sort, &c.Tag); err != nil {
				return
			}
			currentPage = append(currentPage, c)
			//每10条为1页
			if len(currentPage) == 10 {
				classes = append(classes, currentPage)
				currentPage = []class{}
			}
		}
		//上述函数 仍会把低于五条的 数据保留在 currentPage中
		if len(currentPage) > 0 {
			classes = append(classes, currentPage)
		}
		//定义 n  为 有多少页数
		n = len(classes)

	})

	e.POST("/api/admin/deleteSelectionClass", func(c *gin.Context) {
		var ids []int
		err := c.BindJSON(&ids)
		if err != nil {
			return
		}

		// 构建 SQL 查询，使用字符串拼接构建 IN 子句
		deleteSql := "DELETE FROM class WHERE id IN ("
		for i, id := range ids {
			if i > 0 {
				deleteSql += ","
			}
			deleteSql += fmt.Sprintf("%d", id)
		}
		deleteSql += ")"

		_, err = db.Exec(deleteSql)
		if err != nil {

			fmt.Println("Failed to execute SQL query:", err)

			c.JSON(500, "no")
			return
		}

		c.JSON(200, "yes")

		row, err := db.Query("select * from class ORDER BY sort ASC")
		if err != nil {
			return
		}

		//初始化classes,currentPage
		classes = make([][]class, 0)
		currentPage = make([]class, 0)

		for row.Next() {
			//定义c作为临时储存
			var c class
			if err := row.Scan(&c.Id, &c.Name, &c.Info, &c.Sort, &c.Tag); err != nil {
				return
			}
			currentPage = append(currentPage, c)
			//每10条为1页
			if len(currentPage) == 10 {
				classes = append(classes, currentPage)
				currentPage = []class{}
			}
		}
		//上述函数 仍会把低于五条的 数据保留在 currentPage中
		if len(currentPage) > 0 {
			classes = append(classes, currentPage)
		}
		//定义 n  为 有多少页数
		n = len(classes)

	})

	e.Run(":8084")
}
