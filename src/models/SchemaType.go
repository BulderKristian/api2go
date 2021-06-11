package models

import "fmt"

type SchemaType string

const (
	AsyncApiType SchemaType = "asyncapi"
)

var (
	asyncApi = "./templates/template.asyncapi.mustache"
)

func ValidateAndParseSchemaType(schemaTypeInput string) (string, error) {
	schemaType := SchemaType(schemaTypeInput)
	switch schemaType {
	case AsyncApiType:
		return asyncApi, nil
	}
	return "", fmt.Errorf("invalid schematype: valid schema types are [asyncapi]")
}
