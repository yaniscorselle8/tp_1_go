package main

import (
	"flag"
	"fmt"
	"os"
)

const VERSION = "1.0"

func main() {

	version := flag.Bool("version", true, "version")
	flag.Parse()
	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

}
