package proxy_core

import (
	"errors"
)

func Start(port, num int, log bool, server func(int, chan<- bool) error, ready chan<- bool) ([]int, error) {
	if !IsPortAvailable(port) {
		return nil, errors.New("Port not available")
	}

	var counter int
	var i int

	counter = 0
	i = 1

	var ports []int


	for counter < num {
		if !IsPortAvailable(port + i) {
			i += 1
			continue
		}
		go server(port + i, ready)

		counter += 1
		i += 1
		ports = append(ports, port + i)
	}

	return ports, nil
}
