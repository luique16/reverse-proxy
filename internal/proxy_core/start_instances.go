package proxy_core

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func StartInstances(port, num int, server *http.ServeMux, log bool) ([]int, error) {
	if !IsPortAvailable(port) {
		return nil, errors.New("Port not available")
	}

	var counter int
	var i int

	counter = 0
	i = 1

	var ports []int

	ready := make(chan bool)

	for counter < num {
		if !IsPortAvailable(port + i) {
			i += 1
			continue
		}

		go func(port int) {
			if log {
				fmt.Println("Instance running on port " + strconv.Itoa(port) + " ✅")
			}

			ready <- true
			http.ListenAndServe(":" + strconv.Itoa(port), server)
		}(port + i)

		ports = append(ports, port + i)

		counter += 1
		i += 1
	}

	for range ports {
		<-ready
	}

	return ports, nil
}
