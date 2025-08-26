package capture

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Parser interface {
	Parse(packet gopacket.Packet)
}

type Capturer struct {
	iface  string
	filter string
	parser Parser
}

func NewCapturer(iface, filter string, parser Parser) *Capturer {
	return &Capturer{iface: iface, filter: filter, parser: parser}
}

func (c *Capturer) Start() {
	handle, err := pcap.OpenLive(c.iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("Failed to open device: %v", err)
	}
	defer handle.Close()

	if c.filter != "" {
		if err := handle.SetBPFFilter(c.filter); err != nil {
			log.Fatalf("Failed to set filter: %v", err)
		}
		fmt.Printf("Applied filter: %s\n", c.filter)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		c.parser.Parse(packet)
	}
}

func StartCapture(iface string) (*pcap.Handle, error) {
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		return nil, err
	}
	return handle, nil
}

type PacketParser interface {
	Parse(gopacket.Packet)
}

func CapturePackets(handle *pcap.Handle, parser PacketParser) {
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		parser.Parse(packet)
	}
}
