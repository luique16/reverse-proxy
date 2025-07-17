package cmd

import (
	"fmt"
	"github.com/luique16/reverse-proxy/internal/cli"
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

	fmt.Printf("Running %d instances after port %d\n", num, port)
}
