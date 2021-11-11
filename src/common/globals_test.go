package common_test

import (
	"github.com/codedevstem/api2go/src/common"
	"testing"
)

func TestSetInputPath(t *testing.T) {
	common.SetInputPath("./inputs/test.yaml")
}
