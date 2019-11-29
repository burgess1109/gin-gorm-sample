package server

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	messageDB "gin-gorm-sample/adapter/database/message"
	userDB "gin-gorm-sample/adapter/database/user"
	messageWeb "gin-gorm-sample/adapter/web/message"
	userWeb "gin-gorm-sample/adapter/web/user"
	"gin-gorm-sample/application/message"
	messagePort "gin-gorm-sample/application/message/port"
	"gin-gorm-sample/application/user"
	userPort "gin-gorm-sample/application/user/port"
	"gin-gorm-sample/database"
)

type Service struct {
	user    userPort.Web
	message messagePort.Web
}

type Initiator struct {
	db      *gorm.DB
	service Service
	router  *gin.Engine
}

func Init() *Initiator {
	return &Initiator{}
}

func (b *Initiator) InitConfig() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath("./configs")
	viper.SetConfigType("yml")
	configFile := fmt.Sprintf("config.%s", os.Getenv("ENV"))
	viper.SetConfigName(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Config file get error:" + err.Error())
	}
}

func (b *Initiator) InitDB() {
	db, err := database.ConnectMySQL()
	if err != nil {
		log.Fatal("DB connect error:" + err.Error())
	}

	b.db = db

	if viper.GetBool("mysql.debug.mode") {
		b.db.LogMode(true)
	}
}

func (b *Initiator) InitService() {
	userRepo := userDB.NewRepository(b.db)
	userService := user.NewService(userRepo)

	messageRepo := messageDB.NewRepository(b.db)
	messageService := message.NewService(messageRepo, userRepo)

	b.service = Service{
		user:    userService,
		message: messageService,
	}
}

func (b *Initiator) InitRouter() {
	b.router = gin.Default()

	userRouter := userWeb.NewRouter(b.service.user)
	userRouter.SetRoutes(b.router)

	messageRouter := messageWeb.NewRouter(b.service.message)
	messageRouter.SetRoutes(b.router)
}

func (b *Initiator) GetRouter() *gin.Engine {
	return b.router
}

func (b *Initiator) GetDB() *gorm.DB {
	return b.db
}
