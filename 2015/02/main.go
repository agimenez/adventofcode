package main

import (
	"flag"
	"io"
	"log"
	"math"
	"os"
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

	part1, part2 := 0, 0
	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

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

func sqFeet(box string) int {
	res := 0
	parts := strings.Split(box, "x")

	l := utils.ToInt(parts[0])
	w := utils.ToInt(parts[1])
	h := utils.ToInt(parts[2])

	minSide := math.MaxInt
	areas := []int{
		l * w,
		w * h,
		h * l,
	}

	dbg("box: %v, areas: %v", box, areas)
	for _, a := range areas {
		res += 2 * a
		if a < minSide {
			minSide = a
		}
	}

	res += minSide

	return res
}

func solve1(s []string) int {
	res := 0
	for _, box := range s {
		res += sqFeet(box)
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
