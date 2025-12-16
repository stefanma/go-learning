package lesson7

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type strArr []string

type UserD struct {
	ID     int
	Skills strArr
}

func (arr strArr) Value() (driver.Value, error) {
	return strings.Join(arr, ","), nil
}

func (arr *strArr) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	*arr = strings.Split(string(bytes), ",")
	return nil
}

func RunDefinition(db *gorm.DB) {
	db.AutoMigrate(&UserD{})
	// db.Create(&UserD{Skills: strArr{"s1", "s2", "s3"}})

	var user UserD
	db.First(&user, "id = 1")
	fmt.Println(user)

	// sql.NullString{}
	// sql.NullTime{}
}
