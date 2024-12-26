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

type grid map[Point]bool

func (g grid) toggleLights(start Point, end Point) {
	for y := start.Y; y <= end.Y; y++ {
		for x := start.X; x <= end.X; x++ {
			p := Point{y, x}
			g[p] = !g[p]
		}
	}
}

func (g grid) turnLights(s string, start Point, end Point) {
	val := false
	if s == "on" {
		val = true
	}

	for y := start.Y; y <= end.Y; y++ {
		for x := start.X; x <= end.X; x++ {
			g[Point{y, x}] = val
		}
	}

}

func solve1(s []string) int {
	res := 0
	g := make(grid, 1000*1000)
	for _, l := range s {
		parts := strings.Split(l, " ")

		if parts[0] == "turn" {
			s := strings.Split(parts[2], ",")
			start := Point{
				Y: ToInt(s[0]),
				X: ToInt(s[1]),
			}

			e := strings.Split(parts[4], ",")
			end := Point{
				Y: ToInt(e[0]),
				X: ToInt(e[1]),
			}

			g.turnLights(parts[1], start, end)
		} else if parts[0] == "toggle" {
			s := strings.Split(parts[1], ",")
			start := Point{
				Y: ToInt(s[0]),
				X: ToInt(s[1]),
			}

			e := strings.Split(parts[3], ",")
			end := Point{
				Y: ToInt(e[0]),
				X: ToInt(e[1]),
			}
			g.toggleLights(start, end)

		}

	}

	for _, on := range g {
		if on {
			res++
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
