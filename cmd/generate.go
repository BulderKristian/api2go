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
	generateCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", ".yaml spec to parse")
	generateCmd.PersistentFlags().StringVarP(&outputFolder, "output", "o", "", "output folder")
	generateCmd.PersistentFlags().StringVarP(&schemaType, "schema", "s", "", "schema to read from [asyncapi]")
	rootCmd.AddCommand(generateCmd)
}

var (
	inputFile    string
	outputFolder string
	schemaType   string
	generateCmd  = &cobra.Command{
		Use:   "generate",
		Short: "Generate golang models from .yaml specs",
		Run: func(cmd *cobra.Command, args []string) {
			runCommand()
		},
	}
)

func runCommand() {
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
			fmt.Printf("failed to unmarshall to asyncApi: %v", err)
		}
		modelsMap = generators.GenerateAsyncModelMaps(parsedFileModel)
	case models.OpenApi3Type:
		parsedFileModel := openapi3.Openapi3{}
		err = yaml.Unmarshal(content, &parsedFileModel)
		if err != nil {
			panic(fmt.Errorf("failed to unmarshall to openapi3: %v", err))
		}
		modelsMap = openapi3.GenerateOpenapi3ModelMaps(parsedFileModel)
	default:
		fmt.Printf("invalid schematype")
		return
	}

	template, err := mustache.ParseFile(templatePath)
	if err != nil {
		fmt.Printf("failed to parse templates file: %v", err)
	}
	for modelName, modelMap := range modelsMap {
		var contentBuffer bytes.Buffer
		err = template.FRender(&contentBuffer, modelMap)
		if err != nil {
			fmt.Printf("failed to render template: %v", err)
		}
		utils.WriteToFile(contentBuffer, fmt.Sprintf("%s/%s", outputFolder, common.GetInputFileName()), fmt.Sprintf("%s%s", strings.ToLower(modelName[:1]), modelName[1:]))

	}
}
