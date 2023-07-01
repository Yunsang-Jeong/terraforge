package app

import (
	"bytes"
	"text/template"

	"github.com/Yunsang-Jeong/terraforge/internal/util"
	"gopkg.in/yaml.v3"
)

type Config struct {
	AWSProviders []AWSProvider `yaml:"aws_provider,omitempty"`
	S3Backend    S3Backend     `yaml:"s3_backend,omitempty"`
	Variables    []Variable    `yaml:"variable,omitempty"`
}

type AWSProvider struct {
	Name              string            `yaml:"name"`
	Region            string            `yaml:"region"`
	DefaultTags       map[string]string `yaml:"default_tags,omitempty"`
	AssumeRoleArn     string            `yaml:"asssume_role_arn,omitempty"`
	AssumeSessionName string            `yaml:"asssume_session_name,omitempty"`
	Condtion          []string          `yaml:"condition,omitempty"`
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

const configName = "terraforge.yaml"

func (a *app) parseConfig() error {
	configRelPath, err := util.GetSomethingPath(".", configName, true)
	if err != nil {
		a.lg.Error(err.Error(), "configName", configName)
		return err
	}

	rawConfig, err := util.LoadFileAsString(configRelPath)
	if err != nil {
		a.lg.Error(err.Error(), "configRelPath", configRelPath)
		return err
	}

	configFunc := template.FuncMap{
		"metadata":               a.metadataTmplFunc,
		"need_multiple_provider": a.needMultipleProviderTmplFunc,
	}

	tmpl, err := template.New("config").Funcs(configFunc).Parse(rawConfig)
	if err != nil {
		a.lg.Error(err.Error(), "rawConfig", rawConfig)
		return err
	}

	var executedRawConfig bytes.Buffer
	if err := tmpl.Execute(&executedRawConfig, nil); err != nil {
		a.lg.Error(err.Error(), "rawConfig", rawConfig)
		return err
	}

	config := Config{}
	if err := yaml.Unmarshal(executedRawConfig.Bytes(), &config); err != nil {
		a.lg.Error(err.Error(), "rawConfig", rawConfig)
		return err
	}

	a.config = &config

	return nil
}

//lint:ignore U1000 Ignore template function
func (a *app) metadataTmplFunc(target string) string {
	if value, ok := a.metadata[target]; ok {
		return value
	}

	return "NotFound"
}

//lint:ignore U1000 Ignore template function
func (a *app) needMultipleProviderTmplFunc() interface{} {
	return false
}
