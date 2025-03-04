package main

import (
	"gin-boiler-plate/config"
	"gin-boiler-plate/database"
	"gin-boiler-plate/handler"
	"gin-boiler-plate/repository"
	"gin-boiler-plate/routes"
	"gin-boiler-plate/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()
	db, err := database.NewDatabase(cfg.Driver, cfg.DSN, true)
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	r := repository.NewRepositories(db)
	s := service.NewServices(r)
	h := handler.NewHandlers(s)

	e := gin.Default()
	routes.NewRoutes(e, h)
	e.Run(":3000")
}
