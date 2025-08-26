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

	var srcIP, dstIP string
	var srcPort, dstPort string
	var label string

	if netLayer := packet.NetworkLayer(); netLayer != nil {
		switch ipLayer := netLayer.(type) {
		case *layers.IPv4:
			dst := ipLayer.DstIP
			label = classifyIP(dst)
			dstStr := dst.String()
			if p.doResolve && isPublicIP(dst) {
				dstStr = p.resolver.Resolve(dstStr)
			}
			p.stats.AddIPStats(fmt.Sprintf("[%s] %s", label, dstStr), len(packet.Data()))
			srcIP = ipLayer.SrcIP.String()
			dstIP = ipLayer.DstIP.String()

		case *layers.IPv6:
			dst := ipLayer.DstIP
			label = classifyIP(dst)
			dstStr := dst.String()
			if p.doResolve && isPublicIP(dst) {
				dstStr = p.resolver.Resolve(dstStr)
			}
			p.stats.AddIPStats(fmt.Sprintf("[%s] %s", label, dstStr), len(packet.Data()))
			srcIP = ipLayer.SrcIP.String()
			dstIP = ipLayer.DstIP.String()
		}
	}

	if transportLayer := packet.TransportLayer(); transportLayer != nil {
		switch layer := transportLayer.(type) {
		case *layers.TCP:
			p.stats.IncrementTCP()
			srcPort = fmt.Sprintf("%d", layer.SrcPort)
			dstPort = fmt.Sprintf("%d", layer.DstPort)

		case *layers.UDP:
			p.stats.IncrementUDP()
			srcPort = fmt.Sprintf("%d", layer.SrcPort)
			dstPort = fmt.Sprintf("%d", layer.DstPort)

		default:
			if transportLayer.LayerType() == layers.LayerTypeICMPv4 || transportLayer.LayerType() == layers.LayerTypeICMPv6 {
				p.stats.IncrementICMP()
			} else {
				p.stats.IncrementOther()
			}
		}
	}

	if packet.TransportLayer() == nil {
		p.stats.IncrementOther()
	}

	src := srcIP
	dst := dstIP
	if srcPort != "" {
		src = fmt.Sprintf("%s:%s", src, srcPort)
	}
	if dstPort != "" {
		dst = fmt.Sprintf("%s:%s", dst, dstPort)
	}
	if src != "" && dst != "" {
		p.stats.AddTraffic(src, dst)
	}
}
