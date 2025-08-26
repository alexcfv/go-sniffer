package stats

import "sync"

type IPStats struct {
	Packets int
	Bytes   int
}

type Stats struct {
	TotalPackets int
	TotalBytes   int
	TCPCount     int
	UDPCount     int
	ICMPCount    int
	OtherCount   int

	mu        sync.Mutex
	IPTraffic map[string]int
	IPStats   map[string]*IPStats
}
