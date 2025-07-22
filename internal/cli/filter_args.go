package cli

import (
	"fmt"
	"os"
)


func FilterArgs(args Args) (int, int, bool) {
	if args.HelpOpt {
		Help(args.WrongUsage)
		os.Exit(0)
	}

	var port int

	if args.PortOpt {
		if args.Port < 1 {
			fmt.Println("Port must be greater than 0")
			os.Exit(0)
		}

		port = args.Port
	} else {
		port = 3000
	}

	var num int

	if args.NumOpt {
		if args.Num < 1 {
			fmt.Println("Number of instances must be greater than 0")
			os.Exit(0)
		}

		num = args.Num
	} else {
		num = 3
	}

	return port, num, args.LogOpt

}
