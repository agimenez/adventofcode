package main

import (
	"flag"
	"io"
	"log"
	"maps"
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

	part1, part2 := 0, 0
	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	dbg("lines: %#v", lines)

	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 = solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 = solve2(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func solve1(s []string) int {
	res := 0
	cur := P0
	houses := map[Point]int{cur: 1}
	for _, dir := range s[0] {
		switch dir {
		case '>':
			cur = cur.Right()

		case 'v':
			cur = cur.Down()

		case '<':
			cur = cur.Left()

		case '^':
			cur = cur.Up()
		}
		houses[cur]++
	}

	res = len(houses)

	return res
}

func solve2(s []string) int {
	res := 0

	for _, l := range s {
		if l == "" {
			continue
		}

		houses := map[bool]map[Point]int{
			true:  {P0: 1},
			false: {P0: 1},
		}
		pos := map[bool]Point{
			true:  P0,
			false: P0,
		}
		turnSanta := true
		for _, dir := range l {
			cur := pos[turnSanta]
			switch dir {
			case '>':
				cur = cur.Right()

			case 'v':
				cur = cur.Down()

			case '<':
				cur = cur.Left()

			case '^':
				cur = cur.Up()
			}
			houses[turnSanta][cur]++
			pos[turnSanta] = cur
			turnSanta = !turnSanta
		}

		union := maps.Clone(houses[true])
		for p := range houses[false] {
			union[p]++
		}
		res = len(union)
	}
	return res
}
