package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
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

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) ManhattanDistance() int {
	return Abs(p.X) + Abs(p.Y) + Abs(p.Z)
}

type Particle struct {
	P, V, A Point3D
}

var particleRE = regexp.MustCompile(`p=<([0-9-]+),([0-9-]+),([0-9-]+)>, v=<([0-9-]+),([0-9-]+),([0-9-]+)>, a=<([0-9-]+),([0-9-]+),([0-9-]+)>`)

func parseParticle(s string) Particle {
	p := Particle{}
	res := particleRE.FindStringSubmatch(s)
	p.P = Point3D{
		X: ToInt(res[1]),
		Y: ToInt(res[2]),
		Z: ToInt(res[3]),
	}

	p.V = Point3D{
		X: ToInt(res[4]),
		Y: ToInt(res[5]),
		Z: ToInt(res[6]),
	}

	p.A = Point3D{
		X: ToInt(res[7]),
		Y: ToInt(res[8]),
		Z: ToInt(res[9]),
	}

	return p
}

func solve1(s []string) int {
	res := 0

	particles := make([]Particle, len(s))
	for i, line := range s {
		particles[i] = parseParticle(line)
	}

	// Closest: lower acceleration
	minA := math.MaxInt
	for i, p := range particles {
		dist := p.A.ManhattanDistance()
		if dist < minA {
			minA = dist
			res = i
		} else if dist == minA {
			if p.V.ManhattanDistance() < particles[res].V.ManhattanDistance() {
				minA = dist
				res = i
			}
		}
	}
	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	return res
}
