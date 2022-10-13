package common_test

import (
	"github.com/codedevstem/api2go/src/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTitle(t *testing.T) {
	tt := []struct {
		input    string
		expected string
	}{
		{input: "hello", expected: "Hello"},
		{input: "helloworld", expected: "Helloworld"},
		{input: "helloWorld", expected: "HelloWorld"},
		{input: "HelloWorld", expected: "HelloWorld"},
		{input: "hello world", expected: "Hello World"},
	}
	for _, tc := range tt {
		result := common.Title(tc.input)
		assert.Equal(t, tc.expected, result)
	}
}
