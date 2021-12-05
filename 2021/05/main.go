package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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

func parsePoint(s string) Point {
	p := Point{}

	coords := strings.Split(s, ",")
	p.X, _ = strconv.Atoi(coords[0])
	p.Y, _ = strconv.Atoi(coords[1])

	return p
}

type vents map[Point]int

func (v vents) addVent(start Point, end Point) {
	if start.X == end.X || start.Y == end.Y {
		// part 1: consider only vertical lines
		xStart := Min(start.X, end.X)
		xEnd := Max(start.X, end.X)

		yStart := Min(start.Y, end.Y)
		yEnd := Max(start.Y, end.Y)

		for x := xStart; x <= xEnd; x++ {
			for y := yStart; y <= yEnd; y++ {
				v[Point{x, y}]++
			}
		}

	} else {
		// part2: consider diagonal lines too
		dx := (end.X - start.X) / Abs(end.X-start.X)
		dy := (end.Y - start.Y) / Abs(end.Y-start.Y)
		dbg("%s -> %s, dx: %d, dy: %d", start, end, dx, dy)
		for x, y := start.X, start.Y; x != end.X+dx; x, y = x+dx, y+dy {
			dbg(" -> (%d,%d)", x, y)
			v[Point{x, y}]++
		}
	}
}

func (v vents) addLine(l string) {
	parts := strings.Split(l, " -> ")
	start := parsePoint(parts[0])
	end := parsePoint(parts[1])

	v.addVent(start, end)
}

func (v vents) overlapping(min int) int {
	total := 0
	for _, count := range v {
		if count >= min {
			total++
		}
	}

	return total
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
	vents := vents{}
	for _, l := range lines {
		vents.addLine(l)
		dbg("vents: %+v", vents)
	}

	part1 = vents.overlapping(2)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
