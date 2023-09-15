package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// 如果错误返回 err, 如果正确返回 d
	d, err := gorm.Open("mysql", "root:123456@tcp(172.16.10.140)/simplerest?charset=utf8mb4&parseTime=True&loc=Local")
	// 等于 nil 的情况下会返回 panic(err)
	if err != nil {
		panic(err)
	}
	// 正确返回 d 并将值传递给 db
	db = d
}

func GetDB() *gorm.DB {
	return db
}
