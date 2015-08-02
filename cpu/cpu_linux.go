// +build linux

package cpu

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

/*
#include <unistd.h>
#include <sys/types.h>
#include <pwd.h>
#include <stdlib.h>
*/
import "C"

type ClockTicks struct {
	Value float64
}

var clock_ticks *ClockTicks

func parseCpuTimes(b []byte) (*CpuTimes, error) {

	ct := CpuTimes{}

	lines := strings.Split(string(b), "\n")

	if len(lines) == 0 {
		return nil, errors.New("Empty input")
	}

	// parse the first line that contains combined CPU usage info
	fields := strings.Fields(lines[0])
	if len(fields) != 11 {
		return nil, errors.New("Invalid /proc/stat input")
	}

	if fields[0] == "cpu" {
		for i := 1; i < len(fields); i++ {

			v, err := strconv.ParseUint(fields[i], 10, 64)
			if err != nil {
				return nil, err
			}

			switch i {
			case 1:
				ct.User = float64(v) / GetClockTicks()
			case 2:
				ct.Nice = float64(v) / GetClockTicks()
			case 3:
				ct.System = float64(v) / GetClockTicks()
			case 4:
				ct.Idle = float64(v) / GetClockTicks()
			case 5:
				ct.IOWait = float64(v) / GetClockTicks()
			case 6:
				ct.IRQ = float64(v) / GetClockTicks()
			case 7:
				ct.SoftIRQ = float64(v) / GetClockTicks()
			case 8:
				ct.Steal = float64(v) / GetClockTicks()
			case 9:
				ct.Guest = float64(v) / GetClockTicks()
			case 10:
				ct.GuestNice = float64(v) / GetClockTicks()
			}
		}

	}
	return &ct, nil
}

func Cpu_times() (*CpuTimes, error) {

	b, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return nil, err
	}

	return parseCpuTimes(b)
}

func InitClockTicks() {

	var sc_clk_tck C.long
	sc_clk_tck = C.sysconf(C._SC_CLK_TCK)
	clock_ticks = &ClockTicks{Value: float64(sc_clk_tck)}
}

func GetClockTicks() float64 {

	if clock_ticks == nil {
		InitClockTicks()
	}
	return clock_ticks.Value
}

func SetClockTicks(c float64) {
	clock_ticks = &ClockTicks{Value: c}
}
