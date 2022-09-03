package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func MakeDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModeDir); err != nil {
			log.Error(err.Error())
			return err
		}
	}
	return nil
}

func AddEscapeInLastDir(dir string) string {
	if string(dir[len(dir) - 1]) != "/" {
        return dir + "/"
    }
	return dir
}