package cli

import (
	"fmt"
)

func Help(wrongUsage bool) {
	if wrongUsage {
		fmt.Println("Wrong usage!!!")
	}

	fmt.Println("Usage: reverse-proxy [OPTIONS]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -p, --port <port>    Port to listen on (default: 3000)")
	fmt.Println("  -n, --num <num>      Number of instances to run (default: 3)")
	fmt.Println("  -P, --path <path>    Path to server (default: test server)")
	fmt.Println("  -l, --log            Enable logging")
	fmt.Println("  -h, --help           Show this help message and exit")
	fmt.Println()
}
