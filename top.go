package main

import (
	"fmt"
	"time"

	"github.com/monicasarbu/gotop/cpu"
	"github.com/monicasarbu/gotop/mem"
	"github.com/monicasarbu/gotop/proc"
)

func main() {

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
		fmt.Printf("Pid: %d\n", process.Pid)
		fmt.Printf("Name: %s\n", process.Name())
		fmt.Printf("Cpu: %s\n", process.Cpu_times())
		fmt.Printf("Mem: %s\n", process.Memory_info())
	}
}
