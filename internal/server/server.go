package server

import (
	"fmt"
	"strconv"
	"net/http"
)

func ConfigureServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
}

func Server(port int, ready chan<- bool) error {
	fmt.Println("Instance running on port " + strconv.Itoa(port) + " ✅")

	ready <- true

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)

	if err != nil {
		return err
	}

	return nil
}
