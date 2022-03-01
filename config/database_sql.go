package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"m3gaplazma/gofiber-restapi/exception"
)

var DB *gorm.DB

func ConnectSQL() {
	var err error
	//dsn := "root:root@tcp(127.0.0.1:" +  + ")/go-restapi?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/go-restapi?charset=utf8mb4&parseTime=True&loc=Local",
		Env("DB_USER"),
		Env("DB_PASSWORD"),
		Env("DB_HOST"),
		Env("DB_PORT"),
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	exception.PanicIfError(err)
	fmt.Println("connected to mysql database")
}
