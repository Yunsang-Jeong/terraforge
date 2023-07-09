package configs

import (
	"fmt"
	"path/filepath"

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
	fs afero.Afero
	p  *hclparse.Parser
	wd string
}

func NewParser(workingDir string, fs afero.Afero) *Parser {
	return &Parser{
		fs: fs,
		p:  hclparse.NewParser(),
		wd: workingDir,
	}
}

func (p *Parser) LoadConfigFile(configFile string) (*Config, hcl.Diagnostics) {
	var diags hcl.Diagnostics

	body, diag := p.loadHCLFile(configFile)
	if diag.HasErrors() {
		diags = append(diags, diag...)
		return nil, diags
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
		diags = diags.Append(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Too many metafile block",
			Detail:   "No more than one metadata block can exist",
		})
	}

	for _, block := range blocks["metafile"] {
		metafileBlock, diag := decodeMetafileBlock(block, nil)
		diags = append(diags, diag...)

		for _, metafile := range metafileBlock.Paths {

			metafileBody, diag := p.loadHCLFile(metafile)
			diags = append(diags, diag...)
			if metafileBody == nil {
				return nil, diags
			}

			for name, attr := range metafileBody.Attributes {
				value, diag := attr.Expr.Value(nil)
				diags = append(diags, diag...)

				metadata[name] = value
			}
		}
	}
	config.Ctx.Variables["metadata"] = cty.ObjectVal(metadata)

	config.Generates = map[string][]*Generate{}

	for _, block := range blocks["generate"] {
		generateBlock, diag := decodeGenerateBlock(block, config.Ctx)
		diags = append(diags, diag...)

		config.Generates[block.Labels[0]] = append(config.Generates[block.Labels[0]], generateBlock)
	}

	return config, diags
}

func (p *Parser) loadHCLFile(filename string) (*hclsyntax.Body, hcl.Diagnostics) {
	var diags hcl.Diagnostics

	path, err := utils.GetSomethingPathInParents(p.wd, filename, true)
	if err != nil {
		return nil, diags.Append(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Failed to get hcl file path",
			Detail:   fmt.Sprintf("The file %s could not be read.", filename),
		})
	}

	src, err := p.fs.ReadFile(filepath.Join(p.wd, path))
	if err != nil {
		return nil, diags.Append(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Failed to read hcl file",
			Detail:   fmt.Sprintf("The file %s could not be read.", filename),
		})
	}

	file, diags := p.p.ParseHCL(src, filename)
	if file == nil || file.Body == nil {
		return hcl.EmptyBody().(*hclsyntax.Body), diags
	}

	return file.Body.(*hclsyntax.Body), diags
}
