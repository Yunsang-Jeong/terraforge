package app

import (
	"github.com/zclconf/go-cty/cty"
)

func (app *terraforge) parse(configFile string) (map[string][]generateBlock, error) {
	configBody, err := getHclBodyFromFile(configFile, app.parser)
	if err != nil {
		app.lg.Error("fail to parse config", "configFile", configFile, "err", err.Error())
		return nil, err
	}

	metafiles, _ := getAttribute(configBody.Attributes, "metafiles", app.ctx, false)

	metadata := map[string]cty.Value{}
	for _, metafile := range metafiles.AsValueSlice() {
		metafileBody, err := getHclBodyFromFile(metafile.AsString(), app.parser)
		if err != nil {
			app.lg.Error("fail to parse metafile", "metafile", metafile, "err", err.Error())
			continue
		}

		attributes, err := evalAttributes(metafileBody.Attributes, app.ctx)
		if err != nil {
			app.lg.Error("fail to convert metadata", "err", err.Error())
			return nil, err
		}

		for k, v := range attributes {
			metadata[k] = v
		}
	}
	app.ctx.Variables["meta"] = cty.ObjectVal(metadata)

	generateBlocks := map[string][]generateBlock{}
	for _, block := range configBody.Blocks {
		if block.Type != "generate" {
			continue
		}

		when, _ := getAttribute(block.Body.Attributes, "when", app.ctx, false)
		if !when.IsNull() && when.False() {
			continue
		}

		for _, blockInBlock := range block.Body.Blocks {
			if blockInBlock.Type != "config" {
				continue
			}

			l0 := block.Labels[0]

			generateBlocks[l0] = append(generateBlocks[l0], generateBlock{
				labels: block.Labels,
				config: blockInBlock,
			})
		}

	}

	return generateBlocks, nil
}
