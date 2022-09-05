package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const chunkSize = 64000

func DeepCompareDirectories(dir1, dir2 string) (bool, error) {
	filesDir1, err := ioutil.ReadDir(dir1)
	if err != nil {
		return false, err
	}
	filesDir2, err := ioutil.ReadDir(dir2)
	if err != nil {
		return false, err
	}
	if len(filesDir1) != len(filesDir2) {
		return false, nil
	}
	for i := 0; i < len(filesDir1); i++ {
		fileEq, err := DeepCompareFiles(fmt.Sprintf("%s/%s", dir1, filesDir1[i].Name()), fmt.Sprintf("%s/%s", dir2, filesDir2[i].Name()))
		if err != nil {
			return false, err
		}
		if !fileEq {
			return false, nil
		}
	}
	return true, nil
}

func DeepCompareFiles(file1, file2 string) (bool, error) {

	// Open file1
	f1, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	// Open file2
	f2, err := os.Open(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true, nil
			} else if err1 == io.EOF || err2 == io.EOF {
				return false, nil
			} else {
				log.Fatal(err1, err2)
			}
		}

		if !bytes.Equal(b1, b2) {
			//TODO: Figure out way to print out diff if not equal
			return false, nil
		}
	}
}
