package t468

import (
	"fmt"
	"net"
	"strings"
)

func validIPAddress(IP string) string {
	i := net.ParseIP(IP)
	if i != nil {
		if i.To4() != nil {
			if i.String() != IP {
				return "Neither"
			}
			return "IPv4"
		}

		tmp := strings.Split(IP, ":")
		for i := range tmp {
			if len(tmp[i]) > 4 || len(tmp[i]) == 0 {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}

func ExampleValidIPAddress() {
	ips := []string{"172.16.254.1", "2001:0db8:85a3:0:0:8A2E:0370:7334", "256.256.256.256", "2001:0db8:85a3:0:0:8A2E:0370:7334:", "1e1.4.5.6"}
	for _, ip := range ips {
		fmt.Println(validIPAddress(ip))
	}

	// Output:
	// IPv4
	// IPv6
	// Neither
	// Neither
	// Neither
}
