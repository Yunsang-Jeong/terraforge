package app

import (
	"github.com/hashicorp/hcl/v2"
)

type hclSchema struct {
	attributes map[string]bool
	blokcs     map[string][]string
}

func (s *hclSchema) convertHclBodySchema() *hcl.BodySchema {
	schema := hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{},
		Blocks:     []hcl.BlockHeaderSchema{},
	}

	for k, v := range s.attributes {
		schema.Attributes = append(schema.Attributes, hcl.AttributeSchema{
			Name:     k,
			Required: v,
		})
	}

	for k, v := range s.blokcs {
		schema.Blocks = append(schema.Blocks, hcl.BlockHeaderSchema{
			Type:       k,
			LabelNames: v,
		})
	}

	return &schema
}
