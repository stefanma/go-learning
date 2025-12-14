package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 定义模型
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
	Age  int
}

func main() {
	// 1. 连接数据库（会自动创建 test.db 文件）
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 2. 自动建表
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	// 3. 插入数据
	user := User{Name: "Tom", Age: 18}
	db.Create(&user)

	// 4. 查询数据
	var result User
	db.First(&result, "name = ?", "Tom")

	fmt.Printf("查询结果: %+v\n", result)
}
