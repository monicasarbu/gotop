// +build linux

package mem

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/monicasarbu/gotop/common"
)

func Virtual_memory() (*MemStat, error) {

	b, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	stat := MemStat{}

	for _, line := range lines {

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		switch fields[0] {

		case "MemTotal:":
			v, err := strconv.ParseUint(fields[1], 10, 64)
			if err == nil {
				stat.Total = v
			}
		case "MemFree:":
			v, err := strconv.ParseUint(fields[1], 10, 64)
			if err == nil {
				stat.Free = v
			}
		case "Buffers:":
			v, err := strconv.ParseUint(fields[1], 10, 64)
			if err == nil {
				stat.Buffers = v
			}
		case "Cached:":
			v, err := strconv.ParseUint(fields[1], 10, 64)
			if err == nil {
				stat.Cached = v
			}
		case "Active:":
			v, err := strconv.ParseUint(fields[1], 10, 64)
			if err == nil {
				stat.Active = v
			}

		case "Inactive:":
			v, err := strconv.ParseUint(fields[1], 10, 64)
			if err == nil {
				stat.Inactive = v
			}

		}
	}

	stat.Available = stat.Free + stat.Buffers + stat.Cached
	stat.Used = stat.Total - stat.Free
	stat.Used_p = common.Round(float64(stat.Total-stat.Available) / float64(stat.Total))

	return &stat, nil

}
