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

func solve1(s []string) int {
	res := 0
	cur := P0
	dir := P0.Up()
	for _, instr := range strings.Split(s[0], ", ") {
		dbg("cur: %v, instr: %v", cur, instr)
		if instr[0] == 'R' {
			dir = dir.Rotate90CW()
		} else {
			dir = dir.Rotate90CCW()
		}

		for range ToInt(string(instr[1:])) {
			cur = cur.Sum(dir)
		}
	}

	res = cur.ManhattanDistance(P0)

	return res
}

func solve2(s []string) int {
	res := 0

	cur := P0
	dir := P0.Up()
	visited := map[Point]bool{P0: true}
	minP, maxP := Point{0, 0}, Point{0, 0}
	for _, instr := range strings.Split(s[0], ", ") {
		dbg("cur: %v, instr: %v", cur, instr)
		if instr[0] == 'R' {
			dir = dir.Rotate90CW()
		} else {
			dir = dir.Rotate90CCW()
		}

		for range ToInt(string(instr[1:])) {
			cur = cur.Sum(dir)
			minP = minP.Min(cur)
			maxP = maxP.Max(cur)

			if !visited[cur] {
				visited[cur] = true
			} else {
				return cur.ManhattanDistance(P0)

			}
			dbg("minmax: %v - %v", minP, maxP)
			printVisited(visited, minP, maxP)
			time.Sleep(20 * time.Millisecond)
		}

	}

	res = cur.ManhattanDistance(P0)

	return res
}

func printVisited(visited map[Point]bool, minP, maxP Point) {
	println("== MAP ==")
	for y := minP.Y; y <= maxP.Y; y++ {
		for x := minP.X; x <= maxP.X; x++ {
			p := Point{x, y}
			if !visited[p] {
				print(".")
			} else {
				print("#")
			}
		}
		println("")
	}
	println("-------")
}
