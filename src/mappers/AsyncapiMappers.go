package mappers

import (
	"fmt"
	"github.com/codedevstem/api2go/src/models"
	"github.com/codedevstem/api2go/src/utils"
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
	sort.Sort(models.ModelAttributes(properties))
	primaryModelMap["properties"] = properties
	primaryModelMap["structName"] = strings.Title(messageName)
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
	sort.Sort(models.ModelAttributes(properties))
	primaryModelMap["properties"] = properties
	primaryModelMap["structName"] = strings.Title(schemaName)
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
	} else if property.Type != "object" && len(property.OneOf) != 0 {
		oneOfModelAttributes := make([]string, 0)
		for i, oneOfSchema := range property.OneOf {
			if oneOfSchema.Ref == "" {
				attributeName := fmt.Sprintf("%s_%d", propertyName, i)
				nestedModelMaps = append(nestedModelMaps, MapSchemaToModel(oneOfSchema, attributeName)...)
				oneOfModelAttributes = append(oneOfModelAttributes, attributeName)
			} else {
				refParts := strings.Split(oneOfSchema.Ref, "/")
				attributeName := refParts[len(refParts)-1]
				oneOfModelAttributes = append(oneOfModelAttributes, attributeName)
			}

		}
		nestedModelMaps = append(nestedModelMaps, CreateOneOfModelMap(strings.Title(propertyName), oneOfModelAttributes))
		propertyMap["attributeType"] = strings.Title(propertyName)
	}
	propertyMap["attributeDescription"] = strings.Replace(property.Description, "\n", " ", -1)
	propertyMap["attributeExample"] = property.Example

	properties = append(properties, propertyMap)
	return imports, properties, nestedModelMaps
}

func CreateOneOfModelMap(modelName string, propertyNames []string) models.MappedModel {
	propertiesMaps := make([]map[string]interface{}, 0)
	for _, name := range propertyNames {
		propertiesMap := make(map[string]interface{}, 0)
		propertiesMap["titeledAttributeName"] = strings.Title(name)
		propertiesMap["attributeName"] = fmt.Sprintf("%s%s", strings.ToLower(name[:1]), name[1:])
		propertiesMap["attributeType"] = strings.Title(name)
		propertiesMap["required"] = false
		propertiesMaps = append(propertiesMaps, propertiesMap)
	}
	model := make(map[string]interface{}, 0)
	model["properties"] = propertiesMaps
	model["structName"] = strings.Title(modelName)
	model["packageName"] = "spec"
	return models.MappedModel{
		ModelName: strings.Title(modelName),
		Model:     model,
	}
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
