package main

import (
	"log"
	"os"

	"gin-gorm-sample/server"
)

func main() {
	initiator := server.Init()
	initiator.InitConfig()
	initiator.InitDB()
	initiator.InitService()
	initiator.InitRouter()

	db := initiator.GetDB()
	defer db.Close()

	router := initiator.GetRouter()

	port, exist := os.LookupEnv("WEB_PORT")
	if !exist {
		port = "8080"
	}

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Server run error:" + err.Error())
	}
}
