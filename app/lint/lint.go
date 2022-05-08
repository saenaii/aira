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
	fmt.Println(data)
	v, err := fastjson.Parse(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	issues := v.Get("Issues").GetArray()
	fmt.Println(issues)
	for _, issue := range issues {
		linter := issue.Get("FromLinter").String()
		text := issue.Get("Text").String()
		fmt.Println(linter, text)
	}
}
