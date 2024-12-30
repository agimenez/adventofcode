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

func decompress(l string) string {
	var sb strings.Builder

	remain := l[:]
	for {
		maxdbg := 50
		if len(remain) < 50 {
			maxdbg = len(remain)
		}

		dbg("Remaining: %q[...]", remain[:maxdbg])
		openParen := strings.Index(remain, "(")

		// No more open parens
		if openParen == -1 {
			sb.WriteString(remain)
			break
		}

		// Copy until now
		dbg("copying %d - %d [%q]", 0, openParen, remain[:openParen])
		sb.WriteString(remain[:openParen])
		remain = remain[openParen+1:]
		dbg("NewRemain: %q", remain)

		closeParen := strings.Index(remain, ")")
		var lookahead int
		var reps int
		fmt.Sscanf(remain[:closeParen], "%dx%d", &lookahead, &reps)
		dbg(" -> MARK: %q: %v x %v", remain[0:closeParen], lookahead, reps)

		startCopy := closeParen + 1
		endCopy := closeParen + 1 + lookahead
		dbg(" -> REP [%v:%v] %q", startCopy, endCopy, remain[startCopy:endCopy])
		for ; reps > 0; reps-- {
			sb.WriteString(remain[startCopy:endCopy])
		}

		dbg("OUT: %q", sb.String())

		remain = remain[endCopy:]
		// time.Sleep(500 * time.Millisecond)
	}

	return sb.String()
}

func solve1(s []string) int {
	res := 0
	for _, l := range s {
		d := decompress(l)

		res += len(d)
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
