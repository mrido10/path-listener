package service

import (
	"errors"
	"io/ioutil"

	"github.com/mrido10/path-listener/util"
)

func (p Path) validateField(pth ListPath) error {
	if pth.FuncProcessing == nil {
        return errors.New("Function ListPath.FuncProcessed can't nil")
    }
    if pth.PathOrigin == "" {
        return errors.New("Directory ListPath.Origin can't empty")
    }

    _, err := ioutil.ReadDir(pth.PathOrigin)
    if err != nil {
		return errors.New("ListPath.PathOrigin" + err.Error()) 
    }
    return nil
}

func (p Path) additionalSet(pth *ListPath) {
    pth.PathOrigin = util.AddEscapeInLastDir(pth.PathOrigin)
    if pth.PathDone == "" {
        pth.PathDone = util.AddEscapeInLastDir(pth.PathOrigin) + "done/"
	} else {
        pth.PathDone = util.AddEscapeInLastDir(pth.PathDone)
    }
}