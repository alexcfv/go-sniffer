package stats

import (
	"fmt"
)

func NewStats() *Stats {
	return &Stats{
		IPTraffic: make(map[string]int),
		IPStats:   make(map[string]*IPStats),
	}
}

func (s *Stats) AddTraffic(src, dst string) {
	key := fmt.Sprintf("%s -> %s", src, dst)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.IPTraffic[key]++
}

func (s *Stats) AddIPStats(ip string, size int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.IPStats[ip]; !ok {
		s.IPStats[ip] = &IPStats{}
	}
	s.IPStats[ip].Packets++
	s.IPStats[ip].Bytes += size
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

func (s *Stats) GetIPStats() map[string]IPStats {
	s.mu.Lock()
	defer s.mu.Unlock()
	copy := make(map[string]IPStats)
	for ip, stat := range s.IPStats {
		copy[ip] = *stat
	}
	return copy
}
