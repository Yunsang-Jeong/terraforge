package app

func (a *app) parseMetadata() error {
	metadata := map[string]string{
		"region": "ap-northeast-2",
		"state":  "a/b/c.tfstate",
	}

	a.metadata = metadata

	return nil
}
