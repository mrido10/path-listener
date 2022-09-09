package util

import (
	"os"
)

func MoveFile(origin, dest, fileName string) error {
	if err := MakeDirIfNotExist(dest); err != nil {
		return err
	}
	err := os.Rename(origin + fileName, dest + fileName)
	if err != nil {
		return err
	}
	return nil
}