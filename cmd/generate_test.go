package cmd_test

import (
	"fmt"
	"github.com/codedevstem/api2go/cmd"
	"github.com/codedevstem/api2go/test/utils"
	"github.com/matryer/is"
	"github.com/spf13/cobra"
	"testing"
)

func Test_GenerateCommand(t *testing.T) {
	_is := is.New(t)
	tt := []struct {
		args        []string
		err         error
		out         string
		packageName string
	}{
		{
			args:        []string{"--input", "../test/input/enum.yaml", "--output", "../test/actualOutput", "--schema", "openapi3"},
			err:         nil,
			packageName: "enum",
		}, {
			args:        []string{"--input", "../test/input/types.yaml", "--output", "../test/actualOutput", "--schema", "openapi3"},
			err:         nil,
			packageName: "types",
		}, {
			args:        []string{"--input", "../test/input/lists.yaml", "--output", "../test/actualOutput", "--schema", "openapi3"},
			err:         nil,
			packageName: "lists",
		}, {
			args:        []string{"--input", "../test/input/nestedObjects.yaml", "--output", "../test/actualOutput", "--schema", "openapi3"},
			err:         nil,
			packageName: "nestedObjects",
		}, {
			args:        []string{"--input", "../test/input/oneOf.yaml", "--output", "../test/actualOutput", "--schema", "openapi3"},
			err:         nil,
			packageName: "oneOf",
		},
	}

	root := &cobra.Command{Use: "generate", RunE: cmd.GenerateCmd.RunE}
	cmd.GenerateCmdFlags(root)
	for _, tc := range tt {
		out, err := utils.Execute(t, root, tc.args...)
		_is.Equal(tc.err, err)
		if tc.err == nil {
			_is.Equal(tc.out, out)
		}
		eq, err := utils.DeepCompareDirectories(fmt.Sprintf("../test/actualOutput/%s", tc.packageName), fmt.Sprintf("../test/expectedOutput/%s", tc.packageName))
		if err != nil {
			t.Error(err)
		}
		_is.True(eq)
	}
}
