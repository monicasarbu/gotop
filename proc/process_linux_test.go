// +build linux

package proc

import (
	"testing"

	"github.com/monicasarbu/gotop/cpu"
	"github.com/stretchr/testify/assert"
)

func TestProcStat(t *testing.T) {

	type io struct {
		Input  string
		Output string
	}

	tests := []io{
		io{
			Input: `
14544 (puppet) S 1 14544 1054 0 -1 1077961024 4760 0 0 0 59 59 0 0 20 0 2 0 1090 186933248 8607 18446744073709551615 1 1
0 0 0 0 0 4096 33582663 18446744073709551615 0 0 17 0 0 0 0 0 0 0 0 0 0 0 0 0 0`,
			Output: "pid=14544, ppid=1, threads=2, name=puppet, state=sleeping, vsize=182552, rss=747732992, utime=0.59, stime=0.59",
		},
		io{
			Input: `
14544 (puppet) 5 1 14544 1054 0 -1 1077961024 4760 0 0 0 59 59 0 0 20 0 2 0 1090 186933248 8607 18446744073709551615 1 1
0 0 0 0 0 4096 33582663 18446744073709551615 0 0 17 0 0 0 0 0 0 0 0 0 0 0 0 0 0`,
			Output: "pid=14544, ppid=1, threads=2, name=puppet, state=<unknown>, vsize=182552, rss=747732992, utime=0.59, stime=0.59",
		},
	}

	cpu.SetClockTicks(100)
	SetPageSize(4096)

	process := Process{Pid: 14544}

	for _, test := range tests {
		err := process.parseProcStat([]byte(test.Input))
		assert.Nil(t, err)
		assert.Equal(t, test.Output, process.String())
	}
}

func TestProcStatFailure(t *testing.T) {

	type io struct {
		Input  string
		Output string
	}

	tests := []io{
		io{
			Input: `
14544 (puppet) 5 1 14544 1054 0 -1 1077961024 4760 0 0 0 59 59 0 0 20  1090 186933248 8607 18446744073709551615 1 1
0 0 0 0 0 4096 33582663 18446744073709551615 0 0 17 0 0 0 0 0 0 0 0 0 0 0 0 0 0`,
			Output: "pid=14544, ppid=1, threads=2, name=puppet, state=sleeping, vsize=182552, rss=747732992, utime=0.59, stime=0.59",
		},
	}

	cpu.SetClockTicks(100)
	SetPageSize(4096)

	process := Process{Pid: 14544}

	for _, test := range tests {
		err := process.parseProcStat([]byte(test.Input))
		assert.NotNil(t, err)
	}
}

func TestProcStatM(t *testing.T) {

	type io struct {
		Input  string
		Output string
	}

	tests := []io{
		io{
			Input:  `5397 1042 469 239 0 576 0`,
			Output: "pid=14544, ppid=0, threads=0, name=, state=, vsize=21588, rss=4168, utime=0.00, stime=0.00",
		},
	}

	cpu.SetClockTicks(100)
	SetPageSize(4096)

	process := Process{Pid: 14544}

	for _, test := range tests {
		err := process.parseProcStatM([]byte(test.Input))
		assert.Nil(t, err)
		assert.Equal(t, test.Output, process.String())
	}
}
