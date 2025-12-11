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

type rect struct {
	p1, p2 Point
}

func (r rect) area() int {
	dbg("%v x %v", Abs(r.p2.X-r.p1.X)+1, Abs(r.p2.Y-r.p1.Y)+1)
	return (Abs(r.p2.X-r.p1.X) + 1) * (Abs(r.p2.Y-r.p1.Y) + 1)
}

func solve1(s []string) int {
	res := 0
	tiles := []Point{}
	for _, line := range s {
		parts := strings.Split(line, ",")
		tiles = append(tiles, Point{ToInt(parts[0]), ToInt(parts[1])})
	}

	for i, p1 := range tiles {
		for _, p2 := range tiles[i+1:] {
			r := rect{p1, p2}
			area := r.area()
			dbg("%v - %v: %v", p1, p2, area)
			if area > res {
				res = area
			}
		}
	}

	return res
}

func (r rect) MinX() int {
	return min(r.p1.X, r.p2.X)
}

func (r rect) MaxX() int {
	return max(r.p1.X, r.p2.X)
}

func (r rect) MinY() int {
	return min(r.p1.Y, r.p2.Y)
}

func (r rect) MaxY() int {
	return max(r.p1.Y, r.p2.Y)
}

func AABB(r1, r2 rect) bool {

	return r1.MinX() < r2.MaxX() &&
		r1.MaxX() > r2.MinX() &&
		r1.MaxY() > r2.MinY() &&
		r1.MinY() < r2.MaxY()
}

func solve2(s []string) int {
	res := 0
	tiles := []Point{}
	for _, line := range s {
		parts := strings.Split(line, ",")
		tiles = append(tiles, Point{ToInt(parts[0]), ToInt(parts[1])})
	}

	rects := []rect{}
	for i, p1 := range tiles {
		for _, p2 := range tiles[i+1:] {
			rects = append(rects, rect{p1, p2})
		}
	}
	slices.SortFunc(rects, func(a, b rect) int {
		return b.area() - a.area()
	})

	dbg("%v rectangles found", len(rects))
	// https://www.reddit.com/r/adventofcode/comments/1pibab2/comment/nt87ndo/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
	// https://manbeardgames.github.io/docs/tutorials/monogame-3-8/collision-detection/aabb-collision/
	// Check AABB collisions between my rectangle and the "line" rectangle defined by the edges
	for i, r1 := range rects {
		dbg("%v... (%v) Area: %v", i, r1, r1.area())
		collisions := 0
		for t := range tiles {
			r2 := rect{tiles[t], tiles[(t+1)%len(tiles)]}
			dbg("  >> CHECKING COLLISION: %v - %v", r1, r2)
			if AABB(r1, r2) {
				dbg("  >>>> COLLISION!!!")
				collisions++
			}
		}
		if collisions == 0 && r1.area() > res {
			res = r1.area()
		}

	}

	return res
}
