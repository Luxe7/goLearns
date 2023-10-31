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
	//var user User
	var users []User
	//db.Where("name=?", "luxe2").First(&user)
	//db.Where(&User{Name: "luxe3"}).First(&user)
	/*
		查询方式条件有三种 1.string 2.struct 3.map
		第一种方法需要知道字段在数据库中的名称，第二种办法相当于只需要知道模型中的字段名即可
	*/
	db.Where(&User{Name: "luxe2"}).Find(&users)
	/*
		When querying with struct, GORM will only query with non-zero fields,
		that means if your field’s value is 0, '', false or other zero values, it won’t be used to build query conditions,
		无法查询零值，这跟之前的update语句是一样的，主要是因为go语言会把“缺省”的字段设置为零值，gorm无法判断这个零值来自缺省还是来自使用者
		如果要查询零值，需要写成map[string]interface{}类型
	*/
	for _, user := range users {
		println(user.ID)
	}

}
