package di

import (
	"os"

	"github.com/yangluoshen/broker/app"
	"go.uber.org/dig"
)

func GetContainerByEnv() *dig.Container {

	switch os.Getenv("APP_ENV") {
	case "test":
		return newTestContainer()
	case "prod":
		return newProdContainer()
	default:
		panic("need APP_ENV")
	}
}

func newProdContainer() *dig.Container {
	c := dig.New()
	c.Provide(app.NewHttpApp)

	return c
}

func newTestContainer() *dig.Container {
	c := dig.New()
	c.Provide(app.NewHttpApp)

	return c
}
