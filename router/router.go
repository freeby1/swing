package router

import (
	"ginvue/controller"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/ping", controller.Register)
	return r
}
