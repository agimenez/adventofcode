package main

import (
	"bufio"
	"log"
	"os"
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

	s := bufio.NewScanner(os.Stdin)
	part1, part2 := 0, 0
	for s.Scan() {
		l := s.Text()

		dbg("Line: %v\n", l)
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
