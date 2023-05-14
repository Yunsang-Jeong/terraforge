package main

import (
	"fmt"

	"github.com/Yunsang-Jeong/terraforge/internal/config"
)

func main() {
	config := config.NewConfig()
	config.LoadConfigFromFile()

	block := config.GetConfig()
	fmt.Printf("%v", block)
}
