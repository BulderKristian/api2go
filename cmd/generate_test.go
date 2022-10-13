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
			args:        []string{"--input", "../test/input/generic.yaml", "--output", "../test/actualOutput", "--schema", "openapi3"},
			err:         nil,
			packageName: "generic",
		}, {
			args:        []string{"--input", "../test/input/simple_attributes.yaml", "--output", "../test/actualOutput", "--schema", "openapi3"},
			err:         nil,
			packageName: "generic",
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
