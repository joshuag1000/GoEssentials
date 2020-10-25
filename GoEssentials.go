package goessentials

import (
	"fmt"
	"net"
	"os"
)

// General Subroutines for all my programs.

// Hello Prints hi into the golang window
func Hello() {
	fmt.Println("Hi")
}

// GetServerIp Gets the IP address of the device. If the device has more than one IP it will select the first IP.
func GetServerIP(ipNum int) string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
	var ips [255]string
	var i int
	i = 0
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips[i] = ipnet.IP.String()
				i = i + 1
			}
		}
	}
	return ips[ipNum]
}
