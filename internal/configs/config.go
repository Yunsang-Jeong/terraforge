package configs

import (
	"fmt"
	"path/filepath"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/spf13/afero"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	Metafile  *Metafile
	Generates map[string][]*Generate
	Ctx       *hcl.EvalContext
}

func (c *Config) GenerateTFConfig(workingDir string, fs afero.Afero) hcl.Diagnostics {
	var diags hcl.Diagnostics

	for label0, generates := range c.Generates {
		file := hclwrite.NewEmptyFile()
		body := file.Body()

		for index, generate := range generates {
			if !generate.When {
				continue
			}

			newBlock := body.AppendNewBlock(label0, generate.Labels[1:])
			newBlockBody := newBlock.Body()

			diag := appendBlockContent(newBlockBody, generate.Config, c.Ctx)
			diags = append(diags, diag...)
			if diags.HasErrors() {
				return diags
			}

			if index < len(generate.Config.Blocks)-1 {
				body.AppendNewline()
			}
		}

		filename := filepath.Join(workingDir, fmt.Sprintf("%s.tf", label0))
		if err := fs.WriteFile(filename, file.Bytes(), 0644); err != nil {
			return diags.Append(&hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "F metafile block",
				Detail:   "metafile block must have path attirbute only",
			})
		}
	}

	return nil
}

func appendBlockContent(file *hclwrite.Body, config *hclsyntax.Body, ctx *hcl.EvalContext) hcl.Diagnostics {
	var diags hcl.Diagnostics

	attributes := map[string]cty.Value{}
	for name, attribute := range config.Attributes {
		value, diag := attribute.Expr.Value(ctx)
		diags = append(diags, diag...)
		if diags.HasErrors() {
			return diags
		}

		attributes[name] = value
	}

	for name, attribute := range attributes {
		file.SetAttributeValue(name, attribute)
	}

	for _, block := range config.Blocks {
		blockInBlock := file.AppendNewBlock(block.Type, []string{})
		blockInBlockBody := blockInBlock.Body()

		diag := appendBlockContent(blockInBlockBody, block.Body, ctx)
		diags = append(diags, diag...)
		if diags.HasErrors() {
			return diags
		}
	}
	return diags
}
