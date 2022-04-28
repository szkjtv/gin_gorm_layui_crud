package mysql

import (
	"log"

	"example.com/m/v2/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Dbinit() (db *gorm.DB) {
	db, err = gorm.Open("mysql", "root:aa1451418@tcp(127.0.0.1:3306)/world?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&model.PddData{})
	db.AutoMigrate(&model.IotCard{})
	db.AutoMigrate(&model.RgisterUser{})
	db.AutoMigrate(&model.Commtidy{})
	db.AutoMigrate(&model.ImgTest{}) //测试用的
	// db.AutoMigrate(&model.Article{})
	return db

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
