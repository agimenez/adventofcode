package main

import (
	"flag"
	"io"
	"log"
	"maps"
	"os"
	"strings"
	"time"
	//. "github.com/agimenez/adventofcode/utils"
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

// Graph: Adjacency list/map
type graph map[string]map[string]bool

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

	g := genGraph(lines)
	dbg("Graph: %v", g)

	var now time.Time
	var dur [2]time.Duration

	now = time.Now()

	part1 = solve1(g)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 = solve2(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func genGraph(lines []string) graph {
	g := graph{}
	for _, l := range lines {
		conns := strings.Split(l, "-")

		if _, ok := g[conns[0]]; !ok {
			g[conns[0]] = map[string]bool{}
		}

		if _, ok := g[conns[1]]; !ok {
			g[conns[1]] = map[string]bool{}
		}

		g[conns[0]][conns[1]] = true
		g[conns[1]][conns[0]] = true
	}

	return g
}

func (g graph) vertices() map[string]bool {
	m := map[string]bool{}

	for v := range g {
		m[v] = true
	}

	return m
}

// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
// R (Current Clique): A set representing the current clique being constructed.
// P (Potential Candidates): A set of vertices that can be added to R to form a larger clique.
// X (Excluded Vertices): A set of vertices that have already been processed and should not be reconsidered for the current clique.
func (g graph) BronKerbosch(R, P, X map[string]bool) []map[string]bool {
	clique := []map[string]bool{}

	// This is supposed to be as follows:
	// if len(P) == 0 && len(X) == 0 {
	// However, that finds maxima cliques (not groups of 3, which could be part of a bigger one),
	// so we just return when we have a clique of size=3
	if len(R) == 3 {
		clique = append(clique, R)
		return clique
	}

	for v := range P {
		// add v to P
		newR := maps.Clone(R)
		newR[v] = true

		// Intersection betwween P and X, and the neighbours
		newP := map[string]bool{}
		newX := map[string]bool{}
		for n := range g[v] {

			if P[n] {
				newP[n] = true
			}

			if X[n] {
				newX[n] = true
			}
		}
		clique = append(clique, g.BronKerbosch(newR, newP, newX)...)
		delete(P, v)
		X[v] = true
	}
	return clique
}
func solve1(g graph) int {
	res := 0

	R := map[string]bool{}
	P := g.vertices()
	X := map[string]bool{}
	cliques := g.BronKerbosch(R, P, X)

	for _, c := range cliques {
		dbg("Clique: %v", c)
		if len(c) < 3 {
			continue
		}

		for v := range c {
			if v[0] == 't' {
				res++
				dbg(" -> YES!")
				break
			}
		}

	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
