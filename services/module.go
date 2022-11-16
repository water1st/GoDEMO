package services

import "go.uber.org/dig"

func RegisterDependencyInjection(container *dig.Container) {
	err := container.Provide(NewUserService)
	if err != nil {
		println(err.Error())
	}
}
