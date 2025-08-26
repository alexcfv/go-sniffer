package parser

import "net"

func isPublicIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.IsPrivate() {
		return false
	}
	return true
}

func classifyIP(ip net.IP) string {
	if ip.IsLoopback() {
		return "LOOPBACK"
	}
	if ip.IsPrivate() {
		return "LOCAL"
	}
	if ip.IsMulticast() {
		return "MULTICAST"
	}
	if ip.IsUnspecified() {
		return "UNSPECIFIED"
	}
	return "PUBLIC"
}
