package stats

import (
	"fmt"
	"sync"
)

type Stats struct {
	TotalPackets int
	TotalBytes   int
	TCPCount     int
	UDPCount     int
	ICMPCount    int
	OtherCount   int

	mu        sync.Mutex
	IPTraffic map[string]int
}

func NewStats() *Stats {
	return &Stats{
		IPTraffic: make(map[string]int),
	}
}

func (s *Stats) IncrementPackets(size int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.TotalPackets++
	s.TotalBytes += size
}

func (s *Stats) IncrementTCP()   { s.mu.Lock(); s.TCPCount++; s.mu.Unlock() }
func (s *Stats) IncrementUDP()   { s.mu.Lock(); s.UDPCount++; s.mu.Unlock() }
func (s *Stats) IncrementICMP()  { s.mu.Lock(); s.ICMPCount++; s.mu.Unlock() }
func (s *Stats) IncrementOther() { s.mu.Lock(); s.OtherCount++; s.mu.Unlock() }

func (s *Stats) AddTraffic(src, dst string) {
	key := fmt.Sprintf("%s -> %s", src, dst)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.IPTraffic[key]++
}

func (s *Stats) GetTraffic() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()
	copy := make(map[string]int)
	for k, v := range s.IPTraffic {
		copy[k] = v
	}
	return copy
}
