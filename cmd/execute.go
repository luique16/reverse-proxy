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
		if args.Port < 1 {
			fmt.Println("Port must be greater than 0")
			return
		}

		port = args.Port
	} else {
		port = 3000
	}

	var num int

	if args.NumOpt {
		if args.Num < 1 {
			fmt.Println("Number of instances must be greater than 0")
			return
		}

		num = args.Num
	} else {
		num = 3
	}

	ports, err := proxy_core.StartInstances(port, num, server.NewServer(port), args.LogOpt)

	if err != nil {
		fmt.Println(err)
		return
	}

	ready := make(chan bool)

	go proxy_core.Redirector(ports, port, args.LogOpt, ready)

	<-ready

	fmt.Println("Proxy started ✅")

	select {}
}
