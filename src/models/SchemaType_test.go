package models_test

import (
	"github.com/codedevstem/api2go/src/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSchemaType_correctAsyncInput_assertCorrectResponseAndNoError(t *testing.T) {
	validAsyncInput := "asyncapi"
	templatePath, schemaType, err := models.ValidateAndParseSchemaType(validAsyncInput)
	assert.NoError(t, err)
	assert.NotEmpty(t, templatePath)
	assert.NotEmpty(t, schemaType)
}

func TestParseSchemaType_invalidInput_assertErrorAndEmptyStrings(t *testing.T) {
	validAsyncInput := "invalid_string"
	templatePath, schemaType, err := models.ValidateAndParseSchemaType(validAsyncInput)
	assert.Error(t, err)
	assert.Empty(t, templatePath)
	assert.Empty(t, schemaType)
}
