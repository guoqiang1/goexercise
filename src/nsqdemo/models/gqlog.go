package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Result struct {
	Id    int64
	Title string
}

type GqLog struct {
	FId      int32 `gorm:"primary_key"`
	FContent string
	FAddtime string
}

func (u *GqLog) TableName() string {
	return "t_log"
}

func (u *GqLog) getConnect() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:8889)/rclx?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Println(err)
		return nil
	}

	return db
}

func (u *GqLog) Add(gqlog GqLog) *gorm.DB {
	db := u.getConnect()
	if db == nil {
		fmt.Println("get connect failed")
	}

	//re := db.Create(&log)
	db.Create(&gqlog)
	fmt.Println("add", gqlog)
	//db.NewRecord(log)
	return nil
}
