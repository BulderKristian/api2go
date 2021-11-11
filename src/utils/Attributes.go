package utils

import (
	"github.com/codedevstem/api2go/src/common"
	"strings"
)

func ParseAttributeType(attributeType string, attributeFormat string, items map[string]string) string {
	switch attributeType {
	case "string":
		{
			switch attributeFormat {
			case "date-time":
				return common.DateTime
			case "date":
				return common.DateTime
			case "uuid":
				return common.String
			case "byte":
				return common.Byte
			case "binary":
				return common.Byte
			default:
				return common.String
			}

		}
	case "integer":
		{
			switch attributeFormat {
			case "int32":
				return common.Int32
			case "int64":
				return common.Int64
			default:
				return common.Int
			}
		}
	case "number":
		{
			switch attributeFormat {
			case "float":
				return common.Float32
			case "double":
				return common.Float64
			}
		}
	case "array":
		{
			return common.Array + getRefAttributeType(items)
		}
	case "boolean":
		return common.Bool
	}

	return attributeType

}

func getRefAttributeType(items map[string]string) string {
	if items["$ref"] != "" {
		refParts := strings.Split(items["$ref"], "/")
		return strings.Title(refParts[len(refParts)-1])

	} else if items["type"] != "" {
		return ParseAttributeType(items["type"], items["format"], nil)
	}
	return ""
}
