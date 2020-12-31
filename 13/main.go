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

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
