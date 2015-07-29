package cpu

import (
	"fmt"
	"time"

	"github.com/monicasarbu/gotop/common"
)

type CpuTimes struct {
	User      float64
	Nice      float64
	System    float64
	Idle      float64
	IOWait    float64
	IRQ       float64
	SoftIRQ   float64
	Steal     float64
	Guest     float64
	GuestNice float64
}

var _last_cpu_times *CpuTimes

func (t *CpuTimes) String() string {

	return fmt.Sprintf("%.2f user, %.2f system, %.2f nice, %.2f iddle, %.2f iowait, %.2f irq, %.2f softirq, %.2f steal",
		t.User, t.System, t.Nice, t.Idle, t.IOWait, t.IRQ, t.SoftIRQ, t.Steal)
}

func (t *CpuTimes) sum() float64 {

	return t.User + t.Nice + t.System + t.Idle + t.IOWait + t.IRQ + t.SoftIRQ + t.Steal + t.Guest + t.GuestNice
}

func cpu_percentage(t1 float64, t2 float64, all float64) float64 {
	delta := t2 - t1

	perc := common.Percentage(delta, all)

	// TODO: Work around for windows
	/*
		if perc > 100.0 {
			perc = 100.0
		} else if perc < 0.0 {
			perc = 0.0
		}
	*/
	return perc
}

func cpu_times_diff(t1 *CpuTimes, t2 *CpuTimes) *CpuTimes {

	t := &CpuTimes{}

	all_delta := t2.sum() - t1.sum()

	t.User = cpu_percentage(t1.User, t2.User, all_delta)
	t.System = cpu_percentage(t1.System, t2.System, all_delta)
	t.Nice = cpu_percentage(t1.Nice, t2.Nice, all_delta)
	t.Idle = cpu_percentage(t1.Idle, t2.Idle, all_delta)
	t.IOWait = cpu_percentage(t1.IOWait, t2.IOWait, all_delta)
	t.IRQ = cpu_percentage(t1.IRQ, t2.IRQ, all_delta)
	t.SoftIRQ = cpu_percentage(t1.SoftIRQ, t2.SoftIRQ, all_delta)
	t.Steal = cpu_percentage(t1.Steal, t2.Steal, all_delta)
	t.Guest = cpu_percentage(t1.Guest, t2.Guest, all_delta)
	t.GuestNice = cpu_percentage(t1.GuestNice, t2.GuestNice, all_delta)

	return t
}

func Cpu_times_percent(interval time.Duration) (*CpuTimes, error) {

	var err error

	var t1 *CpuTimes

	blocking := false

	if interval > 0.0 {
		blocking = true
	}

	if blocking {
		t1, err = Cpu_times()
		time.Sleep(interval)
	} else {
		t1 = _last_cpu_times
	}

	_last_cpu_times, err = Cpu_times()
	if err != nil {
		return nil, err
	}

	return cpu_times_diff(t1, _last_cpu_times), nil
}

func init() {
	_last_cpu_times, _ = Cpu_times()
}
