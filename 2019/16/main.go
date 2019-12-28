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
		log.Printf(fmt+"\n", v...)
	}
}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

func pattern(phase int) []int {
	var basePattern = []int{0, 1, 0, -1}
	p := []int{}

	for _, n := range basePattern {
		for i := phase; i > 0; i-- {
			p = append(p, n)
		}
	}

	return p
}

func main() {

	var in string
	fmt.Scan(&in)

	fmt.Printf("Part one: %#v\n", in)

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func mod(a, b int) int {
	return (a%b + b) % b
}
