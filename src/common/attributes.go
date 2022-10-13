package common

import (
	"strings"
)

func ParseAttributeType(attributeType string, attributeFormat string, items map[string]string) string {
	switch attributeType {
	case "string":
		{
			switch attributeFormat {
			case "date-time":
				return DateTime
			case "date":
				return DateTime
			case "uuid":
				return String
			case "byte":
				return Byte
			case "binary":
				return Byte
			default:
				return String
			}

		}
	case "integer":
		{
			switch attributeFormat {
			case "int32":
				return Int32
			case "int64":
				return Int64
			default:
				return Int
			}
		}
	case "number":
		{
			switch attributeFormat {
			case "float":
				return Float32
			case "double":
				return Float64
			}
		}
	case "array":
		{
			return Array + getRefAttributeType(items)
		}
	case "boolean":
		return Bool
	}

	return attributeType

}

func getRefAttributeType(items map[string]string) string {
	if items["$ref"] != "" {
		refParts := strings.Split(items["$ref"], "/")
		return Title(refParts[len(refParts)-1])

	} else if items["type"] != "" {
		return ParseAttributeType(items["type"], items["format"], nil)
	}
	return ""
}
