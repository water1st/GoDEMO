package controllers

import (
	"go.uber.org/dig"
	"log"
)

func RegisterDependencyInjection(container *dig.Container) {
	err := container.Provide(NewUserController)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
