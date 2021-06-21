package main

import (
	"github.com/gin-gonic/gin"
	"server/routes"
)

func main(){
	//config.ConnectionToDatabase()

	r := gin.Default()
	routes.UserRoute(r)


	r.Run()
}
