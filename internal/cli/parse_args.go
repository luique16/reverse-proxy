package cli

import (
	"os"
	"strconv"
)

type Args struct {
	PortOpt    bool
	Port       int
	NumOpt     bool
	Num        int
	LogOpt     bool
	HelpOpt    bool
	WrongUsage bool
}

func ParseArgs() Args {
	var args Args

	if len(os.Args) == 1 {
		args.HelpOpt = true
		args.WrongUsage = true
		return args
	}

	var lookNext bool

	array := os.Args[1:]

	for i, arg := range array {
		var err error

		switch arg {
		case "-p", "--port":
			args.PortOpt = true
			if i+1 < len(array) {
				args.Port, err = strconv.Atoi(array[i+1])
				lookNext = true
			} else {
				args.WrongUsage = true
				args.HelpOpt = true
				return args
			}
		case "-n", "--num":
			args.NumOpt = true
			if i+1 < len(array) {
				args.Num, err = strconv.Atoi(array[i+1])
				lookNext = true
			} else {
				args.WrongUsage = true
				args.HelpOpt = true
				return args
			}
		case "-h", "--help":
			args.HelpOpt = true
		case "-l", "--log":
			args.LogOpt = true
		default:
			if lookNext {
				lookNext = false
			} else {
				args.WrongUsage = true
				args.HelpOpt = true
				return args
			}
		}
		
		if err != nil {
			args.WrongUsage = true
			args.HelpOpt = true
			return args
		}
	}

	return args
}
