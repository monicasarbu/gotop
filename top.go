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

	m, err := mem.Virtual_memory()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("kB Mem: %s\n", m.String())
	}

	fmt.Printf("Pids: %v\n", proc.Pids())

	pids := proc.Pids()

	for _, pid := range pids {
		process, err := proc.GetProcess(pid)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			continue
		}
		fmt.Printf("%s\n", process.String())
	}
}
