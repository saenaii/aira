package cyclomatic

import (
	"log"
	"strings"
	"time"

	"aira/app/command"
	"aira/config"
)

type Cyclo struct {
	Dir      string    `json:"dir"`
	FuncName string    `json:"funcNam"`
	Path     string    `json:"path"`
	Count    string    `json:"count"`
	Time     time.Time `json:"time"`
}

func Analyze(conf config.Config) {
	result := make(map[string][]Cyclo)
	for _, bar := range conf.CycloBar {
		stdout := command.Exec("gocyclo", "-over", bar, ".")
		result[bar] = parse(stdout)
	}

	err := handler[conf.Output](result)
	if err != nil {
		log.Println(err)
	}
}

func parse(input string) []Cyclo {
	arr := strings.Split(input, "\n")

	list := make([]Cyclo, 0, len(arr))
	now := command.GetCurrentTime()
	for _, row := range arr {
		item := strings.Fields(strings.TrimSpace(row))
		if len(item) != 4 {
			continue
		}
		list = append(list, Cyclo{
			Dir:      item[1],
			FuncName: item[2],
			Path:     item[3],
			Count:    item[0],
			Time:     now,
		})
	}
	return list
}
