package stats

type Stats struct {
	TotalPackets int
	TotalBytes   int
	TCPCount     int
	UDPCount     int
	ICMPCount    int
	OtherCount   int
}

func NewStats() *Stats {
	return &Stats{}
}

func (s *Stats) IncrementPackets(size int) {
	s.TotalPackets++
	s.TotalBytes += size
}

func (s *Stats) IncrementTCP()   { s.TCPCount++ }
func (s *Stats) IncrementUDP()   { s.UDPCount++ }
func (s *Stats) IncrementICMP()  { s.ICMPCount++ }
func (s *Stats) IncrementOther() { s.OtherCount++ }
