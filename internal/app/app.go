package app

import (
	"fmt"
	"os"
	"path/filepath"

	tflang "github.com/hashicorp/terraform/lang"

	"github.com/Yunsang-Jeong/terraforge/internal/logger"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

type terraforge struct {
	lg       logger.Logger
	metadata map[string]string
	parser   *hclparse.Parser
	ctx      *hcl.EvalContext
}

func NewTerraforge(debug bool) *terraforge {
	return &terraforge{
		lg:       logger.NewSimpleLogger(debug),
		metadata: map[string]string{},
		parser:   hclparse.NewParser(),
		ctx: &hcl.EvalContext{
			Variables: map[string]cty.Value{},
			Functions: map[string]function.Function{},
		},
	}
}

func (app *terraforge) Run(configFile string) error {
	tfscope := tflang.Scope{
		BaseDir: filepath.Dir("."),
	}
	for k, v := range tfscope.Functions() {
		app.ctx.Functions[k] = v
	}

	generateBlocks, err := app.parse(configFile)
	if err != nil {
		return err
	}

	for blockType, blocks := range generateBlocks {
		file := hclwrite.NewEmptyFile()
		body := file.Body()

		for index, block := range blocks {
			newBlock := body.AppendNewBlock(blockType, block.labels[1:])
			newBlockBody := newBlock.Body()

			if err := generate(newBlockBody, block.config, app.ctx); err != nil {
				return err
			}

			if index < len(blocks)-1 {
				body.AppendNewline()
			}
		}

		writer, err := os.Create(fmt.Sprintf("./%s.tf", blockType))
		if err != nil {
			return err
		}
		defer writer.Close()

		file.WriteTo(writer)
	}

	return nil
}
