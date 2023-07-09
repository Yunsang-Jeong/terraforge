package main

import (
	"github.com/Yunsang-Jeong/terraforge/cmd"
)

func main() {
	run := &cmd.RunCmd{}

	cmd.RootCmd.AddCommand(run.Init())
	cmd.Execute()
}
