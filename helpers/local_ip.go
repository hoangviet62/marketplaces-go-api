package helpers

import (
	"net"

	log "github.com/sirupsen/logrus"
)

// Get preferred outbound ip of this machine
func LocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
