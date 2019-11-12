package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController -
type UserController struct {
	Controller BaseController
}

// InitRoutes -
func (userController *UserController) InitRoutes() {
	routes := []Route{
		Route{
			Mehod:   http.MethodGet,
			Path:    "",
			Handler: userController.getUser,
		},
	}
	userController.Controller.InitRoutes("user", routes)
}

func (userController *UserController) getUser(context *gin.Context) {
	data := "hello world"
	userController.Controller.ResponseData = data
	userController.Controller.SuccessResponse(context)
}
