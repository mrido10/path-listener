package service

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type Path struct {
	List 	 			[]ListPath
	TimeWait 			time.Duration
	AutoDeleteFilesDone bool
	CronExp 			string
}

type ListPath struct {
	FuncProcessing func(fs.FileInfo, string)
	PathOrigin 	   string
	PathDone 	   string
	AutoMoveToDone bool
}

type _listPath struct {
	FuncProcessing string 
	ListPath
}

type plainText log.TextFormatter

var Listen = func(pth Path) *Path {
	logrusInit()
	if pth.TimeWait.Seconds() < 5 {
		pth.TimeWait = 5 * time.Second
	}
	return &pth
}

func (p Path) Loop() {
	runtime.GOMAXPROCS(2)
	p.validateAndSetPathDone()
	p.printConfigurations()
	if p.AutoDeleteFilesDone {
		p.job()
	}
	for {
		for _, path := range p.List {
			go p.readPath(path)
		}
		time.Sleep(p.TimeWait)
	}
}

func (p *Path) validateAndSetPathDone() {
	for k, v := range p.List {
		err := p.validateField(v)
		if err != nil {
			log.Error(err.Error())
			os.Exit(2)
		}
		p.additionalSet(&v)
		p.List[k] = v
	}
}

func (p Path) printConfigurations() {
	var list []_listPath
	for _, v := range p.List {
		list = append(list, _listPath{
			FuncProcessing: runtime.FuncForPC(reflect.ValueOf(v.FuncProcessing).Pointer()).Name(),
			ListPath: ListPath{
				PathOrigin: v.PathOrigin,
				PathDone: v.PathDone,
				AutoMoveToDone: v.AutoMoveToDone,
			},
		})
	}
	liPath, err := json.Marshal(list)
    if err != nil {
        log.Error(err.Error())
    }
	log.Info(fmt.Sprintf(`Running path-listener using config:
		List               : %s
		TimeWait           : %v
		AutoDeleteFilesDone: %t 
		CronExp            : %s`,
		string(liPath), p.TimeWait, p.AutoDeleteFilesDone, p.CronExp))
}

// log init
func logrusInit() {
	log.SetReportCaller(true)
	log.SetFormatter(&plainText{})
}

func (f *plainText) Format(entry *log.Entry) ([]byte, error) {
	timestamp := fmt.Sprint(entry.Time.Format("2006-01-02 15:04:05.000"))
	level := entry.Level
	return []byte(fmt.Sprintf("%s [%s] [path-listener] %s\n", timestamp, strings.ToUpper(level.String()), entry.Message)), nil
}