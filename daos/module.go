package daos

import (
	"go.uber.org/dig"
	"log"
)

func RegisterDependencyInjectionWithMemory(container *dig.Container) {
	logError(container.Provide(newMemoryUserDAO))
}

func RegisterDependencyInjectionWithMySQL(container *dig.Container, config func(options *MySQLOptions)) {
	var options = MySQLOptions{}
	config(&options)
	logError(container.Provide(func() MySQLOptions {
		return options
	}))
	logError(container.Provide(newMySQLUserDAO))
	logError(container.Provide(newConnectionFactory))
}

func logError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
