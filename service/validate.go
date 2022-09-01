package service

import (
	"log"
    "io/fs"
    "io/ioutil"

    "github.com/mrido10/path-listener/util"
)

func (p Path) validateField(pth ListPath) []fs.FileInfo{
	if pth.FuncProcessing == nil {
        log.Fatal("- Function ListPath.FuncProcessed can't nil")
    }
    if pth.PathOrigin == "" {
        log.Fatal("- Directory ListPath.Origin can't empty")
    }

    files, err := ioutil.ReadDir(pth.PathOrigin)
    if err != nil {
		log.Fatal("ListPath.PathOrigin" + err.Error()) 
    }

    return files
}

func (p Path) additionalSet(pth *ListPath) {
    pth.PathOrigin = util.AddEscapeInLastDir(pth.PathOrigin)
    if pth.PathDone == "" {
        pth.PathDone = util.AddEscapeInLastDir(pth.PathOrigin) + "done/"
	} else {
        pth.PathDone = util.AddEscapeInLastDir(pth.PathDone)
    }
}