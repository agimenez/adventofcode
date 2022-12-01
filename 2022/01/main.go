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

func maxCals(in []string) int {
	elfSum := 0
	maxSum := 0

	for _, c := range in {
		if c == "" {
			if elfSum > maxSum {
				maxSum = elfSum
			}
			elfSum = 0
			continue
		}

		cals, _ := strconv.Atoi(c)
		elfSum += cals

	}

	return maxSum
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	dbg("lines: %#v", lines)

	part1 = maxCals(lines)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
