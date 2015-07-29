package main

import (
	"fmt"
	"time"

	"github.com/monicasarbu/gotop/cpu"
	"github.com/monicasarbu/gotop/load"
	"github.com/monicasarbu/gotop/mem"
	"github.com/monicasarbu/gotop/proc"
)

func main() {

	l, _ := load.Load()
	fmt.Printf("Load: %.2f %.2f %.2f\n", l.Load1, l.Load5, l.Load15)

	stat2, _ := cpu.Cpu_times_percent(1 * time.Second)
	fmt.Printf("%%Cpu: %s\n", stat2.String())

	stat2, _ = cpu.Cpu_times_percent(0)
	time.Sleep(1 * time.Second)
	stat2, _ = cpu.Cpu_times_percent(0)

	fmt.Printf("Cpu: %s\n", stat2.String())

	m, _ := mem.Virtual_memory()
	fmt.Printf("kB Mem: %s\n", m.String())

	fmt.Printf("Pids: %v\n", proc.Pids())

	pids := proc.Pids()

	for _, pid := range pids {
		process := &proc.Process{Pid: pid}
		process.LoadProcStat()
		fmt.Printf("Pid: %d\n", process.Pid)
		fmt.Printf("PPid: %d\n", process.Ppid())
		fmt.Printf("Name: %s\n", process.Name())
		fmt.Printf("Status: %s\n", process.Status())

		utime, stime := process.Cpu_times()
		fmt.Printf("Cpu: %.2f utime, %.2f stime\n", utime, stime)

		vms, rss := process.Memory_info()
		fmt.Printf("Mem: %v vms, %v rss\n", vms, rss)

		cmdline := process.Cmdline()
		fmt.Printf("Cmdline: %s\n", cmdline)
	}
}
