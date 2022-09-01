package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"time"

	"github.com/mrido10/path-listener/service"
)

func main() {
	// code example
	var listPath = []service.ListPath {
		{FuncProcessing: readFile1, PathOrigin: "C:/testing/folder1", PathDone: "C:/testing/folder1/done", AutoMoveToDone: true},
		{FuncProcessing: readFile2, PathOrigin: "C:/testing/folder2", AutoMoveToDone: true},
	} 
	service.Listen(listPath, 5 * time.Second).Loop()
}

func readFile1(file fs.FileInfo, fullPath string) {
	body, _ := ioutil.ReadFile(fullPath)
    fmt.Println(string(body))
}

func readFile2(file fs.FileInfo, fullPath string) {
	body, _ := ioutil.ReadFile(fullPath)
    fmt.Println(string(body))
}