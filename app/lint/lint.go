package lint

import (
	"fmt"

	"github.com/valyala/fastjson"

	"aira/app/command"
	"aira/config"
)

func Analyze(conf config.Config) {
	stdout := command.Exec("golangci-lint", "run", "-c", conf.LintConfig, "./...")
	parse(stdout)
}

func parse(data string) {
	v, err := fastjson.Parse(data)
	if err != nil {
		return
	}

	issues := v.Get("Issues").GetArray()
	for _, issue := range issues {
		linter := issue.Get("FromLinter").String()
		text := issue.Get("Text").String()
		filename := issue.Get("Pos").Get("Filename").String()
		line := issue.Get("Pos").Get("Line").String()

		fmt.Println(linter, text, filename, line)
	}
}
