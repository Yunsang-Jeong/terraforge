package main

import (
	"os"

	"github.com/Yunsang-Jeong/terraforge/internal/app"
)

func main() {
	app := app.NewTerraforge(true)
	if err := app.Run("terraforge.hcl"); err != nil {
		os.Exit(1)
	}
}
