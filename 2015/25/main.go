package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
	// . "github.com/agimenez/adventofcode/utils"
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

func GetNthCode(pos int) int {
	prev := 20151125
	for ; pos > 1; pos-- {
		prev = prev * 252533
		prev = prev % 33554393
	}

	return prev

}

func GetPosFromCoords(r, c int) int {
	// The initial column value (row = 1) is just the sum of
	// the sequence of natural numbers up to that column
	colValue := (c * (c + 1)) / 2

	// Once we have the value for column == 1, the succession for
	// the row values is c, c+1, c+2, .... with a max of "r-1" terms
	// for example, in the case of r=5, c=2, we know the first row of c=2
	// is 3 (from above), then we will have the following:
	// 3+(2+3+4+5) = 17
	// So we have the sum formula (n(2a+n-1)/2, where:
	// n is the number of terms (r-1)
	// a is the first term (the column number)
	rowSeqSum := ((r - 1) * (2*c + (r - 1) - 1)) / 2

	return colValue + rowSeqSum
}

func solve1(s []string) int {
	res := 0
	row := 3010
	col := 3019

	res = GetNthCode(GetPosFromCoords(row, col))

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
