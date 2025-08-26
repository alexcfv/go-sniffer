package sniffer

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexcfv/go-sniffer/capture"
	"github.com/alexcfv/go-sniffer/output"
	"github.com/alexcfv/go-sniffer/parser"
	"github.com/alexcfv/go-sniffer/resolver"
	"github.com/alexcfv/go-sniffer/stats"
)

type SnifferConfig struct {
	Interface string
	Filter    string
	Duration  int
	Resolve   bool
}

type Sniffer struct {
	cfg            SnifferConfig
	statsCollector *stats.Stats
	packetParser   *parser.Parser
	capturer       *capture.Capturer
	resolver       *resolver.Resolver
	printer        *output.Printer
}

func NewSniffer(cfg SnifferConfig) *Sniffer {
	statsCollector := stats.NewStats()
	resolver := resolver.NewResolver()
	packetParser := parser.NewParser(statsCollector, resolver, cfg.Resolve)
	capturer := capture.NewCapturer(cfg.Interface, cfg.Filter, packetParser)
	printer := output.NewPrinter(statsCollector)

	return &Sniffer{
		cfg:            cfg,
		statsCollector: statsCollector,
		packetParser:   packetParser,
		capturer:       capturer,
		resolver:       resolver,
		printer:        printer,
	}
}

func (s *Sniffer) Run() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go s.capturer.Start()

	if s.cfg.Duration > 0 {
		go func() {
			time.Sleep(time.Duration(s.cfg.Duration) * time.Second)
			stop <- syscall.SIGTERM
		}()
	}

	<-stop
	fmt.Println("\nCapture stopped. Summary:")
	s.printer.PrintSummary()
	s.printer.PrintIPs()
	s.printer.PrintSrcToDst()
}
