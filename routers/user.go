package routers

import (
	"GoDEMO/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func RegisterUserRouter(engine *gin.Engine, container *dig.Container) {
	var err = container.Invoke(func(userController *controllers.IUserController) {
		var controller = *userController
		engine.GET("/user", controller.GetAll)
		engine.GET("/user/:id", controller.GetById)
		engine.POST("/user", controller.Add)
		engine.PUT("/user", controller.Update)
		engine.DELETE("/user/:id", controller.Delete)
	})

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
