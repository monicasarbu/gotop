package mem

import "fmt"

type MemStat struct {
	Total     uint64  `json:"total"`
	Available uint64  `json:"available"`
	Used      uint64  `json:"used"`
	Used_p    float64 `json:"used_p"`
	Free      uint64  `json:"free"`
	Buffers   uint64  `json:"buffers"`
	Cached    uint64  `json:"cached"`
	Active    uint64  `json:"active"`
	Inactive  uint64  `json:"inactive"`
}

func (m *MemStat) String() string {

	return fmt.Sprintf("%d total, %d used(%.2f%%), %d free, %d buffers", m.Total, m.Used, m.Used_p, m.Free, m.Buffers)
}
