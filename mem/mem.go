package mem

import "fmt"

type MemStat struct {
	Total     uint64 /* in kB*/
	Available uint64
	Used      uint64
	Used_p    float64
	Free      uint64
	Buffers   uint64
	Cached    uint64
	Active    uint64
	Inactive  uint64
}

func (m *MemStat) String() string {

	return fmt.Sprintf("%d total, %d used, %d free, %d buffers", m.Total, m.Used, m.Free, m.Buffers)
}
