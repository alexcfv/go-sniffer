package main

import (
	"flag"
	"log"

	"github.com/alexcfv/go-sniffer/sniffer"
)

func main() {
	iface := flag.String("i", "", "Network interface to capture packets")
	filter := flag.String("f", "", "BPF filter")
	duration := flag.Int("t", 0, "Capture duration in seconds (0=indefinite)")
	flag.Parse()

	if *iface == "" {
		log.Fatal("Please specify an interface using -i")
	}

	cfg := sniffer.SnifferConfig{
		Interface: *iface,
		Filter:    *filter,
		Duration:  *duration,
	}

	s := sniffer.NewSniffer(cfg)
	s.Run()
}
