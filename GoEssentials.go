package goessentials

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func Hello() {
	fmt.Println("Hi")
}

func GetServerIp(ipNum int) string {
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
