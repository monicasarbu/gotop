package load

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	type io struct {
		Input  string
		Output LoadStat
	}

	tests := []io{
		io{
			Input:  "0.00 0.01 0.05 1/84 5293",
			Output: LoadStat{0.00, 0.01, 0.05},
		},
		io{
			Input:  "1.00 2.01 1.05 1/84 5293",
			Output: LoadStat{1.00, 2.01, 1.05},
		},
		io{
			Input:  "1.00 2.01 1.05",
			Output: LoadStat{1.00, 2.01, 1.05},
		},
	}

	for _, test := range tests {
		res, err := parseLoadAvg([]byte(test.Input))
		assert.Nil(t, err)
		assert.Equal(t, test.Output, *res)
	}
}

func TestLoadErr(t *testing.T) {

	tests := []string{
		"0.00",
		"0",
		"0 a",
		"a 0 0",
		"0.00 0.00 0.00a",
	}

	for _, test := range tests {
		_, err := parseLoadAvg([]byte(test))
		assert.NotNil(t, err)
	}
}
