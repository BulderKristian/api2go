package utils

import (
	"bytes"
	"fmt"
	"os"
)

func WriteToFile(buf bytes.Buffer, outDir string, modelName string) {
	// Create output folder if it does not exist already
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		err := os.MkdirAll(outDir, 0755)
		if err != nil {
			panic(fmt.Errorf("failed to create output directory, %v", err))
		}
	}
	// Create output file in directory created above
	file, err := os.Create(fmt.Sprintf("%s/model_%s.go", outDir, modelName))
	if err != nil {
		panic(fmt.Errorf("failed to write file: %v", err))
	}
	err = os.Chmod(fmt.Sprintf("%s/model_%s.go", outDir, modelName), 0777)
	if err != nil {
		panic(fmt.Errorf("failed to edit file permissions: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Write contents to file
	_, err = file.Write(buf.Bytes())
	if err != nil {
		panic("iik!")
	}
}
