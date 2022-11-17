package services

import (
	"go.uber.org/dig"
	"log"
)

func RegisterDependencyInjection(container *dig.Container) {
	err := container.Provide(newUserService)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
