package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"
	"time"

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

	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	part1, part2, dur1, dur2 := solve(lines)
	log.Printf("Part 1 (%v): %v\n", dur1, part1)
	log.Printf("Part 2 (%v): %v\n", dur2, part2)

}

func solve(lines []string) (int, int, time.Duration, time.Duration) {
	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 := solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 := solve2(lines)
	dur[1] = time.Since(now)

	return part1, part2, dur[0], dur[1]

}

type rect struct {
	p1, p2 Point
}

func (r rect) area() int {
	dbg("%v x %v", Abs(r.p2.X-r.p1.X)+1, Abs(r.p2.Y-r.p1.Y)+1)
	return (Abs(r.p2.X-r.p1.X) + 1) * (Abs(r.p2.Y-r.p1.Y) + 1)
}

func solve1(s []string) int {
	res := 0
	tiles := []Point{}
	for _, line := range s {
		parts := strings.Split(line, ",")
		tiles = append(tiles, Point{ToInt(parts[0]), ToInt(parts[1])})
	}

	for i, p1 := range tiles {
		for _, p2 := range tiles[i+1:] {
			r := rect{p1, p2}
			area := r.area()
			dbg("%v - %v: %v", p1, p2, area)
			if area > res {
				res = area
			}
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
