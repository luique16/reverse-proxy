package proxy_core

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"sync"
)

type Proxy struct {
	backends []*httputil.ReverseProxy
	ports    []int
	position int
	mu       sync.Mutex
	log      bool
}

func NewProxy(ports []int, log bool) *Proxy {
	var backends []*httputil.ReverseProxy

	for _, port := range ports {
		target := "http://127.0.0.1:" + strconv.Itoa(port)
		url, err := url.Parse(target)
		if err != nil {
			panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(url)
		backends = append(backends, proxy)
	}

	return &Proxy{
		backends: backends,
		ports:    ports,
		position: 0,
		log:      log,
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.mu.Lock()
	proxy := p.backends[p.position]
	if p.log {
		fmt.Println("Redirecting to", p.ports[p.position])
	}
	p.position = (p.position + 1) % len(p.backends)
	p.mu.Unlock()

	proxy.ServeHTTP(w, r)
}

func Redirector(ports []int, port int, log bool, ready chan<- bool) {
	proxy := NewProxy(ports, log)

	http.Handle("/", proxy)

	ready <- true

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		panic(err)
	}
}

