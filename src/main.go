package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-sniffer/internal/capture"
	"go-sniffer/internal/output"
	"go-sniffer/internal/parser"
	"go-sniffer/internal/stats"
)

func main() {
	iface := flag.String("i", "", "Network interface to capture packets")
	filter := flag.String("f", "", "BPF filter")
	duration := flag.Int("t", 0, "Capture duration in seconds (0=indefinite)")
	flag.Parse()

	if *iface == "" {
		log.Fatal("Please specify an interface using -i")
	}

	// Initialize dependencies
	statsCollector := stats.NewStats()
	packetParser := parser.NewParser(statsCollector)
	capturer := capture.NewCapturer(*iface, *filter, packetParser)
	tablePrinter := output.NewPrinter(statsCollector)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go capturer.Start()

	if *duration > 0 {
		go func() {
			time.Sleep(time.Duration(*duration) * time.Second)
			stop <- syscall.SIGTERM
		}()
	}

	<-stop
	fmt.Println("\nCapture stopped. Summary:")
	tablePrinter.Print()
}
