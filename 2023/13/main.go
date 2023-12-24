package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

func mirrorLine(s string) bool {
	dbg(" -> mirrorLine %s", s)
	for i := 0; i < len(s)/2; i++ {
		dbg("    -> %c == %c", s[i], s[len(s)-i-1])
		if s[i] != s[len(s)-i-1] {
			dbg("    -> NAH!")
			return false
		}
	}

	dbg("    -> YUP!")
	return true
}

func printCol(s []string, row int, col int) {
	if !debug {
		return
	}
	dbg("Printing row=%d, col=%d", row, col)
	for i := row; i < len(s); i++ {
		log.Printf("%c", s[i][col])

	}
}

func mirrorCol(s []string, startRow int, col int) bool {
	dbg(" -> mirrorCol")
	printCol(s, startRow, col)
	for i := startRow; i <= len(s)/2; i++ {
		dbg("    -> (%d) %c == (%d) %c", i, s[i][col], len(s)-i-1, s[len(s)-i-1][col])
		if s[i][col] != s[len(s)-i-1][col] {
			dbg("    -> NAH!")
			return false
		}
	}

	dbg("    -> YUP!")

	return true
}

func solve1(in []string) int {
	res := 0

	dbg("Solving for pattern:\n%s", strings.Join(in, "\n"))
	// First, check vertical reflection
	for col := 0; col < len(in[0])-1; col++ {
		dbg("Checking column %d", col)
		mirror := true
		for _, row := range in {
			mirror = mirror && mirrorLine(row[col:])
			if !mirror {
				break
			}
		}

		if mirror {
			dbg("Found mirror in column %d, returning %d", col, (len(in[0])-col)/2+1)
			return (len(in[0])-col)/2 + 1
		}
	}

	// Check horizontal reflection
	for row := 0; row < len(in); row++ {
		dbg("Checking row %d", row)
		mirror := true
		for col := 0; col < len(in[0]); col++ {
			mirror = mirror && mirrorCol(in, row, col)
			if !mirror {
				break
			}
		}

		if mirror {
			dbg("Found mirror in row %d", row)
			return 100 * (row + 1)
			break
		}
	}

	return res
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
	//dbg("lines: %#v", lines)
	pattern := []string{}
	for _, l := range lines {
		if l == "" {
			part1 += solve1(pattern)
			pattern = []string{}
			continue
		}

		pattern = append(pattern, l)

	}
	part1 += solve1(pattern)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
