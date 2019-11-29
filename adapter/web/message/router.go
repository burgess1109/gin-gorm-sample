package message

import (
	"github.com/gin-gonic/gin"

	messagePort "gin-gorm-sample/application/message/port"
)

type Router struct {
	messageService messagePort.Web
}

func NewRouter(ms messagePort.Web) Router {
	return Router{
		messageService: ms,
	}
}

func (r Router) SetRoutes(g *gin.Engine) {
	g.GET("/messages/:id", r.get)
	g.GET("/messages", r.getAll)
	g.POST("/messages", r.post)
	g.PATCH("/messages/:id", r.patch)
	g.DELETE("/messages/:id", r.delete)
}
