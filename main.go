package main

import (
	"os"

	"github.com/Yunsang-Jeong/terraforge/internal/app"
)

func main() {
	app := app.NewTerraforge(true)
	if err := app.Run("terraforge.yaml"); err != nil {
		os.Exit(1)
	}
}
