package proc

type Process struct {
	Pid   int32
	ppid  int32
	name  string
	state string
	utime float64
	stime float64
	vms   uint64
	rss   uint64
}
