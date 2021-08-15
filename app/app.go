package app

import (
	"log"
	"os"
	"strconv"
	"time"

	"aira-go/app/cyclomatic"
)

const (
	defaultTimeSpan = 3600
)

func Run() {
	span := getTimeSpan()
	for range time.Tick(time.Second * time.Duration(span)) {
		cyclomatic.Analyze()
	}
}

func getTimeSpan() int {
	timeSpan, err := strconv.Atoi(os.Getenv("TIME_SPAN"))
	if err != nil {
		log.Println("get timeSpan failed")
		return defaultTimeSpan
	}
	return timeSpan
}
