package models

import (
	"fmt"
	"os"
	"path/filepath"
)

type SchemaType string

const (
	AsyncApiType  SchemaType = "asyncapi"
	OpenApi3Type  SchemaType = "openapi3"
	templateRPath            = "templates/template.mustache"
)

var (
	golangTemplatePath string
)

func init() {
	f, _ := os.Getwd()
	golangTemplatePath = fmt.Sprintf("%s/%s", filepath.Dir(f), templateRPath)
}

func ValidateAndParseSchemaType(schemaTypeInput string) (string, SchemaType, error) {
	schemaType := SchemaType(schemaTypeInput)
	switch schemaType {
	case AsyncApiType:
		return golangTemplatePath, schemaType, nil
	case OpenApi3Type:
		return golangTemplatePath, schemaType, nil
	}

	return "", "", fmt.Errorf("invalid schematype: valid schema types are [asyncapi]")
}
