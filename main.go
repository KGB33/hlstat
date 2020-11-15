package main

import (
	"hlstat/cli"
	"os"
)

func main() {
	os.Exit(cli.CLI(os.Args[1:]))
}
