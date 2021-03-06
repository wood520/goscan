package utils

import (
	"net"
	"net/http"
)

func Connected() bool {
	_, err := http.Get("https://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}

// ---------------------------------------------------------------------------------------
// CIDRs
// ---------------------------------------------------------------------------------------
// Parse a string and returns the corresponding CIDR
func ParseCIDR(s string) string {
	_, ipv4Net, err := net.ParseCIDR(s)
	if err != nil {
		return ""
	}
	return ipv4Net.String()
}

// Returns all the addresses of the local network interfaces
func ParseLocalIP() map[string]string {
	// Returns a Map of interface:subnet
	res := make(map[string]string)

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			res[i.Name] = addr.String()
			break
		}
	}
	return res
}
