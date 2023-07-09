package main

import (
	"os"

	"github.com/Yunsang-Jeong/terraforge/internal/app"
)

func main() {
	app := app.NewTerraforge("example/dev", "terraforge.hcl", true)
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
}
