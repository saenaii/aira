package cyclomatic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"aira/app/command"
)

func TestParse(t *testing.T) {
	currentTime := time.Now()
	testTable := []struct {
		name   string
		input  string
		expect []Cyclo
	}{
		{
			name: "happy path",
			input: `7 server isOutOfWorkdayRange server/serve.go:67:1
					7 server isOutOfWeekendRange server/serve.go:83:1
					5 server Serve server/serve.go:24:1 `,
			expect: []Cyclo{
				{
					Dir:      "server",
					FuncName: "isOutOfWorkdayRange",
					Path:     "server/serve.go:67:1",
					Count:    "7",
					Time:     currentTime,
				},
				{
					Dir:      "server",
					FuncName: "isOutOfWeekendRange",
					Path:     "server/serve.go:83:1",
					Count:    "7",
					Time:     currentTime,
				},
				{
					Dir:      "server",
					FuncName: "Serve",
					Path:     "server/serve.go:24:1",
					Count:    "5",
					Time:     currentTime,
				},
			},
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				command.GetCurrentTime = func() time.Time {
					return time.Now()
				}
			}()

			command.GetCurrentTime = func() time.Time {
				return currentTime
			}
			assert.Equal(t, c.expect, parse(c.input))
		})
	}
}
