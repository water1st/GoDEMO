package services

import "go.uber.org/dig"

func RegisterDependencyInjection(container *dig.Container) {
	container.Provide(NewUserService)
}
