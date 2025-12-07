package main

import (
	"flag"
	"io"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	. "github.com/agimenez/adventofcode/utils"
)

var (
	debug bool
	loops int
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
	flag.IntVar(&loops, "loops", 10, "Loops")
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

type Box struct {
	x, y, z int
}

type distPair struct {
	a, b Box
	dist int
}

type Circuit map[Box]bool

// Save me some math and floats. We just want to know distance magnitudes
// (i.e. which ones are closer, etc), so we don't really care about Sqrt stuff
func EuclideanDistance(box1, box2 Box) int {
	dx := box2.x - box1.x
	dy := box2.y - box1.y
	dz := box2.z - box1.z

	return dx*dx + dy*dy + dz*dz
}

func getDistPair(a, b Box) distPair {
	return distPair{
		a:    a,
		b:    b,
		dist: EuclideanDistance(a, b),
	}
}

func getDistancePairs(boxes []Box) []distPair {
	distPairs := []distPair{}

	for i, b1 := range boxes {
		for _, b2 := range boxes[i+1:] {
			distPairs = append(distPairs, getDistPair(b1, b2))
		}
	}

	return distPairs
}

func solve1(s []string) int {
	res := 0
	boxes := make([]Box, len(s))
	circuits := make([]Circuit, len(s))
	box2circuit := map[Box]int{}

	for i, coord := range s {
		parts := strings.Split(coord, ",")
		b := Box{
			x: ToInt(parts[0]),
			y: ToInt(parts[1]),
			z: ToInt(parts[2]),
		}
		boxes[i] = b
		circuits[i] = map[Box]bool{b: true}
		box2circuit[b] = i
	}

	//dbg("Boxes: %v", boxes)

	pairs := getDistancePairs(boxes)
	slices.SortFunc(pairs, func(a, b distPair) int {
		return a.dist - b.dist
	})
	//dbg("Sorted pairs: %v", pairs)
	for range loops {
		current := pairs[0]
		pairs = pairs[1:]
		dbg("Current pair: %+v", current)

		// These are the closest boxes. I'm not sure if there can be a case
		// where there are two disjoint circuits with different distance pairs...
		dstC := box2circuit[current.a]
		srcC := box2circuit[current.b]
		if srcC == dstC {
			dbg("Same Circuit! doing nothing")
			continue
		}

		// Move all boxes from srcC to dstC
		for box := range circuits[srcC] {
			circuits[dstC][box] = true
			box2circuit[box] = dstC
		}
		// move box B to the circuit of box A
		// Delete before switching circuit, or the key will not be the same
		circuits[srcC] = nil

		dbg("Merged circuit[%v]: %v", dstC, circuits[dstC])
		dbg("")

	}

	lens := []int{}
	for i, circuit := range circuits {
		dbg("%v: %v (%v)", i, circuit, len(circuit))
		lens = append(lens, len(circuit))
	}
	slices.Sort(lens)
	slices.Reverse(lens)
	dbg("FINAL lengths", lens)
	res = lens[0] * lens[1] * lens[2]

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
