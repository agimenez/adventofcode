package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

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

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	part1 = countScores(lines)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

func countScores(s []string) int {
	count := 0
	m := parse(s)
	dbg("Parsed: %v", m)
	starts := findTrailHeads(m)
	dbg("Starts: %v", starts)
	for _, s := range starts {
		count += scoreTrailHead(m, s)
	}

	return count
}

func parse(s []string) map[Point]int {
	m := map[Point]int{}
	for y, l := range s {
		for x, c := range l {
			v, err := strconv.Atoi(string(c))
			if err != nil {
				continue
			}

			m[Point{x, y}] = v
		}
	}

	return m
}

func findTrailHeads(m map[Point]int) []Point {
	h := []Point{}
	for p := range m {
		if v, ok := m[p]; ok && v == 0 {
			h = append(h, p)
		}
	}

	return h
}

func scoreTrailHead(m map[Point]int, start Point) int {
	paths := 0
	queue := []Point{start}

	visited := map[Point]bool{}
	for len(queue) > 0 {
		dbg("Queue: %v", queue)
		cur := queue[0]
		queue = queue[1:]
		dbg("Cur: %v", cur)

		// Do not add to queue if already processed
		if visited[cur] {
			continue
		}
		visited[cur] = true

		if v, ok := m[cur]; ok {
			if v == 9 {
				paths++
				continue
			}
		}

		adj := getAdjacent(m, cur)
		queue = append(queue, adj...)
	}

	return paths
}

func getAdjacent(m map[Point]int, p Point) []Point {
	adj := []Point{}

	candidates := []Point{
		p.Up(),
		p.Right(),
		p.Down(),
		p.Left(),
	}

	for _, c := range candidates {
		if v, ok := m[c]; ok && v == m[p]+1 {
			adj = append(adj, c)
		}
	}

	return adj
}
