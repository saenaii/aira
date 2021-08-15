package cyclomatic

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"aira-go/app/command"
)

type Cyclo struct {
	Dir      string `json:"dir"`
	FuncName string `json:"funcNam"`
	Path     string `json:"path"`
	Count    string `json:"count"`
	Time     time.Time `json:"time"`
}

var handler = map[string]func(map[string][]Cyclo){
	command.DefaultOutput: handleStdOut,
}

func Analyze() {
	bars := getCycloBar()
	result := make(map[string][]Cyclo)
	for _, bar := range bars {
		stdout := command.Exec("gocyclo", "-over", bar, ".")
		result[bar] = Parse(stdout)
	}
	handler[command.OutputType()](result)
}

func Parse(input string) []Cyclo {
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

func getCycloBar() []string {
	envBar := os.Getenv("CYCLO_BAR")
	return strings.Split(envBar, ",")
}

var handleStdOut = func(input map[string][]Cyclo) {
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}