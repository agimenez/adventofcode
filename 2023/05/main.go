package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

type rangeMap struct {
	src  int
	dst  int
	size int
}
type categoryMap struct {
	ranges []rangeMap
	next   string
}

type almanac struct {
	seeds []int
	maps  map[string]*categoryMap
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)

	return n
}

func parseAlmanac(lines []string) almanac {
	a := almanac{
		seeds: []int{},
		maps:  make(map[string]*categoryMap),
	}

	var curMap string
	for _, l := range lines {
		if l == "" {
			continue
		}

		if strings.HasPrefix(l, "seeds:") {
			vals := strings.Split(l, ": ")[1]
			for _, v := range strings.Fields(vals) {
				n, _ := strconv.Atoi(v)
				a.seeds = append(a.seeds, n)
			}
			continue
		}

		// Start a new map
		if !unicode.IsDigit(rune(l[0])) {
			parts := strings.Split(strings.Fields(l)[0], "-")
			src, dst := parts[0], parts[2]
			dbg("Found new map: %s -> %s", src, dst)
			a.maps[src] = &categoryMap{
				ranges: []rangeMap{},
				next:   dst,
			}
			curMap = src

			continue
		}

		// Keep parsing current map
		parts := strings.Fields(l)
		dst, src, length := atoi(parts[0]), atoi(parts[1]), atoi(parts[2])
		dbg(" -> %s[%d] -> %d (x%d)", curMap, src, dst, length)
		rm := rangeMap{
			src:  src,
			dst:  dst,
			size: length,
		}

		// why do I have to do this in 2023?
		a.maps[curMap].ranges = append(a.maps[curMap].ranges, rm)

	}

	return a
}

func (r rangeMap) translate(p int) int {
	gap := p - r.src
	if gap >= 0 && gap < r.size {
		return r.dst + gap
	}

	return p
}

func (a almanac) getNextPosition(curMap string, curPos int) int {
	ranges := a.maps[curMap].ranges
	for _, r := range ranges {
		newPos := r.translate(curPos)
		if newPos != curPos {
			return newPos
		}
	}

	return curPos
}
func (a almanac) traceLocation(startMap string, startPos int) (int, string) {
	dbg("traceLocation %q (%d)", startMap, startPos)
	if startMap == "location" {
		return startPos, "location"
	}

	nextPos := a.getNextPosition(startMap, startPos)

	return a.traceLocation(a.maps[startMap].next, nextPos)
}

func solve1(a almanac) int {
	minLoc := math.MaxInt
	for _, s := range a.seeds {
		loc, _ := a.traceLocation("seed", s)
		if loc < minLoc {
			minLoc = loc
		}

	}

	return minLoc
}

func solve2(a almanac) int {
	minLoc := math.MaxInt

	for i := 0; i < len(a.seeds); i += 2 {
		dbg("Seeds %d/%d -> %d", a.seeds[i], len(a.seeds)/2, a.seeds[i]+a.seeds[i+1]-1)
		log.Printf("Seeds %d/%d -> %d\n", a.seeds[i], len(a.seeds)/2, a.seeds[i]+a.seeds[i+1]-1)
		for seed, n, dots := a.seeds[i], 1, 0; seed < a.seeds[i]+a.seeds[i+1]; seed++ {
			dbg("Seed %d", seed)
			loc, _ := a.traceLocation("seed", seed)
			dbg("Found: %d", loc)

			if loc < minLoc {
				minLoc = loc
			}
			if n%10000 == 0 {
				fmt.Print(".")
				if dots == 100 {
					fmt.Printf(" (%d)\n", n)
					dots = 0
				}
				dots++
			}
			n++
		}
	}

	return minLoc
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
	a := parseAlmanac(lines)
	part1 = solve1(a)
	part2 = solve2(a)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
