// +build linux

package mem

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/monicasarbu/gotop/common"
)

func (stat *MemStat) parseProcMemInfo(b []byte) error {

	var all_errors []string

	lines := strings.Split(string(b), "\n")

	for _, line := range lines {

		fields := strings.Fields(line)
		if len(fields) != 3 {
			all_errors = append(all_errors, fmt.Sprintf("ERR: Too few elements %d on a line", len(fields)))
			continue
		}

		v, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			all_errors = append(all_errors, err.Error())
		} else {

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
	}

	stat.Available = stat.Free + stat.Buffers + stat.Cached
	stat.Used = stat.Total - stat.Free
	if stat.Total > 0 {
		stat.Used_p = common.Round(float64(stat.Total-stat.Available) / float64(stat.Total))
	}

	if len(all_errors) > 0 {
		return fmt.Errorf(strings.Join(all_errors, "; "))
	}
	return nil
}

func Virtual_memory() (*MemStat, error) {

	stat := MemStat{}

	b, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	err = stat.parseProcMemInfo(b)
	if err != nil {
		return nil, err
	}
	return &stat, nil
}
