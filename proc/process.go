package proc

import "fmt"

type Process struct {
	Pid        int32   `json:"pid"`
	Ppid       int32   `json:"ppid"`
	Name       string  `json:"name"`
	State      string  `json:"state"`
	Utime      float64 `json:"utime"`
	Stime      float64 `json:"stime"`
	Vsize      uint64  `json:"vsize"`
	Rss        uint64  `json:"rss"`
	NumThreads int32   `json:"num_threads"`
}

func (p *Process) String() string {

	return fmt.Sprintf("pid=%d, ppid=%d, threads=%d, name=%s, state=%s, vsize=%v, rss=%v, utime=%.2f, stime=%.2f",
		p.Pid, p.Ppid,
		p.NumThreads, p.Name, p.State, p.Vsize, p.Rss, p.Utime, p.Stime)
}
