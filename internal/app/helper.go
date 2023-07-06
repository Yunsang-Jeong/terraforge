package app

import (
	"fmt"

	"github.com/Yunsang-Jeong/terraforge/internal/util"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

func getHclBodyFromFile(filename string, parser *hclparse.Parser) (*hclsyntax.Body, error) {
	path, err := util.GetSomethingPathInParents(".", filename, true)
	if err != nil {
		return nil, err
	}

	file, diag := parser.ParseHCLFile(path)
	if diag.HasErrors() {
		return nil, diag.Errs()[0]
	}

	body := file.Body.(*hclsyntax.Body)
	if diag.HasErrors() {
		return nil, diag.Errs()[0]
	}

	return body, nil
}

func getAttribute(attributes hclsyntax.Attributes, target string, ctx *hcl.EvalContext, errorNotExist bool) (cty.Value, error) {
	attribute, ok := attributes[target]
	if !ok {
		if errorNotExist {
			return cty.NilVal, fmt.Errorf("%s is not in attributes", target)
		} else {
			return cty.NilVal, nil
		}
	}

	value, diag := attribute.Expr.Value(ctx)
	if diag.HasErrors() {
		return cty.NilVal, diag.Errs()[0]
	}

	return value, nil
}

func evalAttributes(attributes hclsyntax.Attributes, ctx *hcl.EvalContext) (map[string]cty.Value, error) {
	result := map[string]cty.Value{}

	for name, attribute := range attributes {
		value, diag := attribute.Expr.Value(ctx)
		if diag.HasErrors() {
			return nil, diag.Errs()[0]
		}

		result[name] = value
	}

	return result, nil
}
