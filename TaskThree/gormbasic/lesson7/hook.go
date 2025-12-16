package lesson7

import (
	"fmt"

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

func (p *Parent) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Parent BeforeCreate")
	return nil
}

func (c *Child) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Child BeforeCreate")
	return nil
}

func RunHook(db *gorm.DB) {
	db.AutoMigrate(&Parent{}, &Child{})

	db.Create(&Child{Parent: Parent{Name: "aa"}, Age: 10})
}
