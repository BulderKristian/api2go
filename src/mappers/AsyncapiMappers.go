package mappers

import (
	"fmt"
	"github.com/codedevstem/api2go/src/common"
	"github.com/codedevstem/api2go/src/models"
	"sort"
	"strings"
)

func MapMessageToModel(message models.Message, messageName string) []models.MappedModel {
	properties := make([]map[string]interface{}, 0)
	imports := make([]map[string]string, 0)
	modelMaps := make([]models.MappedModel, 0)
	primaryModelMap := make(map[string]interface{}, 0)
	payload := message.Payload
	for propertyName, property := range payload.Properties {
		mappedImports, mappedProperties, nestedModelMaps := MapPropertiesToModel(propertyName, property, imports, properties, payload.Required)
		imports = mappedImports
		properties = mappedProperties
		modelMaps = append(modelMaps, nestedModelMaps...)
	}
	if len(imports) > 0 {
		primaryModelMap["hasImports"] = true
		primaryModelMap["imports"] = imports
	}
	sort.Sort(common.ModelProperties(properties))
	primaryModelMap["properties"] = properties
	primaryModelMap["structName"] = common.Title(messageName)
	primaryModelMap["packageName"] = "spec"
	modelMaps = append(modelMaps, models.MappedModel{
		ModelName: messageName,
		Model:     primaryModelMap,
	})
	return modelMaps
}

func MapSchemaToModel(schema models.Schema, schemaName string) []models.MappedModel {
	modelMaps := make([]models.MappedModel, 0)
	imports := make([]map[string]string, 0)
	properties := make([]map[string]interface{}, 0)
	primaryModelMap := make(map[string]interface{}, 0)
	for propertyName, property := range schema.Properties {
		mappedImports, mappedProperties, nestedModelMaps := MapPropertiesToModel(propertyName, property, imports, properties, schema.Required)
		imports = mappedImports
		properties = mappedProperties
		modelMaps = append(modelMaps, nestedModelMaps...)
	}
	if len(imports) > 0 {
		primaryModelMap["hasImports"] = true
		primaryModelMap["imports"] = imports
	}
	sort.Sort(common.ModelProperties(properties))
	primaryModelMap["properties"] = properties
	primaryModelMap["structName"] = common.Title(schemaName)
	primaryModelMap["packageName"] = "spec"
	modelMaps = append(modelMaps, models.MappedModel{
		ModelName: schemaName,
		Model:     primaryModelMap,
	})
	return modelMaps
}

func MapPropertiesToModel(propertyName string,
	property models.Property,
	imports []map[string]string,
	properties []map[string]interface{},
	requiredProperties []string) (
	[]map[string]string, []map[string]interface{}, []models.MappedModel) {
	nestedModelMaps := make([]models.MappedModel, 0)
	propertyMap := make(map[string]interface{}, 0)
	isRequired := false
	for _, requiredProperty := range requiredProperties {
		if requiredProperty == propertyName {
			isRequired = true
			break
		}
	}
	propertyMap["required"] = isRequired
	propertyMap["titledAttributeName"] = common.Title(propertyName)
	propertyMap["attributeName"] = propertyName
	if property.Type == "" && property.Ref != "" {
		refParts := strings.Split(property.Ref, "/")
		propertyMap["attributeType"] = refParts[len(refParts)-1]
	}
	if property.Type != "object" && property.Type != "" {
		attributeType := common.ParseAttributeType(property.Type, property.Format, property.Items)
		switch attributeType {
		case common.DateTime:
			{
				imports = common.AddImportIfNotExisting(imports, "time")
			}
		}
		propertyMap["attributeType"] = attributeType
	} else if property.Type != "object" && len(property.OneOf) != 0 {
		oneOfModelProperties := make([]string, 0)
		oneOfModelProperties, nestedModelMaps = MapOneOfProperty(propertyName, property, nestedModelMaps)
		nestedModelMaps = append(nestedModelMaps, CreateOneOfModelMap(common.Title(propertyName), oneOfModelProperties))
		propertyMap["attributeType"] = common.Title(propertyName)
	} else if property.Type == "object" {
		nestedModelMaps = append(nestedModelMaps, MapSchemaToModel(MapObjectPropertyToSchema(property), propertyName)...)
		propertyMap["attributeType"] = common.Title(propertyName)
	}
	propertyMap["attributeDescription"] = strings.Replace(property.Description, "\n", " ", -1)
	propertyMap["attributeExample"] = property.Example

	properties = append(properties, propertyMap)
	return imports, properties, nestedModelMaps
}

func MapObjectPropertyToSchema(property models.Property) models.Schema {
	return models.Schema{
		Type:       "object",
		Example:    fmt.Sprintf("%s", property.Example),
		Pattern:    fmt.Sprintf("%s", property.Pattern),
		Properties: property.Properties,
	}
}

func MapOneOfProperty(propertyName string, property models.Property, nestedModelMaps []models.MappedModel) ([]string, []models.MappedModel) {
	oneOfModelProperties := make([]string, 0)
	for i, oneOfSchema := range property.OneOf {
		if oneOfSchema.Ref == "" {
			attributeName := fmt.Sprintf("%s_%d", propertyName, i)
			nestedModelMaps = append(nestedModelMaps, MapSchemaToModel(oneOfSchema, attributeName)...)
			oneOfModelProperties = append(oneOfModelProperties, attributeName)
		} else {
			refParts := strings.Split(oneOfSchema.Ref, "/")
			attributeName := refParts[len(refParts)-1]
			oneOfModelProperties = append(oneOfModelProperties, attributeName)
		}
	}
	return oneOfModelProperties, nestedModelMaps
}

func CreateOneOfModelMap(modelName string, propertyNames []string) models.MappedModel {
	propertiesMaps := make([]map[string]interface{}, 0)
	for _, name := range propertyNames {
		propertiesMap := make(map[string]interface{}, 0)
		propertiesMap["titledAttributeName"] = common.Title(name)
		propertiesMap["attributeName"] = fmt.Sprintf("%s%s", strings.ToLower(name[:1]), name[1:])
		propertiesMap["attributeType"] = common.Title(name)
		propertiesMap["required"] = false
		propertiesMaps = append(propertiesMaps, propertiesMap)
	}
	model := make(map[string]interface{}, 0)
	model["properties"] = propertiesMaps
	model["structName"] = common.Title(modelName)
	model["packageName"] = "spec"
	return models.MappedModel{
		ModelName: common.Title(modelName),
		Model:     model,
	}
}

func MapSchemaToEnum(schema models.Schema, schemaName string) map[string]interface{} {
	modelMap := make(map[string]interface{}, 0)
	enums := make([]map[string]string, 0)
	for _, enum := range schema.Enums {
		enumConst := make(map[string]string, 0)
		enumConst["enumConstName"] = schemaName + common.Title(enum)
		enumConst["enumStringValue"] = enum
		enums = append(enums, enumConst)
	}
	modelMap["enumType"] = schema.Type
	modelMap["isEnum"] = true
	sort.Sort(common.ModelEnums(enums))
	modelMap["enums"] = enums
	modelMap["structName"] = common.Title(schemaName)
	modelMap["packageName"] = "spec"
	return modelMap
}
