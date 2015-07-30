// +build darwin

package proc

import (
	"errors"
)

func Pids() []int32 {

	return nil
}

func GetProcess(pid int32) (*Process, error) {
	return nil, errors.New("not implemented")
}
