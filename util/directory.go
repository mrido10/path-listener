package util

import (
	"log"
	"os"
)

func MakeDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModeDir); err != nil {
			log.Println(err.Error())
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