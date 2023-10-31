package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Language struct {
	//gorm.Model
	Name    string
	AddTime time.Time
}

// 在gorm中可以通过给某一个struct.添加TableName.方法来自定义表名
// func (Language)TableName（）string{
// return "my_language"
// }

/*
1,我们自己定义表名是什么
2,统一的给所有的表名加上一个前缀 NamingStrategy: schema.NamingStrategy{TablePrefix: "Hello_"}
这两者是冲突的，同时写入时优先运行TableName方法
3.如果要添加模型的缺省值，可以使用钩子函数
*/

func (u *Language) BeforeCreate(tx *gorm.DB) (err error) {
	u.AddTime = time.Now()
	return
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
		NamingStrategy: schema.NamingStrategy{TablePrefix: "Hello_"},
		Logger:         newLogger,
	})

	_ = db.AutoMigrate(&Language{})
	db.Create(&Language{
		Name: "go",
	}) //因为设置了hook所以会自动添加时间

}
