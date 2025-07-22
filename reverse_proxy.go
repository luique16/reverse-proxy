package reverse_proxy

import (
	"fmt"
	"net/http"

	"github.com/luique16/reverse-proxy/internal/proxy_core"
)

func ReverseProxy(port int, num int, server *http.ServeMux, log bool) {
	ports, err := proxy_core.StartInstances(port, num, server, log)

	if err != nil {
		fmt.Println(err)
		return
	}

	ready := make(chan bool)

	go proxy_core.Redirector(ports, port, log, ready)

	<-ready

	fmt.Println("Proxy started ✅")

	select {}
}
