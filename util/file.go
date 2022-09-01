package util

import (
	"fmt"
	"os"
)

func MoveFile(origin, dest, fileName string) error {
	if err := MakeDirIfNotExist(dest); err != nil {
		return err
	}
	err := os.Rename(origin + fileName, dest + fileName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}