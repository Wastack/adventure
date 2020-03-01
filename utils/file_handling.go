package utils

import (
	"io/ioutil"
	"os"
)

func GetYamlFromFile(file_path string) []byte {
	file, err := os.Open(file_path)
	if err != nil {
		panic("missing test file")
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Error reading file")
	}
	return b
}
