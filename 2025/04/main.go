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
	valid := 0
	for y := range s {
		for x := range s[0] {
			p := Point{x, y}
			if ch, _ := GetChInPoint(s, p); ch != '@' {
				continue
			}

			adj := p.Adjacent(true)
			curvalid := 0
			dbg("CHECKING %+v (%c) (valid: %v)", p, s[p.Y][p.X], valid)
			for _, point := range adj {
				dbg("  - ADJ: %+v (curvalid: %v)", point, curvalid)
				ch, valid := GetChInPoint(s, point)
				dbg("      -> ch: %c, valid: %v", ch, valid)
				if valid && ch == '@' {
					curvalid++
					dbg("      -> CURVALID++ -> %v", curvalid)
				}

			}
			dbg("CURVALID == %v", curvalid)
			if curvalid < 4 {
				dbg("VALID POINT: %v", p)
				valid++
			}
		}
	}

	return valid
}

func solve2(s []string) int {
	res := 0

	return res
}
