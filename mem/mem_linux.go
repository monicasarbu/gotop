// +build linux

package mem

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/monicasarbu/gotop/common"
)

func parseProcMemInfo(b []byte) (*MemStat, error) {

	stat := MemStat{}

	lines := strings.Split(string(b), "\n")

	for _, line := range lines {
		if len(line) == 0 {
			// ignore empty lines
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			return nil, errors.New(fmt.Sprintf("Invalid input line %s: %q", line, fields))
		}

		v, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			return nil, err
		}

		switch fields[0] {

		case "MemTotal:":
			stat.Total = v
		case "MemFree:":
			stat.Free = v
		case "Buffers:":
			stat.Buffers = v
		case "Cached:":
			stat.Cached = v
		case "Active:":
			stat.Active = v
		case "Inactive:":
			stat.Inactive = v
		}
	}

	stat.Available = stat.Free + stat.Buffers + stat.Cached
	stat.Used = stat.Total - stat.Free
	if stat.Total > 0 {
		stat.Used_p = common.Round(float64(stat.Total-stat.Available) / float64(stat.Total))
	}

	return &stat, nil
}

func Virtual_memory() (*MemStat, error) {

	b, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	return parseProcMemInfo(b)
}
