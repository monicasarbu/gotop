// +build linux

package load

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseLoadAvg(b []byte) (*LoadStat, error) {

	fields := strings.Fields(string(b))
	if len(fields) < 3 {
		return nil, errors.New("Invalid input")
	}

	l1, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return nil, err
	}

	l2, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return nil, err
	}

	l3, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return nil, err
	}

	return &LoadStat{
		Load1:  l1,
		Load5:  l2,
		Load15: l3,
	}, nil

}

func Load() (*LoadStat, error) {

	b, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return nil, err
	}
	return parseLoadAvg(b)
}
