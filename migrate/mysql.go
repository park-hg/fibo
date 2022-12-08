package migrate

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:1234@(127.0.0.1:3306)/products"), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}

	models := []interface{

	}
	err = db.AutoMigrate(models...)
	if err != nil {
		return nil
	}
	return db
}
