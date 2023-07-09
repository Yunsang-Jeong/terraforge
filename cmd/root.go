package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "terraforge",
	Short:         "Genraete Terraform configuration",
	Long:          "Genraete Terraform configuration",
	SilenceErrors: true,
}

func Execute() {
	// Remove help for root command
	RootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	// Remove shell completion
	RootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd:   true,
		DisableNoDescFlag:   true,
		DisableDescriptions: true,
		HiddenDefaultCmd:    true,
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
