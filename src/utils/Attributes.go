package utils

import (
	"github.com/codedevstem/api2go/src/models"
	"strings"
)

func ParseAttributeType(attributeType string, attributeFormat string, items map[string]string) string {
	switch attributeType {
	case "string":
		{
			switch attributeFormat {
			case "date-time":
				return models.DateTime
			case "uuid":
				return models.String
			}
		}
	case "array":
		{
			return models.Array + getRefAttributeType(items)
		}
	case "boolean":
		return models.Bool
	}

	return models.String

}

func getRefAttributeType(items map[string]string) string {
	if items["$ref"] != "" {
		refParts := strings.Split(items["$ref"], "/")
		return refParts[len(refParts)-1]

	} else if items["type"] != "" {
		return ParseAttributeType(items["type"], items["format"], nil)
	}
	return ""
}
