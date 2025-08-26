package parser

import (
	"fmt"

	"github.com/alexcfv/go-sniffer/resolver"
	"github.com/alexcfv/go-sniffer/stats"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Parser struct {
	stats     *stats.Stats
	resolver  *resolver.Resolver
	doResolve bool
}

func NewParser(stats *stats.Stats, r *resolver.Resolver, doResolve bool) *Parser {
	return &Parser{stats: stats, resolver: r, doResolve: doResolve}
}

func (p *Parser) Parse(packet gopacket.Packet) {
	p.stats.IncrementPackets(len(packet.Data()))

	if netLayer := packet.NetworkLayer(); netLayer != nil {
		switch ipLayer := netLayer.(type) {
		case *layers.IPv4:
			dst := ipLayer.DstIP
			label := classifyIP(dst)
			dstStr := dst.String()
			if p.doResolve && isPublicIP(dst) {
				dstStr = p.resolver.Resolve(dstStr)
			}
			p.stats.AddIPStats(fmt.Sprintf("[%s] %s", label, dstStr), len(packet.Data()))

		case *layers.IPv6:
			dst := ipLayer.DstIP
			label := classifyIP(dst)
			dstStr := dst.String()
			if p.doResolve && isPublicIP(dst) {
				dstStr = p.resolver.Resolve(dstStr)
			}
			p.stats.AddIPStats(fmt.Sprintf("[%s] %s", label, dstStr), len(packet.Data()))
		}
	}

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
