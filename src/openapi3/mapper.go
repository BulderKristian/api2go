package openapi3

import (
	"fmt"
	"github.com/codedevstem/api2go/src/common"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sort"
	"strings"
)

func MapSchemasIntoModelsMap(modelsMap map[string]interface{}, schemaMap map[string]Schema, specificFileToExtract *string) {
	for schemaName, schema := range schemaMap {
		if specificFileToExtract != nil && schemaName != *specificFileToExtract {
			continue
		}
		if schema.Type == "string" && len(schema.Enums) > 0 {
			MapSchemaToEnum(modelsMap, schema, schemaName)
		} else {
			MapSchemaIntoModelsMap(modelsMap, schema, schemaName)
		}
	}
}

func MapSchemaToEnum(modelsMap map[string]interface{}, schema Schema, enumName string) {
	modelMap := make(map[string]interface{}, 0)
	enums := make([]map[string]string, 0)
	for _, enum := range schema.Enums {
		enumConst := make(map[string]string, 0)
		enumConst["enumConstName"] = enumName + common.Title(enum)
		enumConst["enumStringValue"] = enum
		enums = append(enums, enumConst)
	}
	modelMap["enumType"] = schema.Type
	modelMap["isEnum"] = true
	sort.Sort(common.ModelEnums(enums))
	modelMap["enums"] = enums
	modelMap["structName"] = common.Title(enumName)
	modelMap["packageName"] = "spec"
	modelsMap[enumName] = modelMap
}

func MapSchemaIntoModelsMap(modelsMap map[string]interface{}, schema Schema, schemaName string) {
	modelMap := make(map[string]interface{}, 0)
	MapPropertiesIntoModel(modelsMap, modelMap, schema.Properties, schema.Required)
	modelMap["structName"] = common.Title(schemaName)
	modelMap["packageName"] = "spec"
	modelsMap[schemaName] = modelMap
}

// MapPropertiesIntoModel
/**
 * Once a property list including one or more property is found, we can transform it,
 * based on the information parsed from the spec, to a golang
 */
func MapPropertiesIntoModel(
	rootModelsMap map[string]interface{},
	parent map[string]interface{},
	properties map[string]Schema,
	requiredProperties []string,
) {
	mappedProperties := make([]map[string]interface{}, 0)
	imports := make([]map[string]string, 0)
	for propertyName, property := range properties {
		mappedProperty := make(map[string]interface{})
		mappedProperty["titledAttributeName"] = common.Title(propertyName)
		mappedProperty["attributeName"] = propertyName
		attributeType := ""
		if property.Type == "" && property.Ref != "" {
			refParts := strings.Split(property.Ref, "/")
			attributeType = common.Title(refParts[len(refParts)-1])
			if strings.Contains(property.Ref, ".yaml") {
				parsedFileModel := LoadExternalFileAndConvertToOpenapi3Model(property.Ref)
				MapSchemasIntoModelsMap(rootModelsMap, parsedFileModel.Components.Schemas, &refParts[len(refParts)-1])
			}
		} else if property.AllOf != nil {
			allOfSchemaName := fmt.Sprintf("allOf%s", common.Title(propertyName))
			MapSchemaIntoModelsMap(rootModelsMap, MapAllOfToNewAllOfSchema(rootModelsMap, property.AllOf), allOfSchemaName)
			attributeType = common.Title(allOfSchemaName)
		} else if property.OneOf != nil {
			oneOfSchemaName := fmt.Sprintf("oneOf%s", common.Title(propertyName))
			MapSchemaIntoModelsMap(rootModelsMap, MapOneOfToNewOneOfSchema(rootModelsMap, property.OneOf), oneOfSchemaName)
			attributeType = common.Title(oneOfSchemaName)
		} else if property.AnyOf != nil {
			anyOfPropertyName := fmt.Sprintf("anyOf%s", common.Title(propertyName))
			MapSchemaIntoModelsMap(rootModelsMap, MapAnyOfToNewAnyOfSchema(rootModelsMap, property.AnyOf), anyOfPropertyName)
			attributeType = common.Title(anyOfPropertyName)
		} else if property.Type == "object" {
			MapSchemaIntoModelsMap(rootModelsMap, property, common.Title(propertyName))
			attributeType = common.Title(propertyName)
		} else {
			attributeType = common.ParseAttributeType(property.Type, property.Format, property.Items)
		}
		mappedProperty["attributeType"] = attributeType
		for _, requiredProperty := range requiredProperties {
			if requiredProperty == propertyName {
				mappedProperty["required"] = true
			}
		}
		imports = addImportIfApplicable(imports, attributeType)
		mappedProperties = append(mappedProperties, mappedProperty)
	}
	if len(imports) > 0 {
		parent["hasImports"] = true
	}
	parent["imports"] = imports
	sort.Sort(common.ModelProperties(mappedProperties))
	parent["properties"] = mappedProperties
}

func addImportIfApplicable(imports []map[string]string, attributeType string) []map[string]string {
	switch attributeType {
	case common.DateTime, common.DateTimeArray:
		if !checkIfImportExistsInArrayMap(imports, "time") {
			imports = append(imports, map[string]string{"importValue": "time"})
		}
	}

	return imports
}
func checkIfImportExistsInArrayMap(imports []map[string]string, value string) bool {
	for _, imp := range imports {
		for _, existingValue := range imp {
			if existingValue == value {
				return true
			}
		}
	}
	return false
}

func MapAllOfToNewAllOfSchema(rootModelsMap map[string]interface{}, allOf []Schema) Schema {
	allOfSchema := Schema{}
	allOfSchema.Type = "object"
	items := make([]string, 0)
	properties := make(map[string]Schema)
	for i, schema := range allOf {
		propertyName := FindPropertyName(rootModelsMap, &schema, i)
		properties[propertyName] = schema
		items = append(items, propertyName)
	}
	allOfSchema.Required = items
	allOfSchema.Properties = properties
	return allOfSchema
}

func MapOneOfToNewOneOfSchema(rootModelsMap map[string]interface{}, oneOf []Schema) Schema {
	oneOfSchema := Schema{}
	oneOfSchema.Type = "object"
	properties := make(map[string]Schema)
	for i, schema := range oneOf {
		propertyName := FindPropertyName(rootModelsMap, &schema, i)
		properties[propertyName] = schema
	}
	oneOfSchema.Properties = properties
	return oneOfSchema
}

func MapAnyOfToNewAnyOfSchema(rootModelsMap map[string]interface{}, anyOf []Schema) Schema {
	anyOfSchema := Schema{}
	anyOfSchema.Type = "object"
	properties := make(map[string]Schema)
	for i, schema := range anyOf {
		propertyName := FindPropertyName(rootModelsMap, &schema, i)
		properties[propertyName] = schema
	}
	anyOfSchema.Properties = properties
	return anyOfSchema
}

func FindPropertyName(rootModelsMap map[string]interface{}, schema *Schema, i int) string {
	propertyName := ""
	if schema.Ref != "" {
		refParts := strings.Split(schema.Ref, "/")
		propertyName = common.Title(refParts[len(refParts)-1])
		if strings.Contains(schema.Ref, ".yaml") {
			parsedFileModel := LoadExternalFileAndConvertToOpenapi3Model(schema.Ref)
			MapSchemasIntoModelsMap(rootModelsMap, parsedFileModel.Components.Schemas, &refParts[len(refParts)-1])
		}
	} else if schema.Title != "" {
		propertyName = common.Title(schema.Title)
		MapSchemaIntoModelsMap(rootModelsMap, *schema, propertyName)
	} else {
		propertyName = common.Title(fmt.Sprintf("inlineObject%d", i))
		MapSchemaIntoModelsMap(rootModelsMap, *schema, propertyName)
		schema.Type = propertyName
	}
	return propertyName
}

func LoadExternalFileAndConvertToOpenapi3Model(externalReference string) Openapi3 {
	refFileAndObjectParts := strings.Split(externalReference, "#/")
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", common.GetInputPath(), refFileAndObjectParts[0]))
	if err != nil {
		panic(fmt.Sprintf("failed to read inputfile: %v", err))
	}
	parsedFileModel := Openapi3{}
	err = yaml.Unmarshal(content, &parsedFileModel)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshall to openapi3: %v", err))
	}
	return parsedFileModel
}
