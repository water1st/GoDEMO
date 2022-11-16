package main

import (
	"GoDEMO/controllers"
	"GoDEMO/daos"
	"GoDEMO/routers"
	"GoDEMO/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {
	var host = gin.Default()
	var contianer = createContianer()

	routers.RegisterUserRouter(host, contianer)

	host.Run()
}

func createContianer() *dig.Container {
	var contianer = dig.New()
	daos.RegisterDependencyInjection(contianer)
	services.RegisterDependencyInjection(contianer)
	controllers.RegisterDependencyInjection(contianer)

	return contianer
}
