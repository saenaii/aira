package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	TimeSpan int
	Output   string
	CycloBar []string
}

const (
	defaultTimeSpan = 3600
	DefaultOutput   = "stdout"
)

var output = map[string]struct{}{
	DefaultOutput: {},
}

func LoadConfig() Config {
	return Config{
		TimeSpan: getTimeSpan(),
		Output:   getOutput(),
		CycloBar: getCycloBar(),
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

func getOutput() string {
	value := os.Getenv("OUTPUT")
	if _, ok := output[value]; ok {
		return value
	}
	return DefaultOutput
}

func getCycloBar() []string {
	envBar := os.Getenv("CYCLO_BAR")
	return strings.Split(envBar, ",")
}
