package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	TimeSpan   int
	Output     string
	LintConfig string
	Codes      []string
	CycloBar   []string
	HTTPHost   string
}

const (
	defaultTimeSpan = 3600
)

const (
	OutputStdout = "stdout"
	OutputHTTP   = "http"
)

var output = map[string]struct{}{
	OutputStdout: {},
	OutputHTTP:   {},
}

func LoadConfig() Config {
	return Config{
		TimeSpan:   getTimeSpan(),
		LintConfig: os.Getenv("LINT_CONFIG"),
		Codes:      getCodes(),
		Output:     getOutput(),
		CycloBar:   getCycloBar(),
		HTTPHost:   os.Getenv("HTTP_HOST"),
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
	return OutputStdout
}

func getCycloBar() []string {
	envBar := os.Getenv("CYCLO_BAR")
	return strings.Split(envBar, ",")
}

func getCodes() []string {
	value := os.Getenv("CODES")
	return strings.Split(value, ",")
}
