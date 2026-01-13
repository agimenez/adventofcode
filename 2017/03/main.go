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

type Spiral map[int]Point

func genSpiral(pos int) (Spiral, int) {
	s := Spiral{}

	curPoint := P0
	dir := P0.Down()
	seen := map[Point]bool{}

	neighbourSum := map[Point]int{P0: 1}
	max := 0

	// The logic of this is:
	// - try to turn left (always)
	// - If the left point is already seen (we're in a loop) -> keep going
	// - If the point of turning left is not seen, then we need to actually turn
	for curPos := 1; curPos <= pos; curPos++ {
		seen[curPoint] = true
		s[curPos] = curPoint

		neighbourSum[curPoint] = ReduceCollect(curPoint.Adjacent(true), neighbourSum[curPoint], func(acc int, p Point) int {
			return acc + neighbourSum[p]
		})

		if neighbourSum[curPoint] > pos && max == 0 {
			max = neighbourSum[curPoint]
		}

		if turnLeft := curPoint.Sum(dir.Rotate90CCW()); seen[turnLeft] == false {
			curPoint = turnLeft
			dir = dir.Rotate90CCW()
		} else {
			curPoint = curPoint.Sum(dir)
		}
	}

	return s, max
}

// This part is Ulam's spiral, and could be solved mathemathically, but
// who wants math when we can simulate!!!
func solve1(s []string) int {
	res := 0

	pos := ToInt(s[0])
	sp, _ := genSpiral(pos)
	res = sp[pos].ManhattanDistance(P0)
	dbg("%v", sp)

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")
	pos := ToInt(s[0])
	_, res = genSpiral(pos)

	return res
}
