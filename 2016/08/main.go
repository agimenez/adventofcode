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

func printScreen(s [6][50]bool) {
	for y := 0; y < 6; y++ {
		for x := 0; x < 50; x++ {
			if s[y][x] {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
	println()
}

func rotateCol(s [6][50]bool, col int, offset int) [6][50]bool {
	newCol := [6]bool{}
	for row := range 6 {
		oldRow := (row - offset)
		if oldRow < 0 {
			oldRow = 6 + oldRow
		}
		newCol[row] = s[oldRow][col]
	}
	dbg("New Col: %v", newCol)

	for row := range 6 {
		s[row][col] = newCol[row]
	}

	printScreen(s)

	return s
}

func rotateRow(s [6][50]bool, row int, offset int) [6][50]bool {
	newRow := [50]bool{}
	for col := range 50 {
		oldCol := (col - offset)
		if oldCol < 0 {
			oldCol = 50 + oldCol
		}
		newRow[col] = s[row][oldCol]
	}
	dbg("New Row: %v", newRow)

	for col := range 50 {
		s[row][col] = newRow[col]
	}

	printScreen(s)

	return s

}

func solve1(s []string) int {
	res := 0
	screen := [6][50]bool{}

	for _, l := range s {
		dbg("Instruction: %v", l)
		parts := strings.Split(l, " ")
		if parts[0] == "rect" {
			dims := strings.Split(parts[1], "x")
			maxX := ToInt(dims[0])
			maxY := ToInt(dims[1])

			for y := 0; y < maxY; y++ {
				for x := 0; x < maxX; x++ {
					screen[y][x] = true
				}
			}
		} else {
			if parts[1] == "column" {
				var col int
				fmt.Sscanf(parts[2], "x=%d", &col)

				offset := ToInt(parts[4])
				dbg("Rotate column %v by %v", col, offset)
				screen = rotateCol(screen, col, offset)

			} else {
				var row int
				fmt.Sscanf(parts[2], "y=%d", &row)

				offset := ToInt(parts[4])
				dbg("Rotate column %v by %v", row, offset)
				screen = rotateRow(screen, row, offset)
			}
		}
		printScreen(screen)
	}

	for y := range screen {
		for x := range screen[y] {
			if screen[y][x] {
				res++
			}
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
