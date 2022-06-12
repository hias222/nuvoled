package udpserver

import (
	"fmt"
	"net"
)

func GetIpaddress() string {
	ifaces, err := net.Interfaces()
	// handle err
	if err != nil {
		fmt.Println("error get IP")
		return ""
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		// handle err
		if err != nil {
			fmt.Println("error get iface")
			return ""
		}
		for _, addr := range addrs {
			var ip net.IP
			var name string
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				name = v.IP.String()
			case *net.IPAddr:
				ip = v.IP
				name = v.IP.String()
			}

			if len(ip) > 15 {
				if ip[12] == 169 {
					if ip[13] == 254 {
						fmt.Println("found self assigned ", ip)
						return name
					}
				}
			}

			// process IP address
		}
	}
	return ""
}

func GetLocaladdress() string {
	ifaces, err := net.Interfaces()
	// handle err
	if err != nil {
		fmt.Println("error get IP")
		return ""
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		// handle err
		if err != nil {
			fmt.Println("error get iface")
			return ""
		}
		for _, addr := range addrs {
			var ip net.IP
			var name string
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				name = v.IP.String()
			case *net.IPAddr:
				ip = v.IP
				name = v.IP.String()
			}

			if len(ip) > 15 {
				if ip[12] == 192 {
					if ip[13] == 168 {
						fmt.Println("found internal 192er ", ip)
						return name
					}
				}
			}

			// process IP address
		}
	}
	return ""
}
