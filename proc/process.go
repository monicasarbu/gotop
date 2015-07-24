package proc

import (
	"fmt"

	"github.com/monicasarbu/gotop/cpu"
)

type MemoryInfoStat struct {
	RSS uint64
	VMS uint64
}

type Process struct {
	Pid    int32
	name   string
	status string
	mem    *MemoryInfoStat
	cpu    *cpu.CpuTimes
}

func (m *MemoryInfoStat) String() string {
	return fmt.Sprintf("vms %d, rss %d", m.VMS, m.RSS)

}
