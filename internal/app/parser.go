package app

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

var configSchema = hclSchema{
	attributes: map[string]bool{
		"metafiles": false,
	},
	blokcs: map[string][]string{
		"generate": {"type", "name"},
	},
}

var generateBlockConfigSchema = hclSchema{
	attributes: map[string]bool{
		"source": false,
		"when":   true,
	},
	blokcs: map[string][]string{
		"config": {},
	},
}

var configBlockConfigSchema = hclSchema{
	attributes: map[string]bool{
		"alias":        false,
		"region":       false,
		"assume_role":  false,
		"default_tags": false,
	},
}

type hclProviderBlock struct {
	name   string               `hcl:"name"`
	source string               `hcl:"source"`
	config map[string]cty.Value `hcl:"config,block"`
}

var parser = hclparse.NewParser()

var ctx = hcl.EvalContext{
	Variables: map[string]cty.Value{},
	Functions: map[string]function.Function{},
}

func (app *terraforge) parseConfig(configFile string) error {
	rootBody, err := getHclBodyFromFile(configFile, configSchema.convertHclBodySchema())
	if err != nil {
		app.lg.Error("fail to parse config", "configFile", configFile, "err", err.Error())
		return err
	}

	metafiles, err := getSliceAttribute(rootBody.Attributes, "metafiles", &ctx, false)
	if err != nil {
		app.lg.Error("fail to parse metafiles", "err", err.Error())
		return err
	}

	metadata := map[string]cty.Value{}
	for _, metafile := range metafiles {
		attrs, err := getHclAttributesFromFile(metafile)
		if err != nil {
			app.lg.Error("fail to parse metafile", "metafile", metafile, "err", err.Error())
			continue
		}

		for name, attr := range *attrs {
			value, diag := attr.Expr.Value(nil)
			if diag.HasErrors() {
				app.lg.Error("fail to evaluate attribute", "metafile", metafile, "attr-name", name, "err", diag.Errs()[0])
				continue
			}

			metadata[name] = value
		}
	}
	ctx.Variables["metadata"] = cty.ObjectVal(metadata)

	for _, block := range rootBody.Blocks {
		if block.Type != "generate" {
			continue
		}

		blockType := block.Labels[0]
		blockname := block.Labels[1]

		switch blockType {
		case "provider":
			providerBlock, _ := app.parseGenerateProviderBlock(blockname, block)
			fmt.Println(providerBlock)
		}
	}

	return nil
}

func (app *terraforge) parseGenerateProviderBlock(providerName string, block *hcl.Block) (*hclProviderBlock, error) {
	result := &hclProviderBlock{}

	body, diag := block.Body.Content(generateBlockConfigSchema.convertHclBodySchema())
	if diag.HasErrors() {
		app.lg.Error("fail to parse generate block for provider", "providerName", providerName, "err", diag.Errs()[0])
		return nil, diag.Errs()[0]
	}

	providerWhen, err := getBoolAttribute(body.Attributes, "when", &ctx, true)
	if err != nil {
		app.lg.Error("fail to parse condition attribute in provider", "block.Type", block.Type, "block.Labels", block.Labels, "err", err.Error())
		return nil, err
	}

	if !providerWhen {
		return nil, nil
	}

	providerSource, err := getStringAttribute(body.Attributes, "source", &ctx, false)
	if err != nil {
		app.lg.Error("fail to parse source attribute in provider", "block.Type", block.Type, "block.Labels", block.Labels, "err", err.Error())
		return nil, err
	}
	if providerSource == "" {
		providerSource = fmt.Sprintf("hashicorp/%s", providerName)
	}
	result.source = providerSource

	configBlockCount := len(body.Blocks)
	if configBlockCount > 1 {
		app.lg.Debug("too many config blcok in generate block", "block.Type", block.Type, "block.Labels", block.Labels, "err", err.Error())
		return result, fmt.Errorf("too many config blcok in generate block")
	} else if configBlockCount == 0 {
		return result, nil
	}

	configBlockBody, diag := body.Blocks[0].Body.Content(configBlockConfigSchema.convertHclBodySchema())
	if diag.HasErrors() {
		app.lg.Error("fail to parse generate block for provider", "providerName", providerName, "err", diag.Errs()[0])
		return nil, diag.Errs()[0]
	}

	result.config = map[string]cty.Value{}
	for name, attr := range configBlockBody.Attributes {
		value, diag := attr.Expr.Value(&ctx)
		if diag.HasErrors() {
			return nil, diag.Errs()[0]
		}

		result.config[name] = value
	}

	return result, nil
}
