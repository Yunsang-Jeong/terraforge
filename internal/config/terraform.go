package config

type Block struct {
	ProviderBlocks []ProviderBlock `mapstructure:"provider"`
	VariableBlocks []VariableBlock `mapstructure:"variable"`
}

type ProviderBlock struct {
	Name        string                  `mapstructure:"name"`
	Region      string                  `mapstructure:"region"`
	AccessKey   string                  `mapstructure:"access_key"`
	SecretKey   string                  `mapstructure:"secret_key"`
	Profile     string                  `mapstructure:"profile"`
	DefaultTags map[string]string       `mapstructure:"default_tags"`
	AssumeRole  ProviderBlockAssumeRole `mapstructure:"assume_role"`
}

type ProviderBlockAssumeRole struct {
	RoleArn     string `mapstructure:"role_arn"`
	SessionName string `mapstructure:"session_name"`
}

type VariableBlock struct {
	Name        string      `mapstructure:"name"`
	Type        string      `mapstructure:"type"`
	Description string      `mapstructure:"description"`
	Default     interface{} `mapstructure:"default"`
}
