package main

import (
	"errors"
	"fmt"
	"net"
)

type LocalInfo struct {
	ip      string
	macAddr string
}

func GetLocalIPAndMacAddr() (*LocalInfo, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		// interface down or loopback address
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok { // check the address type
				if ipnet.IP.To4() != nil {
					return &LocalInfo{
						ip:      ipnet.IP.String(),
						macAddr: iface.HardwareAddr.String(),
					}, nil
				}
			}
		}
	}
	return nil, errors.New("No local ip and mac")
}

func main() {
	info, err := GetLocalIPAndMacAddr()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", info)
}
