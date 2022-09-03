package service

import (
	"io/ioutil"
	"os"

	cron "github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func (p Path) job() {
	c := cron.New()
	c.AddFunc(p.CronExp, p.jobDeleteFilesFromDir)
	go c.Start()
}

func (p Path) jobDeleteFilesFromDir() {
	for _, v := range p.List {
		files, err := ioutil.ReadDir(v.PathDone)
		if err != nil {
			log.Error(err.Error())
			return
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			fileName := v.PathDone + file.Name()
			err := os.Remove(fileName)
    		if err != nil {
				log.Error(err.Error())
				return
			}
			log.Info("Delete File > ", fileName)
		}
	}
}