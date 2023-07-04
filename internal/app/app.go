package app

import (
	"github.com/Yunsang-Jeong/terraforge/internal/logger"
)

type terraforge struct {
	lg       logger.Logger
	metadata map[string]string
}

func NewTerraforge(debug bool) *terraforge {
	return &terraforge{
		lg:       logger.NewSimpleLogger(debug),
		metadata: map[string]string{},
	}
}

func (app *terraforge) Run(configFile string) error {
	if err := app.parseConfig(configFile); err != nil {
		return err
	}

	// if err := app.generateAWSProvider("provider.tf"); err != nil {
	// 	return err
	// }

	// if err := app.generateVariable("variable.tf"); err != nil {
	// 	return err
	// }

	// mytf()
	return nil
}
