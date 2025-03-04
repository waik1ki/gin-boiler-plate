package routes

import (
	"gin-boiler-plate/handler"

	"github.com/gin-gonic/gin"
)

func NewRoutes(e *gin.Engine, handler *handler.Handlers) {
	newUserRoutes(e.Group("/users"), handler.User)
}

func newUserRoutes(g *gin.RouterGroup, handler *handler.UserHandler) {
	g.POST("", handler.CreateUser)
	g.GET("", handler.GetUsers)
	g.GET("/:id", handler.GetUser)
	g.PUT("/:id", handler.UpdateUser)
	g.DELETE("/:id", handler.DeleteUser)
}
