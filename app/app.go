package app

import (
	"time"

	"aira-go/app/cyclomatic"
	"aira-go/config"
)

func Run() {
	conf := config.LoadConfig()
	for range time.Tick(time.Second * time.Duration(conf.TimeSpan)) {
		cyclomatic.Analyze(conf)
	}
}
