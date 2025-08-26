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

func NewPrinter(stats *stats.Stats) *Printer {
	return &Printer{stats: stats}
}

func (p *Printer) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "CATEGORY\tCOUNT")
	fmt.Fprintf(w, "Total Packets\t%d\n", p.stats.TotalPackets)
	fmt.Fprintf(w, "Total Bytes\t%d\n", p.stats.TotalBytes)
	fmt.Fprintf(w, "TCP\t%d\n", p.stats.TCPCount)
	fmt.Fprintf(w, "UDP\t%d\n", p.stats.UDPCount)
	fmt.Fprintf(w, "ICMP\t%d\n", p.stats.ICMPCount)
	fmt.Fprintf(w, "Other\t%d\n", p.stats.OtherCount)
	w.Flush()

	fmt.Println("\nTraffic by IP pairs:")
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

	w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "SRC -> DST\tPACKETS")
	for _, entry := range sorted {
		fmt.Fprintf(w, "%s\t%d\n", entry.Key, entry.Value)
	}
	w.Flush()
}
