package main

import (
	"github.com/axzed/go-jwt-demo/controllers"
	"github.com/axzed/go-jwt-demo/initializers"
	"github.com/axzed/go-jwt-demo/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	//initializers.SyncDataBase()
}

func main() {
	r := gin.Default()

	r.POST("/sigup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
