package ip

import (
	"fmt"
	"log"
	"net"
)

// GET preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	return localAddr.IP.String()
}

func GetLocalIP() (ip string) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		ip = "0.0.0.0"
		return
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}
	if len(ips) != 0 {
		ip = ips[0]
	} else {
		ip = "0.0.0.0"
	}
	return
}
