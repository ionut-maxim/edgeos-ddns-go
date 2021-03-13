package main

import (
	"log"
	"net"
	"os/exec"
)

func getInterfaceIP(ifaceName string) net.IP {
	iface, err := net.InterfaceByName(ifaceName)

	if err != nil {
		log.Fatal(err)
	}

	addrs, err := iface.Addrs()

	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			return ipnet.IP
		}
	}

	return nil
}

func ipInRange(ipAddr net.IP, ipRange string) bool {
	_, network, err := net.ParseCIDR(ipRange)

	if err != nil {
		log.Fatal(err)
	}

	if network.Contains(ipAddr) {
		return true
	}

	return false
}

func runProcess(ifaceName string, command string) {
	cmd := exec.Command("/opt/vyatta/bin/vyatta-op-cmd-wrapper", command, "interface", ifaceName)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func restartInterface(ifaceName string) {
	runProcess(ifaceName, "disconnect")
	runProcess(ifaceName, "connect")
}
