package main

import (
	"flag"
	"fmt"
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

func dbg(f string, v ...interface{}) {
	if debug {
		fmt.Printf(f+"\n", v...)
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

func ConwayStep(g Grid) Grid {
	next := NewGrid(g.Width(), g.Height())

	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			p := NewPoint(x, y)
			r := g.GetRune(p)
			next.SetRune(p, r)

			isOn := r == '#'
			countOn := 0
			for v := range g.AdjacentPoints(p, true) {
				if g.GetRune(v) == '#' {
					countOn++
				}
			}

			if isOn && !(countOn == 2 || countOn == 3) {
				next.SetRune(p, '.')
			} else if !isOn && countOn == 3 {
				next.SetRune(p, '#')
			}
		}
	}

	return next
}

func SwitchCornersOn(g Grid) Grid {
	w := g.Width() - 1
	h := g.Height() - 1
	for _, p := range []Point{NewPoint(0, 0), NewPoint(0, h), NewPoint(w, 0), NewPoint(w, h)} {
		g.SetRune(p, '#')
	}

	return g
}

func ConwaysGame(g Grid, steps int, stuckCorners bool) Grid {

	for range steps {
		g = ConwayStep(g)
		if stuckCorners {
			g = SwitchCornersOn(g)
		}

		dbg("%v", g)
	}

	return g
}

func solve1(s []string) int {
	res := 0

	g := NewGridFromStr(s)
	dbg("%v", g)

	// Kind of cheat?
	if len(s) < 10 {
		g = ConwaysGame(g, 4, false)
	} else {
		g = ConwaysGame(g, 100, false)
	}

	g.MapFunc(func(r rune) {
		if r == '#' {
			res++
		}
	})

	return res
}

func solve2(s []string) int {
	res := 0

	g := SwitchCornersOn(NewGridFromStr(s))
	dbg("%v", g)

	// Kind of cheat?
	if len(s) < 10 {
		g = ConwaysGame(g, 5, true)
	} else {
		g = ConwaysGame(g, 100, true)
	}

	g.MapFunc(func(r rune) {
		if r == '#' {
			res++
		}
	})

	return res
}
