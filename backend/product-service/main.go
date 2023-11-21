package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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
		err.Error()
		return
	}

	type commodityInfo struct {
		//id
		ID int `json:"id"`
		//分类id
		ClassID int `json:"classId"`
		//商品名称
		Name string `json:"name"`
		//商品描述
		Info string `json:"info"`
		//成本价
		CostPrice float64 `json:"costPrice"`
		//商品单价
		Price float64 `json:"price"`
		//会员价
		PriceVip float64 `json:"priceVip"`
		//发货方式  自动发货为 true  手动发货为 false
		DeliveryMethod string `json:"deliveryMethod"`
		//发货留言 当 自动发货 后 的 发货模板
		DeliveryMessage string `json:"deliveryMessage"`
		//发货排序 当 自动发货 后 如何发送卡密
		//当 为 0 则 优先 发 旧
		//当 为 1 则 随机 发送
		//当 为 2 则 优先 发 新
		DeliverySort string `json:"deliverySort"`
		//权重
		Sort int `json:"sort"`
		//状态
		Tag string `json:"tag"`
		//发货联系方式 即 邮箱(0) 短信(1)
		Contact string `json:"contact"`
	}

	type requestData struct {
		Commodity commodityInfo         `form:"commodity"`
		File      *multipart.FileHeader `form:"files"`
	}

	//创建商品 基础信息表
	createCommoditySql := `
	create table if not exists commodity(
	 id int unsigned auto_increment primary key ,
	 classId varchar(255),
	 name varchar(255),
	 info varchar(255),
	 costPrice float,
	 price float,
	 priceVip float,
	 DeliveryMethod varchar(255),
	 DeliveryMessage varchar(255),
	 DeliverySort int ,
	 sort int,
	 tag varchar(255),
	 Contact int
);
`
	_, err = db.Exec(createCommoditySql)
	if err != nil {
		err.Error()
		return
	}

	//定义 n  为 有多少页数
	n := 0

	////定义一个 2维切片 用来存储分类  一维的大小 pagesize 为 10
	//var commodities [][]commodityInfo
	////创建一个新的一维切片，用来存储每页的数据
	//var currentPage []commodityInfo
	//
	//var commodityCount int
	//noCommodity := `select count(*) from commodityInfo`
	//err = db.QueryRow(noCommodity).Scan(&commodityCount)
	//if commodityCount != 0 {
	//	row, err := db.Query("select * from commodityInfo ORDER BY sort ASC")
	//	if err != nil {
	//		return
	//	}
	//
	//	for row.Next() {
	//		//定义c作为临时储存
	//		var c commodityInfo
	//		if err := row.Scan(&c); err != nil {
	//			return
	//		}
	//		currentPage = append(currentPage, c)
	//		//每10条为1页
	//		if len(currentPage) == 10 {
	//			commodities = append(commodities, currentPage)
	//			currentPage = []commodityInfo{}
	//		}
	//	}
	//	//上述函数 仍会把低于五条的 数据保留在 currentPage中
	//	if len(currentPage) > 0 {
	//		commodities = append(commodities, currentPage)
	//	}
	//	n = len(commodities)
	//}

	e.GET("/api/admin/commodityPageNum", func(c *gin.Context) {
		c.JSON(200, n)
	})

	//用于添加商品分类
	e.POST("/api/admin/addCommodity", func(c *gin.Context) {
		//// 解析请求中的文件和表单数据
		//err := c.Request.ParseMultipartForm(10 << 20) // 10 MB的最大内存大小
		//if err != nil {
		//	c.JSON(400, "解析文件和表单数据错误")
		//	return
		//}

		//定义并解析数据
		var requestData requestData
		if err := c.ShouldBind(&requestData); err != nil {
			c.JSON(400, "解析表单数据错误")
			println(err.Error())
			return
		}

		//c.BindJSON(&requestData.Commodity)

		// 获取表单字段值 从解析的 json包中

		formData := requestData.Commodity
		addCommoditySql := `insert into commodity (classId, name, info, costPrice, price, priceVip, DeliveryMethod, DeliveryMessage, DeliverySort, sort, tag, Contact) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`
		println(formData.Name)
		println(requestData.Commodity.Name)
		result, err := db.Exec(addCommoditySql, formData.ClassID, formData.Name, formData.Info, formData.CostPrice, formData.Price,
			formData.PriceVip, formData.DeliveryMethod, formData.DeliveryMessage, formData.DeliverySort, formData.Sort, formData.Tag, formData.Contact)

		if err != nil {
			return
		}

		// 获取插入的记录的ID
		insertedID, err := result.LastInsertId()

		// 获取上传的文件
		file, _, err := c.Request.FormFile("files")
		if err != nil {
			c.JSON(http.StatusBadRequest, "获取上传文件错误")
			return
		}
		defer file.Close()

		folderPath := "./data/commodityIcons"

		// 创建文件夹（如果不存在）
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 为文件生成新文件名，使用插入的记录ID
		newFileName := fmt.Sprintf("%d_icon.jpg", insertedID)

		// 创建新文件
		newFilePath := filepath.Join(folderPath, newFileName)
		newFile, err := os.Create(newFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer newFile.Close()

		// 将上传的文件内容复制到新文件中
		_, err = io.Copy(newFile, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 补充 更多信息
		addMoreInfo := `insert into commodity_more_info (id, inventory, todaySell, yesterdaySell, allSell, extendId) VALUES (?,0,0,0,0,?)`

		//将id转为 8位 字符
		IDString := fmt.Sprintf("%08d", insertedID)
		//将内容 生成 插入到 更多信息表
		_, err = db.Exec(addMoreInfo, insertedID, IDString)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "成功添加"})
	})

	//用于给商品管理页面提供 分类信息

	type class struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info string `json:"info"`
		Sort int    `json:"sort"`
		Tag  string `json:"tag"`
	}

	var nowClass []class

	e.GET("/api/admin/commodityGetClass", func(c *gin.Context) {
		nowClass = make([]class, 0)
		row, err := db.Query("select * from class ORDER BY sort ASC")
		if err != nil {
			println(err.Error())
			return
		}
		for row.Next() {
			var class class
			err := row.Scan(&class.Id, &class.Name, &class.Info, &class.Sort, &class.Tag)
			if err != nil {
				println(err.Error())
				c.JSON(500, "查询出错")
				return
			}
			nowClass = append(nowClass, class)
		}

		c.JSON(200, nowClass)
	})

	//用于储存 更多信息 也就是 展示信息

	createMoreCommodityInfo := `
	create table if not exists commodity_more_info(
# 	 商品id
	 id int,
# 	 商品库存
	 inventory int,
# 	 今日销售
	 todaySell float,
# 	 昨日销售
	 yesterdaySell float,
# 	 总销售
	 allSell float,
# 	 推广id
	 extendId varchar(64)                           
)
`
	_, err = db.Exec(createMoreCommodityInfo)
	if err != nil {
		return
	}

	type allCommodityInfo struct {
		commodityInfo
		Inventory     int     `json:"inventory"`
		TodaySell     float64 `json:"todaySell"`
		YesterdaySell float64 `json:"yesterdaySell"`
		AllSell       float64 `json:"allSell"`
		ExtendID      string  `json:"extendId"`
	}

	e.Run(":8085")
}
