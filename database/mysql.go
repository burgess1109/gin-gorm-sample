package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import mysql driver initial
	"github.com/spf13/viper"
)

func ConnectMySQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s",
		viper.GetString("mysql.name"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.dsn"),
	)

	return gorm.Open("mysql", dsn)
}
