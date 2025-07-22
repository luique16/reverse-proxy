package main

import (
	"github.com/luique16/reverse-proxy"
	"github.com/luique16/reverse-proxy/internal/cli"
	"github.com/luique16/reverse-proxy/internal/server"
)

func main() {
	port, num, log := cli.FilterArgs(cli.ParseArgs())

	proxy.Run(port, num, server.NewServer(), log)
}
