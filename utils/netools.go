package utils

import (
	"go-ecm/pkg/log"
	"net"
)

func GetLocalIP() []string {
	ifs, err := net.Interfaces()
	if err != nil {
		log.Error("get local ip err.")
	}

	var ips []string
	for i := 0; i < len(ifs); i++ {
		if (ifs[i].Flags&net.FlagUp) != 0 &&
			ifs[i].Name != "docker0" && ifs[i].Name != "docker_gwbridge" {
			addrs, _ := ifs[i].Addrs()

			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ips = append(ips, ipnet.IP.String())
					}
				}
			}
		}
	}

	return ips
}
