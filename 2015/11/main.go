package main

import (
	"bytes"
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

func solve(lines []string) (string, string, time.Duration, time.Duration) {
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

func increasingStraight(s string) bool {
	for i := range s[:len(s)-2] {
		if s[i] == s[i+1]-1 && s[i] == s[i+2]-2 {
			return true
		}
	}

	return false
}

func hasOverlappingPairs(s string) bool {

	pairs := 0
	for i := 0; i < len(s)-1; i++ {
		dbg("  >> s[%d]: %c, s[i+1]: %c", i, s[i], s[i+1])
		if s[i] == s[i+1] {
			dbg("  >> YES (pairs: %d", pairs)
			pairs++
			i++
		}

		if pairs == 2 {
			return true
		}
	}

	return false
}

func Valid(s string) bool {
	// Include increasing straight of at least 3 letters
	if !increasingStraight(s) {
		return false
	}

	// Does not contain i, o, l
	if strings.IndexAny(s, "iol") != -1 {
		return false
	}

	// Must ocntain two different, non-overlapping pairs
	if !hasOverlappingPairs(s) {
		return false
	}

	return true
}

// We know a rule for sure, which is does not contain i, o, or l
// The fist candidate will be the first occurrence of either or them, incremented, and the rest changed to "a"
// Example:
//
//	ghijklmn -> Everything from the first "i" is useless, so increment the first "i" to "j", and reset the rest:
//	ghjaaaaa
func FirstCandidate(s string) string {

	firstInvalid := strings.IndexAny(s, "iol")
	if firstInvalid == -1 {
		return s
	}

	var b bytes.Buffer
	b.WriteString(s[:firstInvalid])
	b.WriteByte(s[firstInvalid] + 1)
	b.WriteString(strings.Repeat("a", len(s)-(firstInvalid+1)))

	return b.String()

}

func NextPass(s string) string {
	dbg("NextPass %q", s)
	next := FirstCandidate(s)
	if len(next) != len(s) {
		panic(fmt.Sprintf("Len mismatch: %q (%d) -> %q (%d)", next, len(next), s, len(s)))
	}
	dbg(" >> FIRST CANDIDATE: %v", next)
	for !Valid(next) {
		next = Increment(next)
		dbg(" >> NEXT: %v", next)
	}
	dbg(" >> FINAL: %v", next)

	return next
}

func Increment(s string) string {
	runes := []rune(s)
	done := false
	for i := len(runes) - 1; i >= 0 && !done; i-- {
		if runes[i] == 'z' {
			runes[i] = 'a'
			continue
		}

		runes[i]++
		// Optimization? Do not generate any passwords with i, o, or l
		if runes[i] == 'i' || runes[i] == 'o' || runes[i] == 'l' {
			runes[i]++
		}

		done = true
	}

	return string(runes)
}

func solve1(s []string) string {

	res := NextPass(s[0])

	return res
}

func solve2(s []string) string {
	res := s[0]

	return res
}
