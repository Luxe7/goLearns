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
	//db.Create(&User{Name: "luxe7"})
	//db.Model(&User{ID: 1}).Update("Name", "") //update会更新零值而updates语句不会更新零值

	//empty := ""
	//db.Model(&User{ID: 1}).Updates(User{Email: &empty})

	//解决仅更新非零值字段的方法有两种
	/*
		1.将string设置为*string
		2.使用sql的NULLxxx类型来解决（ch1）
	*/

	user := User{Name: "Luxe7"} //你无法向 ‘create’ 传递结构体，所以你应该传入数据的指针
	result := db.Create(&user)

	println(user.ID)             // 返回插入数据的主键
	println(result.Error)        // 返回 error
	println(result.RowsAffected) // 返回插入记录的条数
}
