package controllers

import "go.uber.org/dig"

func RegisterDependencyInjection(container *dig.Container) {
	container.Provide(NewUserController)
}
