package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// User 拥有并属于多种 language，`user_languages` 是连接表
// 当使用 GORM 的 AutoMigrate 为 User 创建表时，GORM 会自动创建连接表
type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
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
	languages := []Language{}
	languages = append(languages, Language{Name: "go"})
	languages = append(languages, Language{Name: "java"})
	languages = append(languages, Language{Name: "python"})
	user := User{
		Languages: languages,
	}
	db.Create(&user)
	//var user User
	//db.Preload("Languages").First(&user)
	//for _, language := range user.Languages {
	//	println(language.Name)
	//}

	// 开始关联模式
	//关联模式下，先查询到user，然后再关联得到languages
	//var user User
	//db.First(&user)
	//var languages []Language
	//_ = db.Model(&user).Association("Languages").Find(&languages)
	//for _, language := range languages {
	//	println(language.Name)
	//}

}
