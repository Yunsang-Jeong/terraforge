package configs

import (
	"errors"
	"path/filepath"

	"github.com/Yunsang-Jeong/terraforge/internal/logger"
	"github.com/Yunsang-Jeong/terraforge/internal/utils"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	tflang "github.com/hashicorp/terraform/lang"
	"github.com/spf13/afero"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

type Parser struct {
	lg logger.Logger
	fs afero.Afero
	p  *hclparse.Parser
	wd string
}

func NewParser(lg logger.Logger, workingDir string, fs afero.Afero) *Parser {
	return &Parser{
		lg: lg,
		fs: fs,
		p:  hclparse.NewParser(),
		wd: workingDir,
	}
}

func (p *Parser) LoadConfigFile(configFile string) (*Config, error) {
	body, err := p.loadHCLFile(configFile)
	if err != nil {
		p.lg.Error("fail to load config file")
		return nil, nil
	}

	blocks := map[string][]*hclsyntax.Block{}
	for _, block := range body.Blocks {
		blocks[block.Type] = append(blocks[block.Type], block)
	}

	config := &Config{
		Ctx: &hcl.EvalContext{
			Variables: map[string]cty.Value{},
			Functions: map[string]function.Function{},
		},
	}

	tfscope := tflang.Scope{
		BaseDir: filepath.Dir("."),
	}
	for k, v := range tfscope.Functions() {
		config.Ctx.Functions[k] = v
	}

	metadata := map[string]cty.Value{}

	if len(blocks["metafile"]) > 1 {
		p.lg.Error("no more than one metadata block can exist")
		return nil, errors.New("too many metafile block")
	}

	for _, block := range blocks["metafile"] {
		metafileBlock, err := decodeMetafileBlock(block, nil)
		if err != nil {
			p.lg.Error("fail to decode metafile block")
			return nil, err
		}

		for _, metafile := range metafileBlock.Paths {
			metafileBody, err := p.loadHCLFile(metafile)
			if err != nil {
				p.lg.Error("fail to load metafile")
				return nil, err
			}

			for name, attr := range metafileBody.Attributes {
				value, diags := attr.Expr.Value(nil)
				if diags.HasErrors() {
					return nil, errors.Join(diags.Errs()...)
				}

				metadata[name] = value
			}
		}
	}
	config.Ctx.Variables["metadata"] = cty.ObjectVal(metadata)

	config.Generates = map[string][]*Generate{}

	for _, block := range blocks["generate"] {
		generateBlock, err := decodeGenerateBlock(block, config.Ctx)
		if err != nil {
			p.lg.Error("fail to decode generate block")
			return nil, err
		}

		config.Generates[block.Labels[0]] = append(config.Generates[block.Labels[0]], generateBlock)
	}

	return config, nil
}

func (p *Parser) loadHCLFile(filename string) (*hclsyntax.Body, error) {
	path, err := utils.GetSomethingPathInParents(p.wd, filename, true)
	if err != nil {
		p.lg.Error("failed to get hcl file path", "err", err, "filename", filename)
		return nil, err
	}

	src, err := p.fs.ReadFile(filepath.Join(p.wd, path))
	if err != nil {
		p.lg.Error("failed to get read hcl path", "err", err, "filename", filename)
		return nil, err
	}

	file, diag := p.p.ParseHCL(src, filename)
	if file == nil || file.Body == nil || diag.HasErrors() {
		p.lg.Error("fail to load config file")
		return hcl.EmptyBody().(*hclsyntax.Body), errors.Join(diag.Errs()...)
	}

	return file.Body.(*hclsyntax.Body), nil
}
