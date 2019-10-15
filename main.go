package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	"gin-gorm-sample/database"
)

var db *gorm.DB
var err error

func main() {
	initConfig()
	if err != nil {
		log.Fatal("Config file get error:" + err.Error())
	}

	initDB()
	if err != nil {
		log.Fatal("DB connect error:" + err.Error())
	}

	port, exist := os.LookupEnv("WEB_PORT")
	if !exist {
		port = "8080"
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Server run error:" + err.Error())
	}
}

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath("./configs")
	viper.SetConfigType("yml")
	configFile := fmt.Sprintf("config.%s", os.Getenv("ENV"))
	viper.SetConfigName(configFile)

	err = viper.ReadInConfig()
}

func initDB() {
	db, err = database.ConnectMySQL()
	if err != nil {
		return
	}

	if viper.GetBool("mysql.debug.mode") {
		db.LogMode(true)
	}
}
