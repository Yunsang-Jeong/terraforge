package app

import (
	"bytes"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AWSProviders []AWSProvider `yaml:"aws_provider,omitempty"`
	S3Backend    S3Backend     `yaml:"s3_backend,omitempty"`
	Variables    []Variable    `yaml:"variable,omitempty"`
}

type AWSProvider struct {
	Name              string            `yaml:"name", hcl:"name"`
	Region            string            `yaml:"region"`
	DefaultTags       map[string]string `yaml:"default_tags,omitempty"`
	AssumeRoleArn     string            `yaml:"asssume_role_arn,omitempty"`
	AssumeSessionName string            `yaml:"asssume_session_name,omitempty"`
	Condtion          []bool            `yaml:"condition,omitempty"`
}

type S3Backend struct {
	Bucket string `yaml:"bucket"`
	Key    string `yaml:"key"`
	Region string `yaml:"region"`
}

type Variable struct {
	Name        string      `yaml:"name"`
	Type        string      `yaml:"type"`
	Description string      `yaml:"description,omitempty"`
	Default     interface{} `yaml:"default,omitempty"`
}

// equalMetadata
func (app *terraforge) parseConfigAndSave(rawConfig string) error {
	configFunc := template.FuncMap{
		"metadata":               app.metadataTmplFunc,
		"metadataIfEqual":        app.metadataIfEqualTmplFunc,
		"need_multiple_provider": app.needMultipleProviderTmplFunc,
	}

	tmpl, err := template.New("config").Funcs(configFunc).Parse(rawConfig)
	if err != nil {
		app.lg.Error("fail to initialize to execute config-functions", "rawConfig", rawConfig, "err", err.Error())
		return err
	}

	var executedRawConfig bytes.Buffer
	if err := tmpl.Execute(&executedRawConfig, nil); err != nil {
		app.lg.Error("fail to execute config-functions", "rawConfig", rawConfig, "err", err.Error())
		return err
	}

	config := Config{}
	if err := yaml.Unmarshal(executedRawConfig.Bytes(), &config); err != nil {
		app.lg.Error("fail to parse config-function-executed config ", "executedRawConfig", executedRawConfig.String(), "err", err.Error())
		return err
	}

	app.config = config

	return nil
}

func (a *terraforge) metadataTmplFunc(key string) string {
	return a.metadata[key]
}

func (a *terraforge) metadataIfEqualTmplFunc(key string, compreMetadata string, compareValue string) string {
	if a.metadata[compreMetadata] == compareValue {
		return a.metadata[key]
	}
	return ""
}

func (a *terraforge) needMultipleProviderTmplFunc() bool {
	return false
}
