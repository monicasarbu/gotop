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

func (p *Process) parseProcState(s string) string {

	if len(s) == 0 {
		return "<unknown>"
	}
	switch s[0] {
	case 'S':
		return "sleeping"
	case 'R':
		return "running"
	case 'D':
		return "disk sleep"
	case 'T':
		return "stopped"
	case 't': // Linux 2.6.33 onward
		return "tracing stop"
	case 'Z':
		return "zombie"
	case 'X': // from Linux 2.6.0 onward
		return "dead"
	case 'x': // Linux 2.6.33 to 3.13 only
		return "dead"
	}
	return "<unknown>"
}

func (p *Process) parseProcStat(b []byte) error {

	var all_errors []string

	fields := strings.Fields(string(b))

	for i, field := range fields {

		switch i {
		case 0:
			// skip pid
		case 1:
			// process name
			p.name = strings.Trim(field, "()")
		case 2:
			// process state
			p.state = p.parseProcState(field)
		case 3:
			// ppid
			ppid, err := strconv.Atoi(field)
			if err == nil {
				p.ppid = int32(ppid)
			} else {
				all_errors = append(all_errors, err.Error())
			}
		case 13:
			// utime: Amount of time that this process has been scheduled
			// in user mode, measured in clock ticks (divide by
			// sysconf(_SC_CLK_TCK)).
			utime, err := strconv.ParseFloat(field, 64)
			if err == nil {
				p.utime = utime / cpu.GetClockTicks()
			} else {
				all_errors = append(all_errors, err.Error())
			}
		case 14:
			//stime: Amount of time that this process has been scheduled
			// in kernel mode, measured in clock ticks (divide by
			// sysconf(_SC_CLK_TCK)).
			stime, err := strconv.ParseFloat(field, 64)
			if err == nil {
				p.stime = stime / cpu.GetClockTicks()
			} else {
				all_errors = append(all_errors, err.Error())
			}
		}
	}

	if len(all_errors) > 0 {
		return fmt.Errorf(strings.Join(all_errors, "; "))
	}
	return nil
}

func (p *Process) LoadProcStat() error {

	fmt.Printf("clock_ticks = %d", cpu.GetClockTicks())

	fname := fmt.Sprintf("/proc/%d/stat", p.Pid)

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	return p.parseProcStat(b)
}

func (p *Process) Ppid() int32 {
	return p.ppid
}

func (p *Process) Name() string {
	return p.name
}

func (p *Process) Cpu_times() (float64, float64) {
	return p.utime, p.stime
}

func (p *Process) Status() string {
	return p.state
}

func (p *Process) Memory_info() (uint64, uint64) {

	pagesize := GetPageSize()

	fname := fmt.Sprintf("/proc/%d/statm", p.Pid)

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return 0, 0
	}
	fields := strings.Fields(string(b))
	if len(fields) > 2 {
		vms, err := strconv.Atoi(fields[0])
		if err != nil {
			return 0, 0
		}
		p.vms = uint64(vms) / pagesize

		rss, err := strconv.Atoi(fields[1])
		if err != nil {
			return 0, 0
		}
		p.rss = uint64(rss) / pagesize

		return p.vms, p.rss
	}

	return 0, 0
}

func (p *Process) Cmdline() string {

	fname := fmt.Sprintf("/proc/%d/cmdline", p.Pid)

	b, err := ioutil.ReadFile(fname)
	if err == nil {
		return string(b)
	}
	return ""
}

func (p *Process) String() string {

	return fmt.Sprintf("pid=%d, ppid=%d, name=%s, state=%s, vms=%v, rss=%v, utime=%.2f, stime=%.2f", p.Pid, p.ppid, p.name, p.state, p.vms, p.rss, p.utime, p.stime)
}

func GetPageSize() uint64 {
	var pagesize C.long
	pagesize = C.sysconf(C._SC_PAGE_SIZE)
	return uint64(pagesize)
}
