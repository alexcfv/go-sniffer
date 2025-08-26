package output

import (
	"fmt"
	"sort"
	"strings"

	"github.com/alexcfv/go-sniffer/stats"
)

type Printer struct {
	stats *stats.Stats
}

func NewPrinter(s *stats.Stats) *Printer {
	return &Printer{stats: s}
}

func (p *Printer) PrintSummary() {
	fmt.Println("\n=== Traffic Summary ===")
	fmt.Printf("Total Packets: %d\n", p.stats.TotalPackets)
	fmt.Printf("Total Bytes:   %d\n", p.stats.TotalBytes)
	fmt.Printf("TCP:           %d\n", p.stats.TCPCount)
	fmt.Printf("UDP:           %d\n", p.stats.UDPCount)
	fmt.Printf("ICMP:          %d\n", p.stats.ICMPCount)
	fmt.Printf("Other:         %d\n", p.stats.OtherCount)
}

func (p *Printer) PrintIPs() {
	fmt.Println("\n=== IP/Domain Stats ===")

	type stat struct {
		IP      string
		Packets int
		Bytes   int
	}

	var stats []stat
	for ip, count := range p.stats.IPStats {
		stats = append(stats, stat{
			IP:      ip,
			Packets: count.Packets,
			Bytes:   count.Bytes,
		})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Packets > stats[j].Packets
	})

	maxIPLen := 0
	for _, s := range stats {
		if len(s.IP) > maxIPLen {
			maxIPLen = len(s.IP)
		}
	}

	fmt.Printf("%-*s  %-10s  %-10s\n", maxIPLen, "IP/Domain", "Packets", "Bytes")
	fmt.Println(strings.Repeat("-", maxIPLen+25))

	for _, s := range stats {
		fmt.Printf("%-*s  %-10d  %-10d\n", maxIPLen, s.IP, s.Packets, s.Bytes)
	}
}

func (p *Printer) PrintSrcToDst() {
	fmt.Println("\n=== Traffic by IP pairs ===")
	traffic := p.stats.GetTraffic()

	type kv struct {
		Key   string
		Value int
	}
	var sorted []kv
	for k, v := range traffic {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	maxKeyLen := 0
	for _, entry := range sorted {
		if len(entry.Key) > maxKeyLen {
			maxKeyLen = len(entry.Key)
		}
	}

	fmt.Printf("%-*s  %s\n", maxKeyLen, "SRC -> DST", "PACKETS")
	fmt.Println(strings.Repeat("-", maxKeyLen+10))
	for _, entry := range sorted {
		fmt.Printf("%-*s  %d\n", maxKeyLen, entry.Key, entry.Value)
	}
}
