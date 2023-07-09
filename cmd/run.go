package cmd

import (
	"github.com/Yunsang-Jeong/terraforge/internal/app"

	"github.com/spf13/cobra"
)

type RunCmd struct{}

var runCmdFlags = map[string]flag{
	"debug": {
		_type:        "bool",
		shorten:      "d",
		description:  "[opt] enable debug mode",
		requirement:  false,
		defaultValue: false,
	},
	"wd": {
		_type:        "string",
		shorten:      "w",
		description:  "[opt] working directroy",
		requirement:  false,
		defaultValue: ".",
	},
	"cf": {
		_type:        "string",
		shorten:      "c",
		description:  "[opt] config file name",
		requirement:  false,
		defaultValue: "terraforge.hcl",
	},
}

func (r *RunCmd) Init() *cobra.Command {
	c := &cobra.Command{
		Use:   "run",
		Short: "Genraete Terraform configuration",
		Run: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")
			wd, _ := cmd.Flags().GetString("wd")
			cf, _ := cmd.Flags().GetString("cf")

			app.NewTerraforge(wd, cf, debug).Run()
		},
	}

	cobraFlagRegister(c, runCmdFlags)

	return c
}
