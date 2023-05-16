package database

import (
	"investo-rest-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	//Jika panic, tolong periksan di root:password, kalau db kalian
	// ga pakai password langsung saja root:@/namadb
	connection, err := gorm.Open(mysql.Open("root:@/investo-auth"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
