package configs

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Generate struct {
	Labels  []string
	When    bool
	ForEach hcl.Expression
	Config  *hclsyntax.Body
}

func decodeGenerateBlock(block *hclsyntax.Block, ctx *hcl.EvalContext) (*Generate, hcl.Diagnostics) {
	var diags hcl.Diagnostics

	g := &Generate{
		Labels: block.Labels,
		When:   true,
	}

	if attr, ok := block.Body.Attributes["when"]; ok {
		value, diag := attr.Expr.Value(ctx)
		diags = append(diags, diag...)

		if diag.HasErrors() {
			g.When = true
		} else if !value.IsNull() && value.False() {
			g.When = false
		}
	}

	if attr, ok := block.Body.Attributes["for_each"]; ok {
		g.ForEach = attr.Expr
	}

	for _, blockInBlock := range block.Body.Blocks {
		if blockInBlock.Type == "config" {
			g.Config = blockInBlock.Body
			break
		}
	}

	return g, diags
}
