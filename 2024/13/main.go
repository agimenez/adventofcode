package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)
	buttons := map[string]Point{}
	var prize Point

	buttonsRegex := regexp.MustCompile(`Button ([AB]): X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	for _, l := range lines {
		if l == "" {
			continue
		}

		res := buttonsRegex.FindStringSubmatch(l)
		if res != nil {
			buttons[res[1]] = Point{
				X: ToInt(res[2]),
				Y: ToInt(res[3]),
			}

			continue
		}

		res = prizeRegex.FindStringSubmatch(l)
		if res != nil {
			prize.X = ToInt(res[1])
			prize.Y = ToInt(res[2])

			part1 += solve1(prize, buttons)

			prize.X += 10000000000000
			prize.Y += 10000000000000
			part2 += solve1(prize, buttons)
		}

	}

	var dur [2]time.Duration

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func solve1(prize Point, buttons map[string]Point) int {
	res := 0
	dbg("Solve: %v, %v", prize, buttons)
	a := buttons["A"]
	b := buttons["B"]

	det := (a.X*b.Y - a.Y*b.X)

	na := (prize.X*b.Y - prize.Y*b.X) / det
	nb := (prize.Y*a.X - a.Y*prize.X) / det

	if na*a.X+nb*b.X == prize.X && na*a.Y+nb*b.Y == prize.Y {
		res = 3*na + nb
	}

	return res
}

func solve2(s string) int {
	res := 0

	return res
}
