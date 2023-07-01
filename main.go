package main

import (
	"github.com/Yunsang-Jeong/terraforge/internal/app"
)

func main() {
	app := app.NewApp(true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
