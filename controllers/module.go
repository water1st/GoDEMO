package controllers

import (
	"go.uber.org/dig"
	"log"
)

func RegisterDependencyInjection(container *dig.Container) {
	err := container.Provide(newUserController)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
