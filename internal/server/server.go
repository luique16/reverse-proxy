package server

import (
	"fmt"
	"strconv"
	"net/http"
)

func Server(port int, ready chan<- bool) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	fmt.Println("Instance running on port " + strconv.Itoa(port) + " ✅")

	ready <- true

	err := http.ListenAndServe(":" + strconv.Itoa(port), mux)

	if err != nil {
		return err
	}

	return nil
}
