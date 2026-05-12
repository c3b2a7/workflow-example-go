package main

import (
	"flag"
	"fmt"
)

var (
	Version   = "unknown version"
	BuildTime = "unknown time"
)

var (
	showVersion bool
)

func init() {
	flag.BoolVar(&showVersion, "version", false, "show version information")
	flag.Parse()
}

func main() {
	if showVersion {
		fmt.Printf("%s, %s\n", Version, BuildTime)
		return
	}

	fmt.Println("Hello, World!")
}
