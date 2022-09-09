package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func MD5Checksum(path string) string {
	file, err := os.Open(path)

	if err != nil {
		log.Error(err.Error())
		return ""
	}

	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		log.Error(err.Error())
		return ""
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}