package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// User 有多张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

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
	//_ = db.AutoMigrate(&User{}, &CreditCard{})
	var user User
	//db.Create(&user)
	//db.Create(&CreditCard{
	//	Number: "12",
	//	UserID: user.ID,
	//})
	//db.Create(&CreditCard{
	//	Number: "34",
	//	UserID: user.ID,
	//})
	db.Preload("CreditCards").First(&user) //preload CreditCards，而不是CreditCard
	for _, card := range user.CreditCards {
		println(card.Number)
	}

	//在大型的系统中，我个人不建议使用外键约束，外键约束也有很大的优点：数据的完整性
	/*
		外键约束会让给你的数据很完整，即使是业务代码有些人考虑的不严谨
		在大型的系统，高并发的系统中一般不使用外键约束，自己在业务层面保证数据的一致性
	*/
	/*
		在 GORM 中，Preload 方法的参数是结构体的字段名称，而不是表名，这是因为 GORM 是一个 ORM 框架，它的设计目的是为了方便地操作模型对象，而不是直接操作数据库表。
		虽然在 GORM 中，模型通常对应着一个数据库表，但是模型和表之间并不是一一对应的关系。一个模型可能对应着多个表，或者一个表可能对应着多个模型。
		因此，使用模型名而不是表名作为 Preload 方法的参数，可以更加灵活地进行关联查询。
	*/
}
