package configs

import (
	"errors"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

type Generate struct {
	Labels  []string
	When    bool
	ForEach map[string]cty.Value
	Config  *hclsyntax.Body
}

func decodeGenerateBlock(block *hclsyntax.Block, ctx *hcl.EvalContext) (*Generate, error) {
	g := &Generate{
		Labels: block.Labels,
		When:   true,
	}

	if attr, ok := block.Body.Attributes["when"]; ok {
		value, diags := attr.Expr.Value(ctx)
		if diags.HasErrors() {
			g.When = true
		} else if !value.IsNull() && value.False() {
			g.When = false
		}
	}

	if attr, ok := block.Body.Attributes["for_each"]; ok {
		forEach, diags := attr.Expr.Value(ctx)
		if diags.HasErrors() {
			return nil, errors.Join(diags.Errs()...)
		}

		g.ForEach = map[string]cty.Value{}

		switch t := forEach.Type(); {
		case t.IsObjectType():
			for key, value := range forEach.AsValueMap() {
				g.ForEach[key] = value
			}
		case t.IsListType() || t.IsSetType():
			for _, value := range forEach.AsValueSlice() {
				g.ForEach[value.AsString()] = value
			}
		default:
			return nil, errors.New("for_each type must be object or set(string)")
		}
	}

	for _, blockInBlock := range block.Body.Blocks {
		if blockInBlock.Type == "config" {
			g.Config = blockInBlock.Body
			break
		}
	}

	return g, nil
}
