package output

import (
	"fmt"
	"os"
	"text/tabwriter"

	"go-sniffer/internal/stats"
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
}
