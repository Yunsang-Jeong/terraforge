package app

import (
	"fmt"

	"github.com/Yunsang-Jeong/terraforge/internal/util"
	"github.com/hashicorp/hcl/v2"
)

func getHclBodyFromFile(filename string, schema *hcl.BodySchema) (*hcl.BodyContent, error) {
	path, err := util.GetSomethingPathInParents(".", filename, true)
	if err != nil {
		return nil, err
	}

	file, diag := parser.ParseHCLFile(path)
	if diag.HasErrors() {
		return nil, diag.Errs()[0]
	}

	data, diag := file.Body.Content(schema)
	if diag.HasErrors() {
		return nil, diag.Errs()[0]
	}

	return data, nil
}

func getHclAttributesFromFile(filename string) (*hcl.Attributes, error) {
	path, err := util.GetSomethingPathInParents(".", filename, true)
	if err != nil {
		return nil, err
	}

	file, diag := parser.ParseHCLFile(path)
	if diag.HasErrors() {
		return nil, diag.Errs()[0]
	}

	data, _ := file.Body.JustAttributes()

	return &data, nil
}

func getMapAttribute(attrs hcl.Attributes, target string, ctx *hcl.EvalContext, errorNotExist bool) (map[string]string, error) {
	attr, ok := attrs[target]
	if !ok {
		if errorNotExist {
			return nil, fmt.Errorf("%s is not in attributes", target)
		} else {
			return nil, nil
		}
	}

	value, diag := attr.Expr.Value(ctx)
	if diag.HasErrors() {
		return nil, diag.Errs()[0]
	}

	result := map[string]string{}
	for k, v := range value.AsValueMap() {
		result[k] = v.AsString()
	}

	return result, nil
}

func getSliceAttribute(attrs hcl.Attributes, target string, ctx *hcl.EvalContext, errorNotExist bool) ([]string, error) {
	attr, ok := attrs[target]
	if !ok {
		if errorNotExist {
			return nil, fmt.Errorf("%s is not in attributes", target)
		} else {
			return nil, nil
		}
	}

	value, diag := attr.Expr.Value(ctx)
	if diag.HasErrors() {
		return nil, diag.Errs()[0]
	}

	result := []string{}
	for _, v := range value.AsValueSlice() {
		result = append(result, v.AsString())
	}

	return result, nil
}

func getStringAttribute(attrs hcl.Attributes, target string, ctx *hcl.EvalContext, errorNotExist bool) (string, error) {
	attr, ok := attrs[target]
	if !ok {
		if errorNotExist {
			return "", fmt.Errorf("%s is not in attributes", target)
		} else {
			return "", nil
		}
	}

	value, diag := attr.Expr.Value(ctx)
	if diag.HasErrors() {
		return "", diag.Errs()[0]
	}

	return value.AsString(), nil
}

func getBoolAttribute(attrs hcl.Attributes, target string, ctx *hcl.EvalContext, errorNotExist bool) (bool, error) {
	attr, ok := attrs[target]
	if !ok {
		if errorNotExist {
			return false, fmt.Errorf("%s is not in attributes", target)
		} else {
			return false, nil
		}
	}

	value, diag := attr.Expr.Value(ctx)
	if diag.HasErrors() {
		return false, diag.Errs()[0]
	}

	return value.True(), nil
}
