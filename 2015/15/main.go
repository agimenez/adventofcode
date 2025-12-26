package main

import (
	"flag"
	"fmt"
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

type Ingredient struct {
	capacity, durability, flavor, texture, calories int
}

func solve1(s []string) int {
	res := 0

	ings := []Ingredient{}
	for _, line := range s {
		parts := strings.Split(line, ": ")

		properties := strings.Split(parts[1], " ")
		i := Ingredient{
			capacity:   ToInt(properties[1][:len(properties[1])-1]),
			durability: ToInt(properties[3][:len(properties[3])-1]),
			flavor:     ToInt(properties[5][:len(properties[5])-1]),
			texture:    ToInt(properties[7][:len(properties[7])-1]),
			calories:   ToInt(properties[9]),
		}

		ings = append(ings, i)

	}

	// A bit hardcoded, can't bother to make it more general for more ingredients
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			for k := 0; k <= 100-i-j; k++ {
				l := 100 - i - j - k
				dbg("Trying: %v, %v, %v, %v", i, j, k, l)

				counters := []int{i, j, k, l}
				var capacity, durability, flavor, texture int

				for n, ing := range ings {
					capacity += counters[n] * ing.capacity
					durability += counters[n] * ing.durability
					flavor += counters[n] * ing.flavor
					texture += counters[n] * ing.texture
				}

				if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 {
					continue
				}
				score := capacity * durability * flavor * texture

				res = Max(res, score)
			}
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
