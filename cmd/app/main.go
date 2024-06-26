package main

import (
	"github.com/gsxhnd/go-api-template/di"
)

func main() {
	app, err := di.InitApp()
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
