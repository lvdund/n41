package n41types

import (
	"fmt"
	"internal/bytealg"
	"internal/itoa"
	"net"
)

type SbiAdrr struct {
	IP   net.IP `json:"ip"`
	Port int    `json:"port"`
	Zone string `json:"zone,omitempty"`
}

func (s SbiAdrr) GetAddr() string {
	return fmt.Sprintf("%s:%d", s.IP, s.Port)
}

func (a *SbiAdrr) String() string {
	if a == nil {
		return "<nil>"
	}
	ip := ipEmptyString(a.IP)
	if a.Zone != "" {
		return JoinHostPort(ip+"%"+a.Zone, itoa.Itoa(a.Port))
	}
	return JoinHostPort(ip, itoa.Itoa(a.Port))
}

func ipEmptyString(ip net.IP) string {
	if len(ip) == 0 {
		return ""
	}
	return ip.String()
}
func JoinHostPort(host, port string) string {
	// We assume that host is a literal IPv6 address if host has colons.
	if bytealg.IndexByteString(host, ':') >= 0 {
		return "[" + host + "]:" + port
	}
	return host + ":" + port
}
