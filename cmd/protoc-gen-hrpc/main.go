package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.2.0"

func main() {
	versionFlag := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	gen := newGenerator()
	gen.Run()
}
