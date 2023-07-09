package app

import (
	"github.com/Yunsang-Jeong/terraforge/internal/configs"
	"github.com/Yunsang-Jeong/terraforge/internal/logger"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
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

type generateBlock struct {
	labels []string
	config *hclsyntax.Block
}

type metadata map[string]cty.Value

type generateBlocks map[string][]generateBlock

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
	parser := configs.NewParser(app.wd, app.fs)

	config, diag := parser.LoadConfigFile(app.cf)
	if diag.HasErrors() {
		app.lg.Error(diag.Error())
		return nil
	}

	diag = config.GenerateTFConfig(app.wd, app.fs)
	if diag.HasErrors() {
		app.lg.Error(diag.Error())
		return nil
	}

	return nil
}
