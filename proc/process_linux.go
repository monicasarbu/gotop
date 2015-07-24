// +build linux

package proc

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/monicasarbu/gotop/cpu"
)

/*
#include <unistd.h>
#include <sys/types.h>
#include <pwd.h>
#include <stdlib.h>
*/
import "C"

func isdigit(x string) (int32, bool) {

	digit, err := strconv.Atoi(x)

	if err != nil {
		return 0, false
	}
	return int32(digit), true
}

func Pids() []int32 {

	var pids []int32

	files, _ := ioutil.ReadDir("/proc")
	for _, file := range files {
		fileName := file.Name()

		pid, ispid := isdigit(fileName)
		if ispid {
			pids = append(pids, pid)
		}
	}
	return pids
}

func (p *Process) getProcStat() []string {

	fname := fmt.Sprintf("/proc/%d/stat", p.Pid)

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil
	}

	return strings.Fields(string(b))
}

func (p *Process) Name() string {

	if len(p.name) == 0 {
		stats := p.getProcStat()
		if len(stats) > 0 {
			p.name = strings.Trim(stats[1], "()")
		}
	}
	return p.name
}

func (p *Process) Cpu_times() *cpu.CpuTimes {

	t := &cpu.CpuTimes{}

	stats := p.getProcStat()

	clock_ticks := cpu.GetClockTicks()
	us, _ := strconv.ParseFloat(stats[11], 64)
	sy, _ := strconv.ParseFloat(stats[12], 64)

	t.User = us / clock_ticks
	t.System = sy / clock_ticks

	return t
}

func (p *Process) Status() string {
	stats := p.getProcStat()

	if len(stats) > 2 {
		return stats[2]
	}
	return "<unknown>"
}

func (p *Process) Memory_info() *MemoryInfoStat {

	m := &MemoryInfoStat{}

	pagesize := GetPageSize()

	fname := fmt.Sprintf("/proc/%d/statm", p.Pid)

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil
	}
	fields := strings.Fields(string(b))
	if len(fields) > 2 {
		vms, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil
		}
		m.VMS = uint64(vms) / pagesize

		rss, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil
		}
		m.RSS = uint64(rss) / pagesize

		return m
	}

	return nil
}

func GetPageSize() uint64 {
	var pagesize C.long
	pagesize = C.sysconf(C._SC_PAGE_SIZE)
	return uint64(pagesize)
}
