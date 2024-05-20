package external

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func NewDB() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_ACCOUNT"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_ADDRESS"), os.Getenv("MYSQL_DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
