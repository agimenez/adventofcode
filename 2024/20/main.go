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

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}

type grid struct {
	m map[Point]rune

	start, end Point

	w int
	h int
}

func (g grid) print(distances map[Point]int) {
	if !debug {
		return
	}
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			p := Point{x, y}
			if _, visited := distances[p]; visited {
				fmt.Print("O")
				continue
			}

			fmt.Printf("%c", g.m[p])
		}
		fmt.Println()
	}
	fmt.Println()

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

	g := NewGrid(lines)
	dbg("%v", g)
	//return
	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 = solve1(g, 2)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 = solve1(g, 20)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func NewGrid(lines []string) grid {
	g := grid{
		h: len(lines),
		w: len(lines[0]),

		m: map[Point]rune{},
	}

	for y, l := range lines {
		dbg("y = %v", y)
		for x, c := range l {
			dbg("x = %v", x)
			p := Point{x, y}
			g.m[p] = c
			switch c {
			case 'S':
				g.start = p
			case 'E':
				g.end = p
			}
		}
	}

	return g

}

func solve1(g grid, cheatTime int) int {
	res := 0

	dist := bfs(g, g.start, g.end)
	savings := g.getCheatSavings(dist, cheatTime)

	fmt.Printf("== CheatTime: %v\n", cheatTime)
	for savings, cheats := range savings {
		if savings >= 50 {
			fmt.Printf("There are %v cheats that save %v picoseconds\n", cheats, savings)
		}
		if savings >= 100 {
			res += cheats
		}
	}

	return res
}

func (g grid) getCheatSavings(distances map[Point]int, ps int) map[int]int {
	total := distances[g.end]
	cheats := map[int]int{}

	// for start, cost := range distances {
	// 	dbg("Dist %v -> %v (%v)", start, cost, total)
	// }
	for p, cost := range distances {
		dbg("Dist %v -> %v (%v)", p, cost, total)
		for dy := -ps; dy <= ps; dy++ {
			for dx := -ps; dx <= ps; dx++ {
				dif := Point{dx, dy}
				candidate := p.Sum(dif)
				dist := p.ManhattanDistance(candidate)
				if p == candidate || dist > ps {
					continue
				}

				if candidateCost, exists := distances[candidate]; exists {
					savings := candidateCost - cost - dist
					if savings > 0 {
						dbg(" -> %v cost: %v (shortcut -> %v)", candidate, candidateCost, savings)
						cheats[savings]++
					}
				}
			}
		}
	}

	return cheats

}

func bfs(g grid, start, end Point) map[Point]int {
	dist := map[Point]int{start: 0}
	q := []Point{start}

	var cur Point
	for len(q) > 0 {
		cur, q = q[0], q[1:]
		dbg("cur: %v, Q: %v", cur, q)
		dbg("dist[cur] = %v", dist[cur])
		g.print(dist)

		for _, next := range cur.Adjacent(false) {
			// skip OOB or walls
			if v, ok := g.m[next]; !ok || v == '#' {
				continue
			}

			if _, visited := dist[next]; !visited {
				dist[next] = dist[cur] + 1
				q = append(q, next)
			}
		}

		if debug {
			// time.Sleep(200 * time.Millisecond)
		}
	}

	return dist
}

func solve2(s []string) int {
	res := 0

	return res
}
