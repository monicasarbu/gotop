// +build darwin

package proc

import "github.com/monicasarbu/gotop/cpu"

func Pids() []int32 {

	return nil
}

func (p *Process) Name() string {
	return ""
}

func (p *Process) Cpu_times() *cpu.CpuTimes {
	return nil
}

func (p *Process) Memory_info() *MemoryInfoStat {

	return nil
}
