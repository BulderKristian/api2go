package cmd

import (
	"bytes"
	"fmt"
	"github.com/cbroglie/mustache"
	"github.com/codedevstem/api2go/src/common"
	"github.com/codedevstem/api2go/src/generators"
	"github.com/codedevstem/api2go/src/models"
	"github.com/codedevstem/api2go/src/openapi3"
	"github.com/codedevstem/api2go/src/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

func init() {
	GenerateCmdFlags(RootCmd)
	RootCmd.AddCommand(GenerateCmd)
}

var (
	inputFile    string
	outputFolder string
	schemaType   string
	GenerateCmd  = &cobra.Command{
		Use:   "generate",
		Short: "Generate golang models from .yaml specs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Generate()
		},
	}
)

func GenerateCmdFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", ".yaml spec to parse")
	cmd.PersistentFlags().StringVarP(&outputFolder, "output", "o", "", "output folder")
	cmd.PersistentFlags().StringVarP(&schemaType, "schema", "s", "", "schema to read from [asyncapi]")
}
func Generate() error {
	common.SetInputPath(inputFile)
	templatePath, schemaType, err := models.ValidateAndParseSchemaType(schemaType)
	if err != nil {
		fmt.Printf("failed to validate and parse schemaType %v", err)
	}

	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("failed to read inputfile: %v", err)
	}

	var modelsMap = make(map[string]interface{}, 0)
	switch schemaType {
	case models.AsyncApiType:
		parsedFileModel := models.AsyncApi{}
		err = yaml.Unmarshal(content, &parsedFileModel)
		if err != nil {
			return fmt.Errorf("failed to unmarshall to asyncApi: %v", err)
		}
		modelsMap = generators.GenerateAsyncModelMaps(parsedFileModel)
	case models.OpenApi3Type:
		parsedFileModel := openapi3.Openapi3{}
		err = yaml.Unmarshal(content, &parsedFileModel)
		if err != nil {
			return fmt.Errorf("failed to unmarshall to openapi3: %v", err)
		}
		modelsMap = openapi3.GenerateOpenapi3ModelMaps(parsedFileModel)
	default:
		return fmt.Errorf("invalid schematype")
	}

	template, err := mustache.ParseFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse templates file: %v", err)
	}
	for modelName, modelMap := range modelsMap {
		var contentBuffer bytes.Buffer
		err = template.FRender(&contentBuffer, modelMap)
		if err != nil {
			return fmt.Errorf("failed to render template: %v", err)
		}
		utils.WriteToFile(contentBuffer, fmt.Sprintf("%s/%s", outputFolder, common.GetInputFileName()), fmt.Sprintf("%s%s", strings.ToLower(modelName[:1]), modelName[1:]))
	}
_:
	return nil
}
