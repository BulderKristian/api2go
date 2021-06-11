package generators

import (
	"github.com/codedevstem/api2go/src/mappers"
	"github.com/codedevstem/api2go/src/models"
)

func GenerateAsyncModelMaps(asyncApiModel models.AsyncApi) map[string]interface{} {
	modelsMap := make(map[string]interface{}, 0)
	messages := asyncApiModel.Components.Messages
	if messages != nil {
		for messageName, message := range *messages {
			payload := message.Payload
			if payload.Type == "object" {
				modelsMap[messageName] = mappers.MapMessageToModel(message, messageName)
			}
		}
	}
	schemas := asyncApiModel.Components.Schemas
	if schemas != nil {
		for schemaName, schema := range *schemas {
			if schema.Type == "object" {
				modelsMap[schemaName] = mappers.MapSchemaToModel(schema, schemaName)
			} else if schema.Type == "string" && len(schema.Enums) > 0 {
				modelsMap[schemaName] = mappers.MapSchemaToEnum(schema, schemaName)
			}
		}
	}

	return modelsMap
}
