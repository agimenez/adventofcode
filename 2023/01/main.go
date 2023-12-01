package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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
func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	for l := range lines {
		first, last := -1, -1
		for _, r := range lines[l] {
			if !unicode.IsDigit(r) {
				continue
			}

			v, _ := strconv.Atoi(string(r))

			// Initial case: no first or last
			if first == -1 {
				first = v
			}

			// if it's a digit, will always be last
			last = v
		}

		dbg("l: %s, %d%d", lines[l], first, last)
		part1 += 10*first + last
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
