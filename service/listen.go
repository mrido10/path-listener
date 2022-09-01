package service

import (
	"fmt"
	"io/fs"
	"log"

	"github.com/mrido10/path-listener/util"
)

func (p Path) readPath(pth ListPath) {
    files := p.validateField(pth)
    p.additionalSet(&pth)
    for _, file := range files {
        go p.processFiles(file, pth)
    }
}

func (p Path) processFiles(file fs.FileInfo, pth ListPath) {
    if file.IsDir() {
        return
    }

    fullPath := pth.PathOrigin + file.Name()
    log.Println(fmt.Sprintf(">> Read file: %s", fullPath))
    pth.FuncProcessing(file, fullPath)
    if pth.AutoMoveToDone {
        util.MoveFile(pth.PathOrigin, pth.PathDone, file.Name())
    }
}