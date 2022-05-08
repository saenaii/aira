package cloc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		codes     []string
		input     string
		expect    map[string]Data
		expectErr error
	}{
		{
			name:  "happy path",
			codes: []string{"Go"},
			input: `{
				"header": {
					"cloc_url": "github.com/AlDanial/cloc",
					"cloc_version": "1.92",
					"elapsed_seconds": 0.0390970706939697,
					"n_files": 14,
					"n_lines": 619,
					"files_per_second": 358.08309296582,
					"lines_per_second": 15832.3881818459
				},
				"Go": {
					"nFiles": 7,
					"blank": 60,
					"comment": 37,
					"code": 301
				},
				"XML": {
					"nFiles": 5,
					"blank": 0,
					"comment": 0,
					"code": 194
				}
			}`,
			expect: map[string]Data{
				"Go": {
					Files:   7,
					Code:    301,
					Blank:   60,
					Comment: 37,
				},
			},
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			result, err := parse(c.codes, c.input)
			assert.Equal(t, c.expect, result)
			assert.Equal(t, c.expectErr, err)
		})
	}
}
