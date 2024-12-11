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
	part2 = solve1(lines[0], 75)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

func solve1(s string, blinks int) int {
	stones := getStones(s)
	dbg("Initial stones: %v", stones)
	for i := 0; i < blinks; i++ {
		stones = blink(stones)
		dbg("New stones (%d): %v", i+1, stones)
	}
	return count(stones)
}

func count(s map[string]int) int {
	count := 0
	for _, v := range s {
		count += v
	}

	return count
}

func getStones(s string) map[string]int {
	m := map[string]int{}
	for _, v := range strings.Fields(s) {
		m[v]++
	}

	return m
}

func blink(stones map[string]int) map[string]int {
	l := []string{}
	currentStones := map[string]int{}
	for k, v := range stones {
		if v == 0 {
			continue
		}
		currentStones[k] = v

		l = append(l, k)
	}

	dbg("blink: %v", l)
	for _, v := range l {
		counts := currentStones[v]
		val := ToInt(v)
		dbg(" -> Stone %v, counts %v", v, counts)
		if v == "0" {
			stones["0"] -= counts
			stones["1"] += counts
		} else if len(v)%2 == 0 {
			first := ToInt(v[:len(v)/2])
			second := ToInt(v[len(v)/2:])
			stones[v] -= counts

			stones[fmt.Sprintf("%d", first)] += counts
			stones[fmt.Sprintf("%d", second)] += counts

		} else {
			val *= 2024
			stones[fmt.Sprintf("%d", val)] += counts
			stones[v] -= counts
		}
	}
	dbg("return: %v", stones)

	return stones
}
