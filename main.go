package main

import (
	"GoDEMO/controllers"
	"GoDEMO/daos"
	"GoDEMO/routers"
	"GoDEMO/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"log"
	"os"
)

func main() {
	var host = gin.Default()
	var contianer = createContianer()

	routers.RegisterUserRouter(host, contianer)

	err := host.Run()
	if err != nil {
		println(err.Error())
	}
}

func createContianer() *dig.Container {
	var contianer = dig.New()
	contianer.Provide(func() *log.Logger {
		return log.Default()
	})
	var providerName = os.Getenv("DAO_PROVIDER")
	if providerName == "mysql" {
		daos.RegisterDependencyInjectionWithMySQL(contianer, func(options *daos.MySQLOptions) {
			//options.ConnectionString = "root:123456@tcp(localhost:3306)/test?charset=utf8"
			options.ConnectionString = os.Getenv("CONNECTION_STRING")
		})
	} else {
		daos.RegisterDependencyInjectionWithMemory(contianer)
	}

	services.RegisterDependencyInjection(contianer)
	controllers.RegisterDependencyInjection(contianer)

	return contianer
}
