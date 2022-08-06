package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "Saud:12345@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Print("dddd")
		panic(err)
	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}
