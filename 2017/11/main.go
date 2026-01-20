package main

import (
	"flag"
	"fmt"
	. "github.com/agimenez/adventofcode/utils"
	"io"
	"log"
	"os"
	"strings"
	"time"
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

// redblobgames.com -> hexagonal grids
type Hex struct {
	q, r, s int
}

var H0 = Hex{0, 0, 0}

var hexDir = map[string]Hex{
	"n":  {0, -1, 1},
	"ne": {1, -1, 0},
	"se": {1, 0, -1},
	"s":  {0, 1, -1},
	"sw": {-1, 1, 0},
	"nw": {-1, 0, 1},
}

func (h Hex) Add(h2 Hex) Hex {
	return Hex{
		q: h.q + h2.q,
		r: h.r + h2.r,
		s: h.s + h2.s,
	}
}

func (h Hex) Move(dir string) Hex {
	return h.Add(hexDir[dir])
}

func (h Hex) Sub(h2 Hex) Hex {
	return Hex{
		q: h.q - h2.q,
		r: h.r - h2.r,
		s: h.s - h2.s,
	}
}

func (h Hex) ManhattanDist(dest Hex) int {
	sub := h.Sub(dest)

	return (Abs(sub.q) + Abs(sub.r) + Abs(sub.s)) / 2

}

func Walk(path string) Hex {
	cur := Hex{}
	for _, dir := range strings.Split(path, ",") {
		cur = cur.Move(dir)
	}

	return cur
}

func CountSteps(path string) int {
	return Walk(path).ManhattanDist(H0)
}

func FurtherDistance(path string) int {
	res := 0
	cur := Hex{}
	for _, dir := range strings.Split(path, ",") {
		cur = cur.Move(dir)

		res = Max(res, cur.ManhattanDist(H0))
	}

	return res
}

func solve1(s []string) int {
	res := 0

	for _, path := range s {
		res += CountSteps(path)
		dbg("%v: %v", path, res)
	}

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	for _, path := range s {
		res = FurtherDistance(path)
	}

	return res
}
