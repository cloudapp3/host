package info

import (
	"fmt"
	"net"
)

var (
	clientIP = ""
)

func getClientIP() string {
	return clientIP
}
func HostIP() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("host ip error " + err.Error())
		return
	}
	for _, i := range ifaces {
		address, err := i.Addrs()
		if err != nil{
			fmt.Println("get address error " + err.Error())
		}
		for _, addr := range address {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			clientIP = ip.String()
		}
	}
}
