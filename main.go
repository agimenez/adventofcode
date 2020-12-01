package main

import (
	"fmt"
	"log"
)

const (
	debug = false
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func main() {
	var in string

	fmt.Scan(&in)

	log.Printf("%v", in)
}
