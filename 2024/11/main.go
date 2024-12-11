package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	. "github.com/agimenez/adventofcode/utils"
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
	part1 = solve1(lines[0], 25)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

func solve1(s string, count int) int {
	stones := strings.Fields(s)
	dbg("Initial stones: %v", stones)
	for i := 0; i < count; i++ {
		stones = blink(stones)
		dbg("New stones (%d): %v", i+1, stones)
	}
	return len(stones)
}

func blink(stones []string) []string {
	res := []string{}

	for _, s := range stones {
		val := ToInt(s)
		if s == "0" {
			res = append(res, "1")
		} else if len(s)%2 == 0 {
			first := ToInt(s[:len(s)/2])
			second := ToInt(s[len(s)/2:])
			res = append(res, fmt.Sprintf("%d", first))
			res = append(res, fmt.Sprintf("%d", second))

		} else {
			val *= 2024
			res = append(res, fmt.Sprintf("%d", val))
		}
	}

	return res
}
