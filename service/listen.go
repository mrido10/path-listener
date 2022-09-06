package service

import (
	"io/fs"
	"io/ioutil"

	"github.com/mrido10/path-listener/util"
	log "github.com/sirupsen/logrus"
)

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
    log.Info("Read file > ", fullPath)
    pth.FuncProcessing(file, fullPath)
    if pth.AutoMoveToDone {
        util.MoveFile(pth.PathOrigin, pth.PathDone, file.Name())
    }
}