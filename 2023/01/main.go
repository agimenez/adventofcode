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

var words = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
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
		dbg("l: %s", lines[l])
		for i, r := range lines[l] {
			if unicode.IsDigit(r) {
				dbg("digit: %c", r)

				v, _ := strconv.Atoi(string(r))

				// Initial case: no first or last
				if first == -1 {
					first = v
				}

				// if it's a digit, will always be last
				last = v
			} else { // handle part 2
				dbg("word: %s", lines[l][i:])
				for w, d := range words {
					if strings.HasPrefix(lines[l][i:], w) {
						if first == -1 {
							first = d
						}

						last = d
					}
				}

			}
			dbg("first: %v, last: %v", first, last)
		}

		dbg("l: %s, %d%d", lines[l], first, last)
		part1 += 10*first + last
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
