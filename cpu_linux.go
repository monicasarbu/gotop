// +build linux

package main

import (
	"fmt"
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

var clock_ticks float64

func Cpu_times() (*CpuTimes, error) {

	b, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	stat := CpuTimes{}

	for _, line := range lines {

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		if fields[0] == "cpu" {
			for i := 1; i < len(fields); i++ {
				v, _ := strconv.ParseUint(fields[i], 10, 64)
				switch i {
				case 1:
					stat.User = float64(v) / clock_ticks
				case 2:
					stat.Nice = float64(v) / clock_ticks
				case 3:
					stat.System = float64(v) / clock_ticks
				case 4:
					stat.Idle = float64(v) / clock_ticks
				case 5:
					stat.IOWait = float64(v) / clock_ticks
				case 6:
					stat.IRQ = float64(v) / clock_ticks
				case 7:
					stat.SoftIRQ = float64(v) / clock_ticks
				case 8:
					stat.Steal = float64(v) / clock_ticks
				case 9:
					stat.Guest = float64(v) / clock_ticks
				case 10:
					stat.GuestNice = float64(v) / clock_ticks
				}
			}
		} else {
			continue
		}

	}

	return &stat, nil
}

func get_clock_ticks() float64 {

	var sc_clk_tck C.long
	sc_clk_tck = C.sysconf(C._SC_CLK_TCK)
	return float64(sc_clk_tck)
}

func init() {

	fmt.Printf("Init cpu_linux\n")

	clock_ticks = get_clock_ticks()
	fmt.Printf("%v\n", clock_ticks)
}
