package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"slices"
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

type Reindeer struct {
	speed    int
	flyTime  int
	restTime int

	flying   bool
	timeLeft int

	distance int
}

func (r *Reindeer) Step() {
	if r.timeLeft == 0 {
		r.flying = !r.flying
		if r.flying {
			r.timeLeft = r.flyTime - 1
			r.distance += r.speed
		} else {
			r.timeLeft = r.restTime - 1
		}

		return
	}

	if r.timeLeft > 0 {
		r.timeLeft--
	}

	if r.flying {
		r.distance += r.speed
	}
}

func (r Reindeer) String() string {
	return fmt.Sprintf("Flying: %v, dist: %v, timeLeft: %v", r.flying, r.distance, r.timeLeft)
}

// I know, this is a math problem and simulate is stupid, but let me be!
func SimulateRace(m map[string]Reindeer, t int) []int {
	reindeers := slices.Collect(maps.Values(m))

	for i := range t {
		dbg("TIME: %v", i+1)
		for r := range reindeers {
			reindeers[r].Step()
			dbg("%v", reindeers[r])
		}

	}

	dist := slices.Collect(func(yield func(int) bool) {
		for _, r := range reindeers {
			if !yield(r.distance) {
				return
			}
		}
	})

	return dist
}

func solve1(s []string) int {
	res := 0

	race := map[string]Reindeer{}

	for _, line := range s {
		parts := strings.Fields(line)
		r := Reindeer{
			speed:    ToInt(parts[3]),
			flyTime:  ToInt(parts[6]),
			restTime: ToInt(parts[len(parts)-2]),
		}

		race[parts[0]] = r
	}

	dists := SimulateRace(race, 2503)

	slices.Sort(dists)
	res = dists[len(dists)-1]

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
