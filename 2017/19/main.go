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

func getStart(g Grid) Point {
	p := P0
	for x := 0; x < g.Width(); x++ {
		p.X = x
		if g.GetRune(p) == '|' {
			return p
		}
	}

	return p
}

func isChar(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func followPath(g Grid, p Point) (string, int) {
	var sb strings.Builder

	dir := P0.Down()
	steps := 0
	for {
		ch := g.GetRune(p)
		dbg("%v (%v) -> '%c'", p, dir, ch)
		if ch == ' ' || g.OutOfBounds(p) {
			break
		}

		if isChar(ch) {
			sb.WriteRune(ch)
		}

		if ch == '+' {
			dbg("  >> DIR CHANGE")
			var newDir Point
			for _, newDir = range []Point{dir.Rotate90CCW(), dir.Rotate90CW()} {
				next := p.Sum(newDir)
				if g.GetRune(next) != ' ' && !g.OutOfBounds(next) {
					dir = newDir
					break
				}
			}
		}

		p = p.Sum(dir)
		steps++
	}

	return sb.String(), steps
}

func solve1(s []string) int {
	res := 0

	g := NewGridFromStr(s)
	start := getStart(g)
	dbg("Got start: %v", start)

	p1, res := followPath(g, start)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", res)

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	return res
}
