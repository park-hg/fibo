package migrate

import (
	"encoding/json"
	commentAdapter "fibo/internal/comment/adapter"
	postAdapter "fibo/internal/post/adapter"
	userAdapter "fibo/internal/user/adapter"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Configuration struct {
	user     string
	password string
	db       string
	host     string
}

func NewMySQLRepository() *gorm.DB {
	file, err := os.Open("../config.json")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("dsn 설정 실패하였습니다.")
		}
	}(file)
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		panic("dsn 설정에 실패하였습니다.")
	}
	dsn := fmt.Sprintf("%s:%s@(%s:3306)/%s", config.user, config.password, config.host, config.db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db 연결에 실패하였습니다.")
	}

	models := []interface{}{
		&commentAdapter.Comment{},
		&postAdapter.Post{},
		&userAdapter.User{},
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		return nil
	}
	return db
}
