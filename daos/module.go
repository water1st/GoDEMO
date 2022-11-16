package daos

import "go.uber.org/dig"

func RegisterDependencyInjection(container *dig.Container) {
	container.Provide(NewMemoryUserDAO)
}
