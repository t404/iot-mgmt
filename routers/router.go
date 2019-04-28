package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"iot-mgmt/controllers/v1"
	_ "iot-mgmt/docs"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	// 先实例化 再用
	userController := new(v1.UserController)

	apiv1.GET("/users", userController.List)
	apiv1.GET("/user/:id", userController.Find)

	//apiv1.Use(middlewares.Evaluation())
	{
		apiv1.POST("/users", userController.Create)
		apiv1.DELETE("/users", userController.Delete)
		apiv1.PUT("/user/:id", userController.Update)
	}
	return r
}
