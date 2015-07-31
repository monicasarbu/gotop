// +build linux

package mem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemStat(t *testing.T) {

	type io struct {
		Input  string
		Output MemStat
	}

	tests := []io{
		io{
			Input: `MemTotal:         501748 kB
MemFree:          176064 kB
Buffers:           29228 kB
Cached:           156580 kB
SwapCached:            0 kB
Active:           172836 kB
Inactive:          96916 kB
Active(anon):      84052 kB
Inactive(anon):      876 kB
Active(file):      88784 kB
Inactive(file):    96040 kB
Unevictable:           0 kB`,
			Output: MemStat{Total: 501748, Available: 361872, Used: 325684, Free: 176064, Buffers: 29228, Cached: 156580, Active: 172836, Inactive: 96916},
		},
	}

	for _, test := range tests {
		stat, err := parseProcMemInfo([]byte(test.Input))
		assert.Nil(t, err)
		assert.Equal(t, test.Output, *stat)
	}

}
