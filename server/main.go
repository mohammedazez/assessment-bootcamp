package main

import (
	"github.com/gin-gonic/gin"
	"server/handler"
	"server/routes"
)

func main(){
	//config.ConnectionToDatabase()

	r := gin.Default()
	r.Use(handler.CORSMiddleware())

	routes.UserRoute(r)
	routes.ListRoute(r)

	r.Run()
}
