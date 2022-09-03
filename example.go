package main

import (
	"io/fs"
	"io/ioutil"
	"time"

	pathListener "github.com/mrido10/path-listener/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	// code example
	var listener = pathListener.Path {
		List: []pathListener.ListPath {
			{FuncProcessing: readFile, PathOrigin: "C:/testing/folder1", PathDone: "C:/testing/folder1/done", AutoMoveToDone: true},
			{FuncProcessing: readFile, PathOrigin: "C:/testing/folder2", AutoMoveToDone: true},
		},
		TimeWait: 5 * time.Second,
		AutoDeleteFilesDone: true,
		CronExp: "22 23 * * *",
	}
	pathListener.Listen(listener).Loop()
}

func readFile(file fs.FileInfo, fullPath string) {
	body, _ := ioutil.ReadFile(fullPath)
    log.Info(string(body))
}