package cmd

import (
	"fmt"

	"github.com/luique16/reverse-proxy/internal/cli"
	"github.com/luique16/reverse-proxy/internal/proxy_core"
	"github.com/luique16/reverse-proxy/internal/server"
)

func Execute() {
	args := cli.ParseArgs()

	if args.HelpOpt {
		cli.Help(args.WrongUsage)
		return
	}

	var port int

	if args.PortOpt {
		port = args.Port
	} else {
		port = 3000
	}

	var num int

	if args.NumOpt {
		num = args.Num
	} else {
		num = 3
	}

	ready := make(chan bool)

	ports, err := proxy_core.Start(port, num, args.LogOpt, server.Server, ready)

	if err != nil {
		fmt.Println(err)
		return
	}
	
	for range ports {
		<-ready
	}

	fmt.Println("Proxy started ✅")

	select {}
}
