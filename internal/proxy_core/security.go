package proxy_core

import (
	"net"
	"strconv"
)

func IsPortAvailable(port int) bool {
	conn, err := net.Dial("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(port)))

	if err != nil {
		return true
	}

	defer conn.Close()

	return false
}
