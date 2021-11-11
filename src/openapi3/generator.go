package openapi3

func GenerateOpenapi3ModelMaps(openapi3Model Openapi3) map[string]interface{} {
	modelsMap := make(map[string]interface{}, 0)
	if openapi3Model.Components.Schemas != nil {
		MapSchemasIntoModelsMap(modelsMap, openapi3Model.Components.Schemas, nil)
	}
	return modelsMap
}
