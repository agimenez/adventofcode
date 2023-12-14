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

func mirrorCol(s []string, startRow int, col int) bool {
	for i := startRow; i < len(s); i++ {
	}
}
func solve1(in []string) int {
	res := 0

	// First, check vertical reflection
	for col := 0; col < len(in[0]); col++ {
		dbg("Checking column %d", col)
		mirror := true
		for _, row := range in {
			mirror = mirror && mirrorLine(row[col:])
			if !mirror {
				break
			}
		}

		if mirror {
			res += (len(in[0])-col)/2 + 1
			break
		}
	}

	// Check horizontal reflection
	for row := 0; row < len(in); row++ {
		dbg("Checking row %d", row)
		mirror := true
		for col := 0; col < len(in[0]); c++ {
			mirror = mirror && mirrorCol(in, row, col)
			if !mirror {
				break
			}
		}

		if mirror {
			res += 100 * ((len(in)-row)/2 + 1)
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

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
