package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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
	flag.Parse()
}

func findEarliest(ts int, buses map[int]int) (int, int) {
	cur := ts
	for {
		for i := range buses {
			if cur%i == 0 {
				return cur, i
			}
		}

		cur++
	}
}

// This seems to be solves using Chinese Remainder Theorem, but the awesome simple
// solution from lizthegrey is more the "intuitive" form I had in mind, but couldn't
// quite get. The following code is copied (and slightly modified to make it more
// meaningful for me) from https://github.com/lizthegrey/adventofcode/blob/main/2020/day13.go
// What this does is to define timestamp steps for the search (tsStep), which is
// basically the LCM of the bus IDs (which are coprime).
// We add the current found timestamp to the LCM of the previous buses until the
// current one in a loop to search for a timestamp for which ts + bus offset is a
// multiple of the bus ID, which means that the condition is satisfied
func findEarliestAndSubsequent(buses map[int]int) int {
	ts := 0
	tsStep := 1

	for bus, offset := range buses {
		for (ts+offset)%bus != 0 {
			ts += tsStep
		}

		tsStep *= bus
	}

	return ts
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	ts, err := strconv.Atoi(lines[0])
	if err != nil {
		panic("could not parse timestamp")
	}

	buses := map[int]int{}
	times := strings.Split(lines[1], ",")
	for i, l := range times {
		if l == "x" {
			continue
		}

		b, err := strconv.Atoi(l)
		if err != nil {
			panic("could not parse busID")
		}
		buses[b] = i
	}
	t, b := findEarliest(ts, buses)
	part1 = (t - ts) * b

	part2 = findEarliestAndSubsequent(buses)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
