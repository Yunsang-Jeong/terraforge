package config

import (
	"path/filepath"

	logger "github.com/Yunsang-Jeong/terraforge/internal/logger"
	"github.com/spf13/viper"
)

const (
	configName = "terraforge"
	configType = "yaml"
)

type config struct {
	viper  *viper.Viper
	logger logger.Logger
	data   Block
}

func NewConfig() *config {
	return &config{
		viper:  viper.NewWithOptions(),
		logger: logger.NewSimpleLogger(true),
	}
}

func (c *config) getParentDirs(base string, dirs []string) ([]string, error) {
	current, err := filepath.Abs(base)
	if err != nil {
		c.logger.Error("fail to get absolute path", "err", err.Error())
		return nil, err
	}

	if current != "/" {
		dirs = append(dirs, current)
		parent := filepath.Dir(current)
		return c.getParentDirs(parent, dirs)
	}

	return dirs, nil
}

func (c *config) LoadConfigFromFile() error {
	c.viper.SetConfigName(configName)
	c.viper.SetConfigType(configType)

	dirs, err := c.getParentDirs(".", []string{})
	if err != nil {
		c.logger.Error("fail to get parents directory list", "err", err.Error())
		return err
	}
	c.logger.Debug("success to get parents directory list", "dirs", dirs)

	for _, dir := range dirs {
		c.viper.AddConfigPath(dir)
	}

	if err := c.viper.ReadInConfig(); err != nil {
		c.logger.Error("fail to read config", "err", err.Error())
		return err
	}

	if err := c.viper.Unmarshal(&c.data); err != nil {
		c.logger.Error("fail to unmarshal config", "err", err.Error())
		return err
	}

	c.logger.Debug("success to load config file")

	return nil
}

func (c *config) GetConfig() *Block {
	return &c.data
}
