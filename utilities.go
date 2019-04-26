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
		fmt.Println(err)
	}

	var probeIP, currentNetworkHardwareName string

	for _, address := range addrs {

		// check the address type and if it is not a loopback the display it
		// = GET LOCAL IP ADDRESS
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("Current IP address : ", ipnet.IP.String())
				probeIP = ipnet.IP.String()
			}
		}
	}
	interfaces, _ := net.Interfaces()
	for _, interf := range interfaces {

		if addrs, err := interf.Addrs(); err == nil {
			for index, addr := range addrs {
				fmt.Println("[", index, "]", interf.Name, ">", addr)

				// only interested in the name with current IP address
				if strings.Contains(addr.String(), probeIP) {
					fmt.Println("Use name : ", interf.Name)
					currentNetworkHardwareName = interf.Name
				}
			}
		}
	}

	fmt.Println("------------------------------")

	// extract the hardware information base on the interface name
	// capture above
	netInterface, err := net.InterfaceByName(currentNetworkHardwareName)

	if err != nil {
		fmt.Println(err)
	}

	name := netInterface.Name
	macAddress := netInterface.HardwareAddr

	fmt.Println("Hardware name : ", name)
	fmt.Println("MAC address : ", macAddress)

	// verify if the MAC address can be parsed properly
	hwAddr, err := net.ParseMAC(macAddress.String())
	if err != nil {
		fmt.Println("No able to parse MAC address : ", err)
		os.Exit(-1)
	}

	fmt.Printf("Physical hardware address : %s \n", hwAddr.String())
	return hwAddr.String(), probeIP

}
