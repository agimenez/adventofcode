package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

type robot struct {
	pos Point
	vel Point
}

type grid struct {
	wide int
	tall int

	robots []robot
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
	g := grid{
		wide: 101,
		tall: 103,
		//wide:   11,
		//tall:   7,
		robots: []robot{},
	}
	for _, l := range lines {
		g = g.addRobot(l)
	}
	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	g = g.simulate(100)
	part1 = g.safetyFactor()

	dur[0] = time.Since(now)

	now = time.Now()
	part2 = solve1(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func (g grid) simulate(secs int) grid {
	dbg("Simulating %v secs", secs)
	for i := 1; i <= secs; i++ {
		dbg("== Step %v", i)
		g = g.step()
	}
	dbg("== GRID: %v", g)

	return g
}

func (g grid) step() grid {
	for i, r := range g.robots {
		//dbg("PRE  Robot %v: %v (%v)", i, r.pos, r.vel)
		pos := r.pos.Sum(r.vel)
		pos.X %= g.wide
		pos.Y %= g.tall

		if pos.X < 0 {
			pos.X += g.wide
		}

		if pos.Y < 0 {
			pos.Y += g.tall
		}

		g.robots[i].pos = pos
		//dbg("POST Robot %v: %v", i, g.robots[i].pos)
	}
	//g.printNumbers(Point{0, 0}, Point{g.wide, g.tall})

	return g
}

func (g grid) safetyFactor() int {
	res := 1
	quadrants := []struct {
		start Point
		end   Point
	}{
		{Point{0, 0}, Point{g.wide / 2, g.tall / 2}},
		{Point{g.wide/2 + 1, 0}, Point{g.wide, g.tall / 2}},
		{Point{0, g.tall/2 + 1}, Point{g.wide / 2, g.tall}},
		{Point{g.wide/2 + 1, g.tall/2 + 1}, Point{g.wide, g.tall}},
	}

	for _, q := range quadrants {
		rq := g.robotsInQuadrant(q.start, q.end)
		res *= rq
		dbg("FOUND %v robots in quadrant, res: %v", rq, res)
	}
	dbg("RETURNING res: %v", res)

	return res
}

func (g grid) robotsInQuadrant(start, end Point) int {
	res := 0
	//dbg("== robots in quadrant %v - %v", start, end)
	//g.printNumbers(start, end)
	for _, r := range g.robots {
		//dbg(" - %v", r.pos)
		if r.pos.X >= start.X && r.pos.Y >= start.Y && r.pos.X < end.X && r.pos.Y < end.Y {
			res++
			//dbg(" - %v YES!!", r.pos)
		}
	}
	dbg("robots in quadrant %v - %v: %d", start, end, res)

	return res
}

func (g grid) addRobot(s string) grid {
	var px, py, vx, vy int

	fmt.Sscanf(s, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
	r := robot{
		pos: Point{px, py},
		vel: Point{vx, vy},
	}

	g.robots = append(g.robots, r)

	return g
}

func solve1(s []string) int {
	res := 0

	return res
}

func (g grid) printNumbers(start, end Point) {
	if !debug {
		return
	}

	for y := start.Y; y < end.Y; y++ {
		for x := start.X; x < end.X; x++ {
			count := 0
			for _, r := range g.robots {
				p := Point{x, y}
				if r.pos == p {
					count++
				}
			}

			if count > 0 {
				fmt.Printf("%d", count)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func solve2(s []string) int {
	res := 0

	return res
}
