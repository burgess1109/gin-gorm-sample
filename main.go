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
	"gin-gorm-sample/messages"
	"gin-gorm-sample/users"
)

var db *gorm.DB
var routes *gin.Engine
var services Services
var err error

type Services struct {
	users    users.UserServiceInterface
	messages messages.MessageServiceInterface
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
	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo)

	messageRepo := messages.NewMessageRepository(db)
	messageService := messages.NewMessageService(messageRepo)

	services.users = userService
	services.messages = messageService
}

func initRoutes() {
	routes.GET("users", services.users.Get)
	routes.GET("users/:id", services.users.GetByID)
	routes.POST("users", services.users.Create)
	routes.PATCH("users/:id", services.users.Update)
	routes.DELETE("users/:id", services.users.Delete)

	routes.GET("messages", services.messages.Get)
	routes.GET("messages/:id", services.messages.GetByID)
	routes.POST("messages", services.messages.Create)
	routes.PATCH("messages/:id", services.messages.Update)
	routes.DELETE("messages/:id", services.messages.Delete)
}
