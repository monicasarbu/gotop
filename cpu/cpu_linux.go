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

	stat := CpuTimes{}

	lines := strings.Split(string(b), "\n")

	if len(lines) == 0 {
		return nil, errors.New("Empty input")
	}

	// parse the first line that contains combined CPU usage info
	fields := strings.Fields(lines[0])
	if len(fields) != 11 {
		return nil, errors.New("cpu line input too short.")
	}

	if fields[0] == "cpu" {
		for i := 1; i < len(fields); i++ {
			v, _ := strconv.ParseUint(fields[i], 10, 64)
			switch i {
			case 1:
				stat.User = float64(v) / GetClockTicks()
			case 2:
				stat.Nice = float64(v) / GetClockTicks()
			case 3:
				stat.System = float64(v) / GetClockTicks()
			case 4:
				stat.Idle = float64(v) / GetClockTicks()
			case 5:
				stat.IOWait = float64(v) / GetClockTicks()
			case 6:
				stat.IRQ = float64(v) / GetClockTicks()
			case 7:
				stat.SoftIRQ = float64(v) / GetClockTicks()
			case 8:
				stat.Steal = float64(v) / GetClockTicks()
			case 9:
				stat.Guest = float64(v) / GetClockTicks()
			case 10:
				stat.GuestNice = float64(v) / GetClockTicks()
			}
		}

		return &stat, nil
	}

	return nil, errors.New("Wrong input data")
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
