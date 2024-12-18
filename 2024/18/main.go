package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"sort"
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
	part1, part2 = solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	//part2 = solve2(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func solve1(s []string) (int, int) {
	res := math.MaxInt
	res2 := 0
	start, end := findNodes(s, 'S', 'E')

	dbg("Start: %v, End: %v", start, end)
	res, paths := findLowestScore(s, start, end)
	res2 = len(paths[res])

	return res, res2
}

type step struct {
	node directedNode
	cost int
	path []Point
}

type directedNode struct {
	pos Point
	dir Point
}

func findLowestScore(s []string, start, end Point) (int, map[int]map[Point]bool) {
	initialNode := directedNode{pos: start, dir: P0.Right()}
	queue := []step{
		{
			node: initialNode,
			cost: 0,
			path: []Point{start},
		}}

	distances := map[directedNode]int{}

	lowestCost := math.MaxInt
	bestPaths := map[int]map[Point]bool{}

	for len(queue) > 0 {
		// Poor man's priority queue
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].cost < queue[j].cost
		})

		cur := queue[0]
		queue = queue[1:]

		dbg("cur: %v, len: %v", cur, len(queue))
		printMap(s, cur.node, cur.path)

		if cur.cost > lowestCost {
			dbg("Abandoning higher cost path: %+v", cur)
			continue
		}

		if cur.node.pos == end {
			dbg("FOUND at %+v", cur)
			if cur.cost <= lowestCost {
				if _, ok := bestPaths[cur.cost]; !ok {
					bestPaths[cur.cost] = map[Point]bool{}
				}

				for _, point := range cur.path {
					bestPaths[cur.cost][point] = true
				}
				lowestCost = cur.cost
				dbg("Lowest cost path: %v", cur.path)

				continue
			}
		}

		// 3 options here: keep same direction, and 2* 90º rotations
		directions := []Point{
			cur.node.dir,
			cur.node.dir.Rotate90CCW(),
			cur.node.dir.Rotate90CW(),
		}

		for _, neighDir := range directions {
			nextNode := directedNode{
				pos: cur.node.pos.Sum(neighDir),
				dir: neighDir,
			}

			if ch, ok := GetChInPoint(s, nextNode.pos); ok && ch != '#' {
				nextCost := cur.cost + 1
				if neighDir != cur.node.dir {
					nextCost += 1000
				}

				if prevCost, ok := distances[nextNode]; ok {
					if prevCost < nextCost {
						continue
					}
				}
				distances[nextNode] = nextCost

				newPath := slices.Clone(cur.path)
				newPath = append(newPath, nextNode.pos)

				queue = append(queue,
					step{
						node: nextNode,
						cost: nextCost,
						path: newPath,
					},
				)

			}
		}

	}

	return lowestCost, bestPaths
}

func findNodes(s []string, start, end rune) (Point, Point) {
	startPoint := Point{}
	endPoint := Point{}

	for y, l := range s {
		for x, c := range l {
			switch c {
			case start:
				startPoint = Point{x, y}
			case end:
				endPoint = Point{x, y}

			}
		}
	}

	return startPoint, endPoint
}

func printMap(m []string, dn directedNode, path []Point) {
	if !debug {
		return
	}

	var ch rune
	switch dn.dir {
	case P0.Right():
		ch = '>'
	case P0.Down():
		ch = 'v'
	case P0.Left():
		ch = '<'
	case P0.Up():
		ch = '^'
	}

	points := map[Point]bool{}
	for _, p := range path {
		points[p] = true
	}

	for y, l := range m {
		for x, c := range l {
			p := Point{x, y}
			if _, ok := points[p]; ok {
				c = '*'
			}
			if dn.pos == p {
				c = ch
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}

	// fmt.Println()
}
func solve2(s []string) int {
	res := 0

	return res
}
