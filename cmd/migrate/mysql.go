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
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
	Host     string `json:"host"`
}

func NewMySQLRepository() *gorm.DB {
	file, err := os.Open("./config.json")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(file, err)
			panic("no such file")
		}
	}(file)
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
		panic("dsn 설정에 실패하였습니다.")
	}
	dsn := fmt.Sprintf("%s:%s@(%s:3306)/%s", config.User, config.Password, config.Host, config.DB)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(dsn, err)
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
