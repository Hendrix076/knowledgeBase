package common

import (
	"github.com/mattn/go-colorable"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Gorm *gorm.DB
func InitDB(){
	Gorm=gormDB()
}
func gormDB() *gorm.DB {
	newLogger := logger.New(
		log.New(colorable.NewColorableStdout(), "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)
	dsn:="devuser:123~!@@tcp(39.105.28.235:3320)/tech?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB,err:=db.DB()
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	return db
}