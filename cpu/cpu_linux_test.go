// +build linux

package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpuTimes(t *testing.T) {

	clock_ticks = GetClockTicks()

	type io struct {
		Input  string
		Output CpuTimes
	}

	tests := []io{
		io{
			Input: `cpu  2226 0 1350 1085505 107 0 138 0 0 0
cpu0 2226 0 1350 1085505 107 0 138 0 0 0
intr 170358 53 10 0 0 0 0 0 0 0 0 0 0 156 0 0 0 0 0 0 36014 4420 10047 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
ctxt 315097
btime 1438072781
processes 5919
procs_running 1
procs_blocked 0
softirq 116034 0 58275 2182 32587 9080 0 1 0 683 13226`,
			Output: CpuTimes{User: 22.26, System: 13.50, Idle: 10855.05, IOWait: 1.07, SoftIRQ: 1.38},
		},
	}

	for _, test := range tests {
		res, err := parseCpuTimes([]byte(test.Input))
		assert.Nil(t, err)
		assert.Equal(t, test.Output, *res)
	}

}

func TestCpuTimesPercent(t *testing.T) {

	clock_ticks = GetClockTicks()

	type io struct {
		Input1 string
		Input2 string
		Output CpuTimes
	}

	tests := []io{
		io{
			Input1: `cpu  4072 512 2421 4079354 275 0 324 0 0 0
cpu0 4072 512 2421 4079354 275 0 324 0 0 0
intr 486766 53 10 0 0 0 0 0 0 0 0 0 0 156 0 0 0 0 0 0 68070 16493 23632 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
ctxt 765455
btime 1438125312
processes 9668
procs_running 2
procs_blocked 0
softirq 319529 0 194245 8181 63314 21065 0 1 0 1129 31594`,
			Input2: `cpu  4072 512 2421 4079454 275 0 324 0 0 0
cpu0 4072 512 2421 4079454 275 0 324 0 0 0
intr 486790 53 10 0 0 0 0 0 0 0 0 0 0 156 0 0 0 0 0 0 68078 16493 23632 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
ctxt 765526
btime 1438125312
processes 9668
procs_running 2
procs_blocked 0
softirq 319545 0 194251 8181 63322 21065 0 1 0 1129 31596`,
			Output: CpuTimes{Idle: 100.0},
		},
	}

	for _, test := range tests {
		res1, err := parseCpuTimes([]byte(test.Input1))
		assert.Nil(t, err)

		res2, err := parseCpuTimes([]byte(test.Input2))
		assert.Nil(t, err)

		res := cpu_times_diff(res1, res2)
		assert.Equal(t, test.Output, *res)
	}

}
func TestCpuTimesError(t *testing.T) {

	clock_ticks = GetClockTicks()

	type io struct {
		Input  string
		Output CpuTimes
	}

	tests := []io{
		io{
			Input: `cpu  2226 0 1350 107 0 13
softirq 116034 0 58275 2182 32587 9080 0 1 0 683 13226`,
			Output: CpuTimes{User: 22.26, System: 13.50, Idle: 10855.05, IOWait: 1.07, SoftIRQ: 1.38},
		},
	}

	for _, test := range tests {
		_, err := parseCpuTimes([]byte(test.Input))
		assert.NotNil(t, err)
	}

}
