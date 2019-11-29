package user

import (
	"github.com/gin-gonic/gin"

	userPort "gin-gorm-sample/application/user/port"
)

type Router struct {
	userService userPort.Web
}

func NewRouter(us userPort.Web) Router {
	return Router{
		userService: us,
	}
}

func (r Router) SetRoutes(g *gin.Engine) {
	g.GET("/users/:id", r.get)
	g.GET("/users", r.getAll)
	g.POST("/users", r.post)
	g.PATCH("/users/:id", r.patch)
	g.DELETE("/users/:id", r.delete)
}
