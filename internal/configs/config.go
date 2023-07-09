package configs

import (
	"errors"
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

func (c *Config) GenerateTFConfig(workingDir string, fs afero.Afero) error {
	for label0, generates := range c.Generates {
		file := hclwrite.NewEmptyFile()
		body := file.Body()

		for index, generate := range generates {
			if !generate.When {
				continue
			}

			ctx := c.Ctx.NewChild()
			ctx.Variables = map[string]cty.Value{}

			if generate.ForEach != nil {
				for key, value := range generate.ForEach {
					ctx.Variables = map[string]cty.Value{
						"each": cty.ObjectVal(map[string]cty.Value{
							"key":   cty.StringVal(key),
							"value": value,
						}),
					}

					newBlock := body.AppendNewBlock(label0, generate.Labels[1:])
					newBlockBody := newBlock.Body()

					if err := appendBlockContent(newBlockBody, generate.Config, ctx); err != nil {
						return err
					}

					if index < len(generate.Config.Blocks)-1 {
						body.AppendNewline()
					}
				}
			} else {
				newBlock := body.AppendNewBlock(label0, generate.Labels[1:])
				newBlockBody := newBlock.Body()

				if err := appendBlockContent(newBlockBody, generate.Config, ctx); err != nil {
					return err
				}

				if index < len(generate.Config.Blocks)-1 {
					body.AppendNewline()
				}
			}
		}

		filename := filepath.Join(workingDir, fmt.Sprintf("%s.tf", label0))
		if err := fs.WriteFile(filename, file.Bytes(), 0644); err != nil {
			return err
		}
	}

	return nil
}

func appendBlockContent(file *hclwrite.Body, config *hclsyntax.Body, ctx *hcl.EvalContext) error {
	attributes := map[string]cty.Value{}
	for name, attribute := range config.Attributes {
		value, diags := attribute.Expr.Value(ctx)
		if diags.HasErrors() {
			return errors.Join(diags.Errs()...)
		}

		attributes[name] = value
	}

	for name, attribute := range attributes {
		file.SetAttributeValue(name, attribute)
	}

	for _, block := range config.Blocks {
		blockInBlock := file.AppendNewBlock(block.Type, []string{})
		blockInBlockBody := blockInBlock.Body()

		if err := appendBlockContent(blockInBlockBody, block.Body, ctx); err != nil {
			return err
		}
	}

	return nil
}
