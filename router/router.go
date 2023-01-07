package router

import "github.com/gin-gonic/gin"

// Controller contains the controller methods
type Controller interface {
	Ping(*gin.Context)
	GetSample(*gin.Context)
}

// New creates the router
func New(controller Controller) *gin.Engine {
	r := gin.New()

	v1 := r.Group("/v1")
	v1.GET("/ping", controller.Ping)
	v1.GET("/sample", controller.GetSample)

	return r
}
