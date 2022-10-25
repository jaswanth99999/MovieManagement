package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	//Mysql server should be running on localhost...
	d, err := gorm.Open("mysql","root:<Password>@tcp(127.0.0.1:3306)/<DBNAME>?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}