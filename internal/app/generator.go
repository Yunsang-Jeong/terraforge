package app

// import (
// 	"os"
// 	"path"

// 	"github.com/hashicorp/hcl/v2/hclwrite"
// 	"github.com/zclconf/go-cty/cty"
// )

//	func checkCondition(condtion []bool) bool {
//		for _, c := range condtion {
//			if !c {
//				return false
//			}
//		}
//		return true
//	}

// type Provider struct {
// 	Name              string            `yaml:"name", hcl:"name"`
// 	Region            string            `yaml:"region"`
// 	DefaultTags       map[string]string `yaml:"default_tags,omitempty"`
// 	AssumeRoleArn     string            `yaml:"asssume_role_arn,omitempty"`
// 	AssumeSessionName string            `yaml:"asssume_session_name,omitempty"`
// }

// func (app *terraforge) generateAWSProvider(filename string) error {
// 	file := hclwrite.NewEmptyFile()
// 	rootBody := file.Body()

// 	for _, provider := range app.config.AWSProviders {
// 		if !checkCondition(provider.Condtion) {
// 			continue
// 		}

// 		providerBlock := rootBody.AppendNewBlock("provider", []string{"aws"})
// 		providerBody := providerBlock.Body()

// 		if provider.Name != "default" {
// 			providerBody.SetAttributeValue("alias", cty.StringVal(provider.Name))
// 		}

// 		if provider.Region != "" {
// 			providerBody.SetAttributeValue("region", cty.StringVal(provider.Region))
// 		} else {
// 			providerBody.SetAttributeValue("region", cty.StringVal("NEED_TO_REPLACE"))
// 		}

// 		if provider.AssumeRoleArn != "" {
// 			assumeRoleBlock := providerBody.AppendNewBlock("assume_role", nil)
// 			assumeRoleBody := assumeRoleBlock.Body()
// 			assumeRoleBody.SetAttributeValue("role_arn", cty.StringVal(provider.AssumeRoleArn))

// 			if provider.AssumeSessionName != "" {
// 				assumeRoleBody.SetAttributeValue("session_name", cty.StringVal(provider.AssumeSessionName))
// 			}
// 		}

// 		if provider.DefaultTags != nil {
// 			defaultTagsBlock := providerBody.AppendNewBlock("default_tags", nil)
// 			defaultTagsBody := defaultTagsBlock.Body()

// 			for tagKey, tagValue := range provider.DefaultTags {
// 				if tagValue != "" {
// 					defaultTagsBody.SetAttributeValue(tagKey, cty.StringVal(tagValue))
// 				}
// 			}
// 		}
// 	}

// 	writer, err := os.Create(path.Join("./", filename))
// 	if err != nil {
// 		return err
// 	}
// 	defer writer.Close()

// 	file.WriteTo(writer)

// 	return nil
// }

// // func (app *terraforge) generateVariable(filename string) error {
// // 	file := hclwrite.NewEmptyFile()
// // 	rootBody := file.Body()

// // 	for _, variable := range app.config.Variables {
// // 		if variable.Name == "" || variable.Type == "" {
// // 			continue
// // 		}

// // 		variableBlock := rootBody.AppendNewBlock("variable", []string{variable.Name})
// // 		variableBody := variableBlock.Body()

// // 		variableBody.SetAttributeRaw("type", hclwrite.TokensForIdentifier(variable.Type))
// // 		// typeexpr.TypeConstraintVal(variable.Type)
// // 		variableBody.SetAttributeValue("description", cty.StringVal(variable.Description))

// // 		// default := reflect.ValueOf(variable.Default)

// // 		var defaultValue cty.Value
// // 		switch variable.Default.(type) {
// // 		case string:
// // 			defaultValue = cty.StringVal(variable.Default.(string))
// // 		case int:
// // 			defaultValue = cty.NumberIntVal(variable.Default.(int64))
// // 		}
// // 		variableBody.SetAttributeValue("default", defaultValue)
// // 	}

// // 	writer, err := os.Create(path.Join("./", filename))
// // 	if err != nil {
// // 		return err
// // 	}
// // 	defer writer.Close()

// // 	file.WriteTo(writer)

// // 	return nil
// // }
