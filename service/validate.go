package service

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"runtime"

	"github.com/mrido10/path-listener/util"
)

func (p Path) validateField(pth ListPath, arrIdx int) error {
	if pth.FuncProcessing == nil {
        return fmt.Errorf("function List[%d].ListPath.FuncProcessed can't nil", arrIdx)
    }
    if pth.PathOrigin == "" {
        return fmt.Errorf("directory List[%d].ListPath.Origin can't empty", arrIdx)
    }

    _, err := ioutil.ReadDir(pth.PathOrigin)
    if err != nil {
		return fmt.Errorf("directory List[%d].ListPath.PathOrigin %s", arrIdx, err.Error()) 
    }

    if MapPathOrigin[pth.PathOrigin] != "" {
        return fmt.Errorf("directory List[%d].ListPath.PathOrigin already used in func %s", arrIdx, MapPathOrigin[pth.PathOrigin])
    }
    MapPathOrigin[pth.PathOrigin] = fmt.Sprintf("%s(fs.FileInfo,string) List[%d].ListPath.PathOrigin", 
        runtime.FuncForPC(reflect.ValueOf(pth.FuncProcessing).Pointer()).Name(), arrIdx)
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