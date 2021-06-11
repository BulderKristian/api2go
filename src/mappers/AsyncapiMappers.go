package mappers

import (
	"github.com/codedevstem/api2go/src/models"
	"github.com/codedevstem/api2go/src/utils"
	"sort"
	"strings"
)

func MapMessageToModel(message models.Message, messageName string) map[string]interface{} {
	properties := make([]map[string]interface{}, 0)
	imports := make([]map[string]string, 0)
	modelMap := make(map[string]interface{}, 0)
	payload := message.Payload
	for propertyName, property := range payload.Properties {
		imports, properties = MapPropertiesToModel(propertyName, property, imports, properties, payload.Required)
	}
	if len(imports) > 0 {
		modelMap["hasImports"] = true
		modelMap["imports"] = imports
	}
	sort.Sort(models.ModelAttributes(properties))
	modelMap["properties"] = properties
	modelMap["structName"] = strings.Title(messageName)
	modelMap["packageName"] = "spec"
	return modelMap
}

func MapSchemaToModel(schema models.Schema, schemaName string) map[string]interface{} {
	modelMap := make(map[string]interface{}, 0)
	imports := make([]map[string]string, 0)
	properties := make([]map[string]interface{}, 0)
	for propertyName, property := range schema.Properties {
		imports, properties = MapPropertiesToModel(propertyName, property, imports, properties, schema.Required)
	}
	if len(imports) > 0 {
		modelMap["hasImports"] = true
		modelMap["imports"] = imports
	}
	sort.Sort(models.ModelAttributes(properties))
	modelMap["properties"] = properties
	modelMap["structName"] = strings.Title(schemaName)
	modelMap["packageName"] = "spec"
	return modelMap
}

func MapPropertiesToModel(propertyName string, property models.Property, imports []map[string]string, properties []map[string]interface{}, requiredProperties []string) ([]map[string]string, []map[string]interface{}) {
	propertyMap := make(map[string]interface{}, 0)
	isRequired := false
	for _, requiredProperty := range requiredProperties {
		if requiredProperty == propertyName {
			isRequired = true
			break
		}
	}
	propertyMap["required"] = isRequired
	propertyMap["titeledAttributeName"] = strings.Title(propertyName)
	propertyMap["attributeName"] = propertyName
	if property.Type == "" && property.Ref != "" {
		refParts := strings.Split(property.Ref, "/")
		propertyMap["attributeType"] = refParts[len(refParts)-1]
	}
	if property.Type != "object" && property.Type != "" {
		attributeType := utils.ParseAttributeType(property.Type, property.Format, property.Items)
		switch attributeType {
		case models.DateTime:
			{
				imports = utils.AddImportIfNotExisting(imports, "time")
			}
		}
		propertyMap["attributeType"] = attributeType
	}
	propertyMap["attributeDescription"] = strings.Replace(property.Description, "\n", " ", -1)
	propertyMap["attributeExample"] = property.Example

	properties = append(properties, propertyMap)
	return imports, properties
}

func MapSchemaToEnum(schema models.Schema, schemaName string) map[string]interface{} {
	modelMap := make(map[string]interface{}, 0)
	enums := make([]map[string]string, 0)
	for _, enum := range schema.Enums {
		enumConst := make(map[string]string, 0)
		enumConst["enumConstName"] = schemaName + strings.Title(enum)
		enumConst["enumStringValue"] = enum
		enums = append(enums, enumConst)
	}
	modelMap["enumType"] = schema.Type
	modelMap["isEnum"] = true
	sort.Sort(models.ModelEnums(enums))
	modelMap["enums"] = enums
	modelMap["structName"] = strings.Title(schemaName)
	modelMap["packageName"] = "spec"
	return modelMap
}
