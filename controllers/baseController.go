package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// BaseController -
type BaseController struct {
	Router         *gin.Engine
	ResponseData   interface{}
	HTTPStatusCode int
	ErrorMessage   string
}

// Route -
type Route struct {
	Mehod   string
	Path    string
	Handler func(*gin.Context)
}

// InitRoutes -
func (baseController *BaseController) InitRoutes(routeGroup string, routes []Route) {
	goupRouter := baseController.Router.Group(routeGroup)
	for _, route := range routes {
		if compareMethod(route, http.MethodGet) {
			goupRouter.GET(route.Path, route.Handler)
		} else if compareMethod(route, http.MethodPost) {
			goupRouter.POST(route.Path, route.Handler)
		} else if compareMethod(route, http.MethodPut) {
			goupRouter.PUT(route.Path, route.Handler)
		} else {
			panic("Unresigered Method. Please resigter")
		}
	}
}

// SuccessResponse -
func (baseController *BaseController) SuccessResponse(context *gin.Context) {
	if baseController.HTTPStatusCode == 0 {
		baseController.HTTPStatusCode = 200
	}
	context.JSON(baseController.HTTPStatusCode, baseController.ResponseData)
}

// ErrorResponse -
func (baseController *BaseController) ErrorResponse(context *gin.Context) {
	if baseController.HTTPStatusCode == 0 {
		baseController.HTTPStatusCode = 500
	}
	context.JSON(baseController.HTTPStatusCode, map[string]string{"errorMessage": baseController.ErrorMessage})
}

func compareMethod(route Route, method string) bool {
	return strings.ToUpper(route.Mehod) == strings.ToUpper(method)
}
