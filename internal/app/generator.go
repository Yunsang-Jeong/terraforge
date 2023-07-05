package app

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

type generateBlock struct {
	labels []string
	config *hclsyntax.Block
}

func generate(body *hclwrite.Body, config *hclsyntax.Block, ctx *hcl.EvalContext) error {
	attributes, err := evalAttributes(config.Body.Attributes, ctx)
	if err != nil {
		return err
	}

	for name, attribute := range attributes {
		body.SetAttributeValue(name, attribute)
	}

	for _, block := range config.Body.Blocks {
		blockInBlock := body.AppendNewBlock(block.Type, []string{})
		blockInBlockBody := blockInBlock.Body()

		if err := generate(blockInBlockBody, block, ctx); err != nil {
			return err
		}
	}
	return nil
}
