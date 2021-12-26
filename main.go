package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		port     int
		firstTTL int
		maxTTL   int
		retry    int
	)
	flag.IntVar(&port, "port", 33434, "port")
	flag.IntVar(&firstTTL, "f", 1, "first TTL")
	flag.IntVar(&maxTTL, "m", 64, "max TTL")
	flag.IntVar(&retry, "n", 3, "retry count")
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
		os.Exit(1)
	}
	hostname := flag.Arg(0)

	err := traceroute(hostname, port, firstTTL, maxTTL, retry)
	if err != nil {
		log.Fatalf("%+v", err)
	}
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] HOSTNAME \n\n", os.Args[0])
	flag.PrintDefaults()
}
