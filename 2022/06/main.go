package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var (
	debug bool
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}

func detectStart(s string, wSize int) int {
	for i := range s {
		set := make(map[byte]bool, wSize)
		for j := i; j < i+wSize; j++ {
			set[s[j]] = true
		}

		if len(set) == wSize {
			return i + wSize
		}
	}

	return 0
}

func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	part1 = detectStart(string(p), 4)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
