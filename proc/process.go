package proc

import "fmt"

type Process struct {
	Pid         int32
	ppid        int32
	name        string
	state       string
	utime       float64
	stime       float64
	vsize       uint64
	rss         uint64
	num_threads int32
}

func (p *Process) String() string {

	return fmt.Sprintf("pid=%d, ppid=%d, threads=%d, name=%s, state=%s, vsize=%v, rss=%v, utime=%.2f, stime=%.2f", p.Pid, p.ppid,
		p.num_threads, p.name, p.state, p.vsize, p.rss, p.utime, p.stime)
}

func (p *Process) Ppid() int32 {
	return p.ppid
}

func (p *Process) Name() string {
	return p.name
}

func (p *Process) Cpu_times() (float64, float64) {
	return p.utime, p.stime
}

func (p *Process) Status() string {
	return p.state
}

func (p *Process) Memory_info() (uint64, uint64) {

	return p.vsize, p.rss
}
