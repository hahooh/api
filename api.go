package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hahooh/flyerAPI/controllers"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	initRoutes(router)

	router.Run()
}

func initMiddlewares(router *gin.Engine) {
	router.Use(optionHandler)
	router.Use(gin.Recovery())
}

func initRoutes(router *gin.Engine) {
	baseController := controllers.BaseController{Router: router}

	userController := controllers.UserController{Controller: baseController}
	userController.InitRoutes()
}

func optionHandler(context *gin.Context) {
	if context.Request.Method == http.MethodOptions {
		context.JSON(204, nil)
		return
	}

	context.Next()
}
