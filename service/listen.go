package service

import (
	"io/fs"
	"io/ioutil"
	"strings"

	"github.com/mrido10/path-listener/util"
	log "github.com/sirupsen/logrus"
)

var mapFileProcessed = make(map[string]string)

func (p Path) readPath(pth ListPath) {
    files, _ := ioutil.ReadDir(pth.PathOrigin)
    for _, file := range files {
        go p.processFiles(file, pth)
    }
}

func (p Path) processFiles(file fs.FileInfo, pth ListPath) {
    if file.IsDir() {
        return
    }

    fullPath := pth.PathOrigin + file.Name()
    checkSum := util.MD5Checksum(fullPath)

    if mapFileProcessed[fullPath] == checkSum {
        return
    }
    log.Info("Read file > ", fullPath)

    mapFileProcessed[fullPath] = checkSum
    pth.FuncProcessing(file, fullPath)
    delete(mapFileProcessed, fullPath)

    if pth.AutoMoveToDone {
        if err := util.MoveFile(pth.PathOrigin, pth.PathDone, file.Name()); err != nil {
            if strings.Contains(err.Error(), "The system cannot find the file specified") {
                return
            }
            log.Error(err.Error())
        }
    }
}