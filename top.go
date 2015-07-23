package main

import (
	"fmt"
	"time"

	"github.com/monicasarbu/gotop/cpu"
	"github.com/monicasarbu/gotop/mem"
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

}
