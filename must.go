package fileCommons

import (
	"errors"
	"os"
)

func MustFileExist(filePath string) (err error) {
	err = nil
	if fileInfo, err := os.Stat(filePath); err == nil {
		if fileInfo.IsDir() {
			err = errors.New(filePath + " is not a directory")
			return err
		}
		return err
	}
	_, err = os.Create(filePath)
	return err
}

func MustDirExist(dirPath string) (err error) {
	//todo: discuss what is better? return errM, return errN,or just like below?
	err = nil
	if fileInfo, err := os.Stat(dirPath); err == nil {
		if !fileInfo.IsDir() {
			err = errors.New(dirPath + " is not a directory")
			return err
		}
		return err
	}

	if _, err := os.Open(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}
	} else {
		return err
	}
	return err
}
