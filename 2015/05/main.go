package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/agimenez/adventofcode/utils"
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

func isNice(word string) bool {
	// At least three vowels
	vowels := 0
	for _, r := range word {
		if strings.ContainsRune("aeiou", r) {
			vowels++
		}
	}
	if vowels < 3 {
		return false
	}

	// contains a letter twice in a row
	twice := false
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			twice = true
			break
		}
	}
	if !twice {
		return false
	}

	// does not contain ab, cd, pq, xy. Could merge this into the previous
	// rule, but just keep'em separated
	for i := 1; i < len(word); i++ {
		sub := word[i-1 : i+1]
		if sub == "ab" || sub == "cd" || sub == "pq" || sub == "xy" {
			return false
		}

	}

	return true
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
	for _, w := range s {
		if isNice(w) {
			res++
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0
	// Well... erm...
	// So RE2 doesn't support backreferences, so I can't be even bothered...
	cmd := `grep -E '([a-zA-Z]{2}).*\1' input.txt | grep -E '(\w).\1' | wc -l`
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		panic("exec")
	}

	dbg("Out: %q", string(out))
	res = utils.ToInt(string(out[:len(out)-1]))
	dbg("res: %v", res)

	return res
}
