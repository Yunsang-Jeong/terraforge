package app

import (
	"github.com/Yunsang-Jeong/terraforge/internal/util"
	"gopkg.in/yaml.v3"
)

type Metafile struct {
	Metafiles []string `yaml:"metafiles,omitempty"`
}

func (app *terraforge) parseMetadataAndSave(rawConfig string) error {
	metafile := Metafile{}
	if err := yaml.Unmarshal([]byte(rawConfig), &metafile); err != nil {
		app.lg.Error("fail to parse metafiles from config", "rawConfig", rawConfig, "err", err.Error())
		return err
	}

	for _, metafile := range metafile.Metafiles {
		rawMetadata, err := util.GetSomethingInParents(".", metafile)
		if err != nil {
			app.lg.Error("fail to load metafile", "metafile", metafile, "err", err.Error())
			continue
		}

		metadata := map[string]string{}
		if err := yaml.Unmarshal([]byte(rawMetadata), metadata); err != nil {
			app.lg.Error("fail to parse metafile", "metafile", metafile, "rawMetadata", rawMetadata, "err", err.Error())
			continue
		}

		for k, v := range metadata {
			app.metadata[k] = v
		}
	}

	return nil
}
