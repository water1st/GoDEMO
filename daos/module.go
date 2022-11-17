package daos

import "go.uber.org/dig"

func RegisterDependencyInjectionWithMemory(container *dig.Container) {
	container.Provide(newMemoryUserDAO)
}

func RegisterDependencyInjectionWithMySQL(container *dig.Container, config func(options *MySQLOptions)) {
	var options = MySQLOptions{}
	config(&options)
	container.Provide(func() MySQLOptions {
		return options
	})
	container.Provide(newMySQLUserDAO)
	container.Provide(newConnectionFactory)
}
