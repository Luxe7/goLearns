package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name      string
	CompanyID int //这个字段在实际使用的时候不需要，但是需要声明，用于数据库操作
	Company   Company
}

//type User struct { 自定义外键名字
//	gorm.Model
//	Name         string
//	CompanyRefer int
//	Company      Company `gorm:"foreignKey:CompanyRefer"`
//	// 使用 CompanyRefer 作为外键
//}

type Company struct {
	ID   int
	Name string
}

// 也可重写引用
//
//	type User struct {
//		gorm.Model
//		Name      string
//		CompanyID string
//		Company   Company `gorm:"references:Code"` // 使用 Code 作为引用
//	}
//
//	type Company struct {
//		ID   int
//		Code string
//		Name string
//	}
func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gormlearn?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level Info会打印出所有的SQL语句
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	//删除原本的表后执行
	//_ = db.AutoMigrate(&User{}) //会自动设置外键
	//db.Create(&User{
	//	Name: "Luxe1",
	//	Company: Company{
	//		Name: "helloworld",
	//	},
	//})
	//db.Create(&User{
	//	Name: "Luxe1",
	//	Company: Company{
	//		ID: 1,
	//	},
	//}) //如果此时再填入公司信息，则会创建一个新的公司数据，这不是我们想要的
	var user User
	//db.First(&user)//这样查询只能得到user表中的信息，无法得到company表中的信息
	//db.Preload("Company").First(&user) //预加载
	db.Joins("Company").First(&user) //做了一个join操作，这两种方法都可以连表查询查询
	println(user.Name, user.CompanyID)
}
