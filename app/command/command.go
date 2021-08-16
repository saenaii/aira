package command

import (
	"bytes"
	"os/exec"
	"time"
)

func Exec(name string, args ...string) string {
	cmd := exec.Command(name, args...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Run()

	return stdout.String()
}

var GetCurrentTime = func() time.Time {
	return time.Now()
}
