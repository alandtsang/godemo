package main

import (
	"fmt"
	"net"
	"os"
)

var Cidrs []*net.IPNet

// https://en.wikipedia.org/wiki/Reserved_IP_addresses
var cidrs = [15]string{
	"0.0.0.0/8",       // 0.0.0.0--0.255.255.255
	"10.0.0.0/8",      // 10.0.0.0--10.255.255.255
	"100.64.0.0/10",   // 100.64.0.0--100.127.255.255
	"127.0.0.0/8",     // 127.0.0.0--127.255.255.255
	"169.254.0.0/16",  // 169.254.0.0-169.254.255.255
	"172.16.0.0/12",   // 172.16.0.0--172.31.255.255
	"192.0.0.0/24",    // 192.0.0.0--192.0.0.255
	"192.0.2.0/24",    // 192.0.2.0--192.0.2.255
	"192.88.99.0/24",  // 192.88.99.0--192.88.99.255
	"192.168.0.0/16",  // 192.168.0.0--192.168.255.255
	"198.18.0.0/15",   // 192.18.0.0--192.19.255.255
	"198.51.100.0/24", // 198.51.100.0--198.51.100.255
	"203.0.113.0/24",  // 203.0.113.0--203.0.113.255
	"224.0.0.0/4",     // 224.0.0.0--239.255.255.255
	"240.0.0.0/4",     // 240.0.0.0--255.255.255.254
}

func ParseCIDR() {
	for _, cidr := range cidrs {
		_, c, err := net.ParseCIDR(cidr)
		if err != nil {
			fmt.Println("ParseCIDR", err)
			os.Exit(1)
		}
		Cidrs = append(Cidrs, c)
	}
}

func IsCIDR(srcIP string) bool {
	for _, cidr := range Cidrs {
		if cidr.Contains(net.ParseIP(srcIP)) {
			return true
		}
	}
	return false
}

func main() {
	ParseCIDR()

	ips := []string{
		"192.168.0.101",
		"100.64.157.228",
		"119.128.220.93",
	}

	for _, ip := range ips {
		ok := IsCIDR(ip)
		fmt.Printf("%q, %v\n", ip, ok)
	}
}
