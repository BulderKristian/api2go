package models

import (
	"fmt"
)

type SchemaType string

const (
	AsyncApiType SchemaType = "asyncapi"
	OpenApi3Type SchemaType = "openapi3"
)

var (
	golangTemplatePath = "./templates/template.mustache"
)

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
