package tests

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUpTestDB() (*gorm.DB, error) {
	//MySQL
	dbName := "books_test"
	dbUser := "mysql"
	dbPassword := "mysql"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func EmptyDB(db *gorm.DB) {
	var tables []string
	db.Raw("SHOW TABLES").Scan(&tables)

	for _, table := range tables {
		db.Exec("DROP TABLE IF EXISTS " + table)
	}
}

func SetUpTestRouter() *gin.Engine {
	return gin.Default()
}
