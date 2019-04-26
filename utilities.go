package utilities

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func GetIPMAC() (mac string, ip string) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}

	var probeIP, currentNetworkHardwareName string

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				probeIP = ipnet.IP.String()
			}
		}
	}

	interfaces, _ := net.Interfaces()
	for _, interf := range interfaces {

		if addrs, err := interf.Addrs(); err == nil {
			for _, addr := range addrs {
				if strings.Contains(addr.String(), probeIP) {
					currentNetworkHardwareName = interf.Name
				}
			}
		}
	}

	netInterface, err := net.InterfaceByName(currentNetworkHardwareName)

	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}

	name := netInterface.Name
	macAddress := netInterface.HardwareAddr

	fmt.Println("Hardware name : ", name)
	fmt.Println("MAC address : ", macAddress)

	hwAddr, err := net.ParseMAC(macAddress.String())
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(-1)
	}

	fmt.Printf("Physical hardware address : %s \n", hwAddr.String())
	return hwAddr.String(), probeIP

}
