package models

import "fmt"

type SchemaType string

const (
	AsyncApiType SchemaType = "asyncapi"
)

var (
	asyncApi = "./templates/template.asyncapi.mustache"
)

func ValidateAndParseSchemaType(schemaTypeInput string) (string, SchemaType, error) {
	schemaType := SchemaType(schemaTypeInput)
	switch schemaType {
	case AsyncApiType:
		return asyncApi, schemaType, nil
	}
	return "", "", fmt.Errorf("invalid schematype: valid schema types are [asyncapi]")
}

func SelectModelTypeFromSchemaType(schemaType SchemaType) (interface{}, error) {
	switch schemaType {
	case AsyncApiType:
		return &AsyncApi{}, nil
	}
	return nil, fmt.Errorf("invalid schematype: valid schema types are [asyncapi]")
}
