package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:1234@(127.0.0.1:3306)/fibo"), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}
	models := []interface{}{}

	err = db.AutoMigrate(models...)
	if err != nil {
		log.Fatal(err)
	}
}
