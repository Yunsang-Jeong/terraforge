package app

import (
	"fmt"

	"github.com/Yunsang-Jeong/terraforge/internal/logger"
)

type app struct {
	lg       logger.Logger
	config   *Config
	metadata map[string]string
}

func NewApp(debug bool) *app {
	return &app{
		lg:       logger.NewSimpleLogger(debug),
		config:   &Config{},
		metadata: map[string]string{},
	}
}

func (a *app) Run() error {
	if err := a.parseMetadata(); err != nil {
		return err
	}
	a.lg.Debug("parsed metadata", "metadata", a.metadata)

	if err := a.parseConfig(); err != nil {
		return err
	}
	a.lg.Debug("parsed aws_provider", "aws_provider", fmt.Sprintf("%+v", a.config.AWSProviders))
	a.lg.Debug("parsed s3_backend", "s3_backend", fmt.Sprintf("%+v", a.config.S3Backend))
	a.lg.Debug("parsed variable", "variable", fmt.Sprintf("%+v", a.config.Variables))

	if err := a.generateTerraformConfiguration(); err != nil {
		return err
	}

	fmt.Print(a.config)

	return nil
}
