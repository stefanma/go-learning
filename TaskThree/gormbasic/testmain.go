package main

import (
	"gorm.io/gorm"
)

type Parent struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Child struct {
	Parent
	Age int
}

func InitDB(dst ...interface{}) *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:st123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(dst...)

	return db
}

func main() {
	// db, err := gorm.Open(mysql.Open("root:st123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	// if err != nil {
	// 	panic(err)
	// }
	// lesson1.Run(db)
	// lesson2.Run(db)
	// lesson3.Run(db)
	// lesson4.Run(db)
	// lesson5.Run(db)
	// lesson6.Run(db)
	// lesson7.Run(db)

	InitDB(&Parent{}, &Child{})
}
