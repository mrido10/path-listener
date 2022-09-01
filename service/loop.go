package service

import (
	"io/fs"
	"runtime"
	"time"
)

type Path struct {
	List 	 []ListPath
	TimeWait time.Duration
}

type ListPath struct {
	FuncProcessing func(fs.FileInfo, string)
	PathOrigin 	   string
	PathDone 	   string
	AutoMoveToDone bool
}

var Listen = func(list []ListPath, timeWait time.Duration) *Path {
	return &Path{
		List: list,
		TimeWait: timeWait,
	}
}

func (p Path) Loop() {
	runtime.GOMAXPROCS(2)
	for {
		for _, path := range p.List {
			go p.readPath(path)
		}
		time.Sleep(p.TimeWait)
	}
}

