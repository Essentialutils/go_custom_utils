package go_custom_utils

import (
	"fmt"
	"net"
)

// GetWifiIP returns the IPv4 address of the Wi-Fi interface
func GetWifiIP() []map[string]string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var result []map[string]string
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// Check if the IP is an IPv4 address and not a loopback address
			if ip != nil && ip.To4() != nil && !ip.IsLoopback() {
				result = append(result, map[string]string{
					"name": i.Name,
					"ip":   ip.String(),
				})
			}
		}
	}
	return result
}
