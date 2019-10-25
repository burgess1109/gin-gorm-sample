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
	"gin-gorm-sample/message"
	"gin-gorm-sample/user"
)

var db *gorm.DB
var routes *gin.Engine
var services Services
var err error

type Services struct {
	user    user.UserServiceInterface
	message message.MessageServiceInterface
}

func main() {
	initConfig()

	initDB()
	defer db.Close()

	port, exist := os.LookupEnv("WEB_PORT")
	if !exist {
		port = "8080"
	}

	initServices()

	routes = gin.Default()
	initRoutes()

	routes.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err = routes.Run(":" + port)
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
	if err != nil {
		log.Fatal("Config file get error:" + err.Error())
	}
}

func initDB() {
	db, err = database.ConnectMySQL()
	if err != nil {
		log.Fatal("DB connect error:" + err.Error())
	}

	if viper.GetBool("mysql.debug.mode") {
		db.LogMode(true)
	}
}

func initServices() {
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)

	messageRepo := message.NewMessageRepository(db)
	messageService := message.NewMessageService(messageRepo)

	services.user = userService
	services.message = messageService
}

func initRoutes() {
	routes.GET("users", services.user.Get)
	routes.GET("users/:id", services.user.GetByID)
	routes.POST("users", services.user.Create)
	routes.PATCH("users/:id", services.user.Update)
	routes.DELETE("users/:id", services.user.Delete)

	routes.GET("messages", services.message.Get)
	routes.GET("messages/:id", services.message.GetByID)
	routes.POST("messages", services.message.Create)
	routes.PATCH("messages/:id", services.message.Update)
	routes.DELETE("messages/:id", services.message.Delete)
}
