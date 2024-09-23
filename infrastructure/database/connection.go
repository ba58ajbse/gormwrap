package database

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLHandler struct {
	DB *gorm.DB
}

// type SQLHandlerIF interface {
// 	Create(v interface{}) error
// 	Find(v interface{}) error
// }

func Init() *SQLHandler {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", user, password, host, port, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return &SQLHandler{
		DB: db,
	}
}

func (s *SQLHandler) Create(v interface{}) error {
	return s.DB.Create(v).Error
}

func (s *SQLHandler) Find(v interface{}) error {
	err := s.DB.Find(v).Error
	if err != nil {
		return errors.New("server error")
	}

	return nil
}
func (s *SQLHandler) First(v interface{}, id uint) error {
	res := s.DB.First(v, id)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("record not found")
		}

		return errors.New("server error")
	}

	return nil
}
