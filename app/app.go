package app

import (
	"time"

	"aira/app/cloc"
	"aira/app/cyclomatic"
	"aira/app/lint"
	"aira/config"
)

func Run() {
	conf := config.LoadConfig()
	for range time.Tick(time.Second * time.Duration(conf.TimeSpan)) {
		cloc.Analyze(conf)
		cyclomatic.Analyze(conf)
		lint.Analyze(conf)
	}
}
