package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDataBase() {
	dns := "root:3482@tpc(127.0.0.1:3306)/taskgo"
	connection, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Error al conectar con la base de datos..")
	}

	db = connection

	db.Logger = db.Logger.LogMode(logger.Info)
}

func GetDB() *gorm.DB {
	return db
}
