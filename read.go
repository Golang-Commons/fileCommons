package fileCommons

import (
	"io/ioutil"
)

func ReadToString(filePath string) (content string, err error) {
	contentByte, err := ioutil.ReadFile(filePath)
	content = string(contentByte)
	return
}

func ReadToBytes(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}
