package services

import (
	"go.uber.org/dig"
	"log"
)

func RegisterDependencyInjection(container *dig.Container) {
	err := container.Provide(NewUserService)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
