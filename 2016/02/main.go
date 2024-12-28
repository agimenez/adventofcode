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

func solve(lines []string) (int, string, time.Duration, time.Duration) {
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

var grid = []string{
	"123",
	"456",
	"789",
}

func solve1(s []string) int {
	res := 0
	cur := Point{1, 1}
	dirs := map[rune]Point{
		'U': P0.Up(),
		'R': P0.Right(),
		'D': P0.Down(),
		'L': P0.Left(),
	}
	curChar, _ := GetChInPoint(s, cur)
	chars := []rune{}
	for _, l := range s {
		for _, c := range l {
			nextPos := cur.Sum(dirs[c])
			nextChar, ok := GetChInPoint(grid, nextPos)
			if ok {
				cur = nextPos
				curChar = nextChar
			}
		}
		dbg("Next: %c", curChar)
		chars = append(chars, rune(curChar))
	}
	dbg("Chars: %v", string(chars))
	res = ToInt(string(chars))

	return res
}

var grid2 = []string{
	"  1",
	" 234",
	"56789",
	" ABC",
	"  D",
}

func solve2(s []string) string {
	res := ""

	cur := Point{0, 2}
	dirs := map[rune]Point{
		'U': P0.Up(),
		'R': P0.Right(),
		'D': P0.Down(),
		'L': P0.Left(),
	}
	curChar, _ := GetChInPoint(grid2, cur)
	chars := []rune{}
	for _, l := range s {
		for _, c := range l {
			nextPos := cur.Sum(dirs[c])
			nextChar, ok := GetChInPoint(grid2, nextPos)
			dbg("DIR: %c, CUR: '%c', CURP: %v", c, curChar, cur)
			dbg(" - NEXTP: %v, NEXTC: '%c'", nextPos, nextChar)
			if ok && nextChar != ' ' {
				cur = nextPos
				curChar = nextChar
				dbg(" - SWAP!")
			}
		}
		dbg("Next: %c\n\n", curChar)
		chars = append(chars, rune(curChar))
	}
	dbg("Chars: %v", string(chars))
	res = string(chars)

	return res
}
