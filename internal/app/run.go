package app

import (
	"github.com/Yunsang-Jeong/terraforge/internal/configs"
	"github.com/Yunsang-Jeong/terraforge/internal/logger"
	"github.com/hashicorp/hcl/v2"
	"github.com/spf13/afero"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

type terraforge struct {
	lg  logger.Logger
	ctx *hcl.EvalContext
	fs  afero.Afero
	wd  string
	cf  string
}

func NewTerraforge(workingDir string, configFile string, debug bool) *terraforge {
	return &terraforge{
		lg: logger.NewSimpleLogger(debug),
		ctx: &hcl.EvalContext{
			Variables: map[string]cty.Value{},
			Functions: map[string]function.Function{},
		},
		fs: afero.Afero{Fs: afero.OsFs{}},
		wd: workingDir,
		cf: configFile,
	}
}

func (app *terraforge) Run() error {
	parser := configs.NewParser(app.lg, app.wd, app.fs)

	config, err := parser.LoadConfigFile(app.cf)
	if err != nil {
		app.lg.Error("fail to load config file", "err", err)
		return nil
	}

	if err := config.GenerateTFConfig(app.wd, app.fs); err != nil {
		app.lg.Error("fail to generate terraform config", "err", err)
		return nil
	}

	return nil
}
