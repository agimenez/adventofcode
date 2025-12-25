package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

func unescapeString(s string) string {
	var out string
	var err error

	if out, err = strconv.Unquote(s); err != nil {
		fmt.Printf("ERROR processing %s", s)
	}

	return out
}

func solve1(s []string) int {
	res := 0

	for _, str := range s {
		out := unescapeString(str)
		dbg("str: %q (%v)", str, len(str))
		dbg("out: %q (%v)", out, len(out))
		dbg("")

		res = res + len(str) - len(out)
	}

	return res
}

func solve2(s []string) int {
	res := 0

	for _, str := range s {
		out := strconv.Quote(str)
		dbg("str: %q (%v)", str, len(str))
		dbg("out: %q (%v)", out, len(out))
		dbg("")

		res = res + len(out) - len(str)
	}

	return res
}
