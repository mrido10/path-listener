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

var Listen = func(pth Path) *Path {
	if pth.TimeWait.Seconds() < 5 {
		pth.TimeWait = 5 * time.Second
	}
	return &pth
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

