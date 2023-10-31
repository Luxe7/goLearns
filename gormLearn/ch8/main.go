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
	DeletedAt    gorm.DeletedAt
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
	_ = db.AutoMigrate(&User{})
	var user User
	//db.Where(&User{Name: "luxe2"}).Delete(&user)
	// 此时模型中不包括gorm.DeletedAt字段，所以删除是硬删除，当加入gorm.DeletedAt字段后，就会有软删除的能力
	//db.Where(&User{Name: "luxe3"}).Delete(&user)
	db.Unscoped().Where(&User{Name: "luxe3"}).Find(&user) //查找被软删除的数据
	db.Unscoped().Delete(&user)                           //永久删除被软删除的数据

}
