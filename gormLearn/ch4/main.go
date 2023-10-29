package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gormlearn?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level Info会打印出所有的SQL语句
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&User{})

	var users = []User{{Name: "luxe1"}, {Name: "luxe2"}, {Name: "luxe3"}}

	result := db.Create(&users) // 传递切片以插入多行数据
	//GORM will generate a single SQL statement to insert all the data and backfill primary key values, hook methods will be invoked too.

	//You can specify batch size when creating with CreateInBatches
	//一次批量提交有性能优势，但是SQL语句有长度限制，所以有如下的分批提交函数
	db.CreateInBatches(&users, 2)

	println(result.Error)        // 返回 error
	println(result.RowsAffected) // 返回插入记录的条数

	/*
		发现了一个很奇怪的问题，同样的users切片，如果不做重新赋值，插入一次后就无法再插入，应该是主键ID的原因
		删除主键字段后，可以重复插入，但是我打印出user.ID的值，发现他们都是一样的，非常不理解为什么会有这样的情况
	*/

	//GORM supports create from map[string]interface{} and []map[string]interface{}{}, e.g:

	//db.Model(&User{}).Create(map[string]interface{}{
	//	"Name": "luxe7", "Age": 18,
	//})
	//
	//// batch insert from `[]map[string]interface{}{}`
	//db.Model(&User{}).Create([]map[string]interface{}{
	//	{"Name": "luxe1", "Age": 18},
	//	{"Name": "luxe2", "Age": 20},
	//})
	//map创建的可维护性不如结构体实例创建

	//关联创建
	//type CreditCard struct {
	//	gorm.Model
	//	Number   string
	//	UserID   uint
	//}
	//
	//type User struct {
	//	gorm.Model
	//	Name       string
	//	CreditCard CreditCard
	//}
	//
	//db.Create(&User{
	//	Name: "jinzhu",
	//	CreditCard: CreditCard{Number: "411111111111"}
	//})
	//// INSERT INTO `users` ...
	//// INSERT INTO `credit_cards` ...同时创建了两个记录
}
