Go code to auto process files in directories 

## Feature
- Listen directories
- Delete files in `/done/` folder by scheduler

## Install
```bash
go get -u github.com/mrido10/path-listener
```
## Usage 
### Create function to process files
The function must have 2 params `(file fs.FileInfo, fullPath string)`

```go
func readFile(file fs.FileInfo, fullPath string) {
  body, _ := ioutil.ReadFile(fullPath)
  log.Info(string(body))
}
```
|Param|Info|
|-|-|
|`file fs.FileInfo`|`io/fs`|
|`fullPath string`|is full path want to read|

### Run code
Using lib from `github.com/mrido10/path-listener/service`

```go
import pathListener "github.com/mrido10/path-listener/service"

func main() {
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
```
|Field|Info|
|-|-|
|`FuncProcessing`|function to process files|
|`PathOrigin`|directory want to auto process files|
|`PathDone`|directory files moved after processed|
||default auto created folder `/done/` into `PathOrigin`|
|`AutoMoveToDone`|if `true` files auto move after processed|
||default `false`|
