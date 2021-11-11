package common

import (
	"fmt"
	"strings"
)

var (
	InputPath     string
	InputFileName string
)

func SetInputPath(input string) {
	inputPaths := strings.Split(input, "/")
	InputPath = fmt.Sprintf("%s/", strings.Join(inputPaths[:len(inputPaths)-1], "/"))
	fmt.Printf("%s\n", InputPath)
	fileNameAndExtension := strings.Split(inputPaths[len(inputPaths)-1], ".")
	InputFileName = fileNameAndExtension[0]
	fmt.Printf("%s\n", InputFileName)

}

func GetInputPath() string {
	return InputPath
}
func GetInputFileName() string {
	return InputFileName
}
