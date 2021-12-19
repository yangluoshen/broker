package main

import (
	"log"

	"github.com/yangluoshen/broker/app"
	"github.com/yangluoshen/broker/di"
)

func main() {
	c := di.GetContainerByEnv()

	err := c.Invoke(func(a *app.HttpApp) error {
		return a.Run()
	})

	if err != nil {
		log.Fatalf("fail to serve, err: %v", err)
	}
}
