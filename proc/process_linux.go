// +build linux

package proc

import (
	"errors"
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

var page_size uint64

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

	if len(fields) != 52 {
		return errors.New("Wrong size of the input")
	}

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
		case 19:
			// Number of threads in this process (since Linux 2.6).
			// Before kernel 2.6, this field was hard coded to 0 as
			// a placeholder for an earlier removed field.
			n, err := strconv.Atoi(field)
			if err != nil {
				all_errors = append(all_errors, err.Error())
			} else {
				p.num_threads = int32(n)
			}
		case 22:
			//Virtual memory size in bytes.
			vsize, err := strconv.ParseUint(field, 10, 64)
			if err != nil {
				all_errors = append(all_errors, err.Error())
			} else {
				p.vsize = vsize / 1024 // in kB
			}

			// Resident Set Size: number of pages the process has
			// in real memory.
			rss, err := strconv.ParseUint(field, 10, 64)
			if err != nil {
				all_errors = append(all_errors, err.Error())
			} else {
				p.rss = rss * page_size / 1024 // in kB
			}

		}
	}

	if len(all_errors) > 0 {
		return fmt.Errorf(strings.Join(all_errors, "; "))
	}
	return nil
}

func (p *Process) LoadProcStat() error {

	fname := fmt.Sprintf("/proc/%d/stat", p.Pid)

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	return p.parseProcStat(b)
}

func GetProcess(pid int32) (*Process, error) {

	process := &Process{Pid: pid}
	err := process.LoadProcStat()
	return process, err
}

func (p *Process) parseProcStatM(b []byte) error {

	var all_errors []string

	fields := strings.Fields(string(b))

	for i, field := range fields {

		value, err := strconv.ParseUint(field, 10, 64)
		if err != nil {
			all_errors = append(all_errors, err.Error())
		} else {
			value = value * page_size / 1024 // in kB

			switch i {
			case 0:
				//  total program size
				p.vsize = value
			case 1:
				//resident set size
				p.rss = value
			}
		}
	}

	if len(all_errors) > 0 {
		return fmt.Errorf(strings.Join(all_errors, "; "))
	}
	return nil
}

func (p *Process) LoadProcStatM() error {

	fname := fmt.Sprintf("/proc/%d/statm", p.Pid)

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	return p.parseProcStatM(b)
}

func (p *Process) Cmdline() string {

	fname := fmt.Sprintf("/proc/%d/cmdline", p.Pid)

	b, err := ioutil.ReadFile(fname)
	if err == nil {
		return string(b)
	}
	return ""
}

func InitPageSize() {
	var sc_page_size C.long
	sc_page_size = C.sysconf(C._SC_PAGE_SIZE)
	page_size = uint64(sc_page_size)
}

func SetPageSize(size uint64) {
	// used by the unit tests
	page_size = size
}

func init() {
	InitPageSize()
}
