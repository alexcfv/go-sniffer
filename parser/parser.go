package parser

import (
	"github.com/alexcfv/go-sniffer/stats"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Parser struct {
	stats *stats.Stats
}

func NewParser(stats *stats.Stats) *Parser {
	return &Parser{stats: stats}
}

func (p *Parser) Parse(packet gopacket.Packet) {
	p.stats.IncrementPackets(len(packet.Data()))

	if netLayer := packet.NetworkLayer(); netLayer != nil {
		src, dst := netLayer.NetworkFlow().Endpoints()
		p.stats.AddTraffic(src.String(), dst.String())

		switch netLayer.LayerType() {
		case layers.LayerTypeIPv4, layers.LayerTypeIPv6:
			if transportLayer := packet.TransportLayer(); transportLayer != nil {
				switch transportLayer.LayerType() {
				case layers.LayerTypeTCP:
					p.stats.IncrementTCP()
				case layers.LayerTypeUDP:
					p.stats.IncrementUDP()
				case layers.LayerTypeICMPv4, layers.LayerTypeICMPv6:
					p.stats.IncrementICMP()
				default:
					p.stats.IncrementOther()
				}
			}
		default:
			p.stats.IncrementOther()
		}
	} else {
		p.stats.IncrementOther()
	}
}
