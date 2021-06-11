package utils

import (
	"bytes"
	"fmt"
	"os"
)

func WriteToFile(buf bytes.Buffer, outDir string, domainName string, modelName string) {
	// Create output folder if it does not exist already
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		err := os.Mkdir(outDir, 0755)
		if err != nil {
			panic(fmt.Errorf("failed to create output directory, %v", err))
		}
	}
	// Create domain folder if it does not exist already
	if _, err := os.Stat(fmt.Sprintf("%s/%s", outDir, domainName)); os.IsNotExist(err) {
		err := os.Mkdir(fmt.Sprintf("%s/%s", outDir, domainName), 0755)
		if err != nil {
			panic(fmt.Errorf("failed to create output directory, %v", err))
		}
	}
	// Create output file in directory created above
	file, err := os.Create(fmt.Sprintf("%s/model_%s.go", outDir, modelName))
	if err != nil {
		panic(fmt.Errorf("failed to write file: %v", err))
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
