package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func run() error {
	return nil
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	err := run()
	if err != nil {
		log.Fatalf("%+v", err)
	}
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] HOSTNAME \n", os.Args[0])
	flag.PrintDefaults()
}
