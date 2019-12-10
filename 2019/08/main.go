package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	debug int
)

func dbg(level int, fmt string, v ...interface{}) {
	if debug >= level {
		log.Printf(fmt, v...)
	}
}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

func main() {

	var in string
	fmt.Scan(&in)

	log.Printf("Max output: %v", in)

}
