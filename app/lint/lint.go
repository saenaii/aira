package lint

import (
	"fmt"

	"aira/app/command"
	"aira/config"
)

func Analyze(conf config.Config) {
	stdout := command.Exec("golangci-lint", "run", "./...")
	fmt.Println(stdout)
}
