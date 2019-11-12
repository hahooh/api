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
	router.Use(setCORS)
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

func setCORS(context *gin.Context) {
	setContextHeader(context, "Access-Control-Allow-Origin", "*")
	setContextHeader(context, "Access-Control-Allow-Credentials", "true")
	setContextHeader(context, "Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, ApiKey")
	setContextHeader(context, "Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	context.Next()
}

func setContextHeader(context *gin.Context, key string, value string) {
	context.Writer.Header().Set(key, value)
}
