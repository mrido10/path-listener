package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func MoveFile(origin, dest, fileName string) error {
	if err := MakeDirIfNotExist(dest); err != nil {
		return err
	}
	err := os.Rename(origin + fileName, dest + fileName)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}