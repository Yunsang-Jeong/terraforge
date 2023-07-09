package configs

import (
	"errors"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Metafile struct {
	Paths []string
}

func decodeMetafileBlock(block *hclsyntax.Block, ctx *hcl.EvalContext) (*Metafile, error) {
	attr, ok := block.Body.Attributes["path"]
	if !ok {
		return nil, errors.New("invalid metafile block. metafile block must have path attirbute only")
	}

	path, diags := attr.Expr.Value(ctx)
	if diags.HasErrors() {
		return nil, errors.Join(diags.Errs()...)
	}

	metafile := &Metafile{
		Paths: []string{},
	}

	for _, p := range path.AsValueSlice() {
		metafile.Paths = append(metafile.Paths, p.AsString())
	}

	return metafile, nil
}
