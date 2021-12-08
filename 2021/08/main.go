package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
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

func parseEntry(s string) ([]string, []string) {
	parts := strings.Split(s, " | ")
	signals := strings.Fields(parts[0])
	outputs := strings.Fields(parts[1])

	return signals, outputs
}

func solvePart1(in []string) int {
	count := 0
	for _, entry := range in {
		_, outputValues := parseEntry(entry)
		dbg("outputs: %v", outputValues)

		for _, v := range outputValues {
			switch len(v) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}

	return count
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	part1 = solvePart1(lines)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
