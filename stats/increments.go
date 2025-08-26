package stats

func (s *Stats) IncrementPackets(size int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.TotalPackets++
	s.TotalBytes += size
}

func (s *Stats) IncrementTCP() {
	s.mu.Lock()
	s.TCPCount++
	s.mu.Unlock()
}

func (s *Stats) IncrementUDP() {
	s.mu.Lock()
	s.UDPCount++
	s.mu.Unlock()
}

func (s *Stats) IncrementICMP() {
	s.mu.Lock()
	s.ICMPCount++
	s.mu.Unlock()
}

func (s *Stats) IncrementOther() {
	s.mu.Lock()
	s.OtherCount++
	s.mu.Unlock()
}
