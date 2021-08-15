package command

import (
	"bytes"
	"os"
	"os/exec"
	"time"
)

const (
	DefaultOutput = "stdout"
)

var output = map[string]struct{}{
	DefaultOutput: {},
}

func Exec(name string, args ...string) string {
	cmd := exec.Command(name, args...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Run()

	return stdout.String()
}

func OutputType() string {
	v := os.Getenv("OUTPUT")
	if _, ok := output[v]; ok {
		return v
	}
	return DefaultOutput
}

var GetCurrentTime = func() time.Time {
	return time.Now()
}
