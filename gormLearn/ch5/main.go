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

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	//var user User
	//// 获取第一条记录（主键升序）
	//println(db.First(&user))
	//// SELECT * FROM users ORDER BY id LIMIT 1;
	//
	//// 获取一条记录，没有指定排序字段
	//println(db.Take(&user))
	//// SELECT * FROM users LIMIT 1;
	//
	//// 获取最后一条记录（主键降序）
	//println(db.Last(&user))
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	//
	//result := db.First(&user)
	//
	////没有找到记录时，它会返回 ErrRecordNotFound 错误
	//// 检查 ErrRecordNotFound 错误
	//errors.Is(result.Error, gorm.ErrRecordNotFound)

	var users []User
	result := db.Find(&users)
	println("总行数：", result.RowsAffected)
	for _, user := range users {
		println(user.Name)
	}

}
