package output

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

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
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "IP/Domain\tPackets")
	for ip, count := range p.stats.IPStats {
		fmt.Fprintf(w, "%s\t%d\n", ip, count)
	}
	w.Flush()
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

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "SRC -> DST\tPACKETS")
	for _, entry := range sorted {
		fmt.Fprintf(w, "%s\t%d\n", entry.Key, entry.Value)
	}
	w.Flush()
}
