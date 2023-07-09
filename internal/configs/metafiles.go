package configs

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Metafile struct {
	Paths []string
}

func decodeMetafileBlock(block *hclsyntax.Block, ctx *hcl.EvalContext) (*Metafile, hcl.Diagnostics) {
	var diags hcl.Diagnostics

	attr, ok := block.Body.Attributes["path"]
	if !ok {
		return nil, diags.Append(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Invalid metafile block",
			Detail:   "metafile block must have path attirbute only",
		})
	}

	path, diags := attr.Expr.Value(ctx)
	if diags.HasErrors() {
		return nil, diags
	}

	metafile := &Metafile{
		Paths: []string{},
	}

	for _, p := range path.AsValueSlice() {
		metafile.Paths = append(metafile.Paths, p.AsString())
	}

	return metafile, nil
}
