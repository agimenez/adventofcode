package main

import (
	"flag"
	"io"
	"log"
	"os"
	"slices"
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

func sqFeet(box string) (int, int) {
	wrap, ribbon := 0, 1
	parts := strings.Split(box, "x")

	l := utils.ToInt(parts[0])
	w := utils.ToInt(parts[1])
	h := utils.ToInt(parts[2])
	sides := []int{
		l,
		w,
		h,
	}
	slices.Sort(sides)

	areas := []int{
		l * w,
		w * h,
		h * l,
	}
	slices.Sort(areas)

	dbg("box: %v, areas: %v", box, areas)
	for i := range areas {
		wrap += 2 * areas[i]
		ribbon *= sides[i]
	}

	wrap += areas[0]
	ribbon += 2*sides[0]   + 2*sides[1]

	return wrap, ribbon
}

func solve1(s []string) int {
	res := 0
	for _, box := range s {
		r, _ := sqFeet(box)
		res += r
	}

	return res
}

func solve2(s []string) int {
	res := 0

	for _, box := range s {
		_, r := sqFeet(box)
		res += r
	}

	return res
}
