package app

import (
	"github.com/Yunsang-Jeong/terraforge/internal/logger"
	"github.com/Yunsang-Jeong/terraforge/internal/util"
)

type terraforge struct {
	lg       logger.Logger
	config   Config
	metadata map[string]string
}

func NewTerraforge(debug bool) *terraforge {
	return &terraforge{
		lg:       logger.NewSimpleLogger(debug),
		config:   Config{},
		metadata: map[string]string{},
	}
}

func (app *terraforge) Run(configFile string) error {
	rawConfig, err := util.GetSomethingInParents(".", configFile)
	if err != nil {
		return err
	}

	if err := app.parseMetadataAndSave(rawConfig); err != nil {
		return err
	}

	if err := app.parseConfigAndSave(rawConfig); err != nil {
		return err
	}

	if err := app.generateAWSProvider("provider.tf"); err != nil {
		return err
	}

	if err := app.generateVariable("variable.tf"); err != nil {
		return err
	}

	return nil
}
