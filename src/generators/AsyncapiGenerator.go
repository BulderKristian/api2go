package generators

import (
	"github.com/codedevstem/api2go/src/mappers"
	"github.com/codedevstem/api2go/src/models"
)

func GenerateAsyncModelMaps(asyncApiModel models.AsyncApi) map[string]interface{} {
	nestedSchemasMap := make([]map[string]models.Schema, 0)
	modelsMap := make(map[string]interface{}, 0)
	messages := asyncApiModel.Components.Messages
	if messages != nil {
		for messageName, message := range *messages {
			payload := message.Payload
			if payload.Type == "object" {
				mappedModels := mappers.MapMessageToModel(message, messageName)
				for _, mappedModel := range mappedModels {
					modelsMap[mappedModel.ModelName] = mappedModel.Model
				}
			}
		}
	}
	schemas := asyncApiModel.Components.Schemas
	if schemas != nil {
		for schemaName, schema := range *schemas {
			if schema.Type == "object" {
				mappedModels := mappers.MapSchemaToModel(schema, schemaName)
				for _, mappedModel := range mappedModels {
					modelsMap[mappedModel.ModelName] = mappedModel.Model
				}
			} else if schema.Type == "string" && len(schema.Enums) > 0 {
				modelsMap[schemaName] = mappers.MapSchemaToEnum(schema, schemaName)
			}
		}
	}
	if len(nestedSchemasMap) != 0 {
		for _, nestedSchemaMap := range nestedSchemasMap {
			for nestedSchemaName, nestedSchema := range nestedSchemaMap {
				modelsMap[nestedSchemaName] = mappers.MapSchemaToModel(nestedSchema, nestedSchemaName)
			}
		}
	}

	return modelsMap
}
