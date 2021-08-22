package app

import (
	"time"

	"aira/app/cyclomatic"
	"aira/config"
)

func Run() {
	conf := config.LoadConfig()
	for range time.Tick(time.Second * time.Duration(conf.TimeSpan)) {
		cyclomatic.Analyze(conf)
	}
}
