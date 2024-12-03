package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/agimenez/adventofcode/utils"
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
	for _, l := range lines {
		part1 += getMuls(l)
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

var pattern = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func getMuls(s string) int {
	total := 0
	matches := pattern.FindAllStringSubmatch(s, -1)
	dbg("%v", matches)
	for _, match := range matches {
		dbg(" - %+v", match)
		total += utils.ToInt(match[1]) * utils.ToInt(match[2])
	}

	return total
}
