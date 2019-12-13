package main

import (
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

type coord struct {
	x, y int
}

func main() {

	var in string
	fmt.Scan(&in)

}
