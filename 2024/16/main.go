package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
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
	part1 = solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 = solve2(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func solve1(s []string) int {
	res := math.MaxInt
	start, end := findNodes(s, 'S', 'E')

	dbg("Start: %v, End: %v", start, end)
	dist, prev := findLowestScore(s, start, end)
	minScoreNode := directedNode{}
	for n, v := range dist {
		dbg("dist[%v]: %v", n.pos, v)
		if n.pos == end {
			if v < res {
				dbg(" - new min score for %v: %v", n, v)
				res = v
				minScoreNode = n
			}
		}

	}

	dbg("PATH")
	dbg("%v, %+v", prev[minScoreNode], prev)
	// for cur := prev[minScoreNode]; cur.pos != start; cur = prev[cur] {
	// 	dbg(" - %v", cur.pos)
	// }

	return res
}

type step struct {
	node directedNode
	cost int
}

type directedNode struct {
	pos Point
	dir Point
}

func findLowestScore(s []string, start, end Point) (map[directedNode]int, map[directedNode]directedNode) {
	queue := []step{
		{node: directedNode{pos: start, dir: P0.Right()}, cost: 0},
	}

	distances := map[directedNode]int{}
	previous := map[directedNode]directedNode{}

	for len(queue) > 0 {
		// Poor man's priority queue
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].cost < queue[j].cost
		})

		cur := queue[0]
		queue = queue[1:]

		printMap(s, cur.node, nil)
		if _, ok := distances[cur.node]; ok {
			dbg("Node %+v already visited", cur.node)
			continue
		}
		distances[cur.node] = cur.cost

		dbg("cur: %v, len: %v", cur, len(queue))
		if cur.node.pos == end {
			dbg("FOUND at %+v", cur)
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
				nextCost := 1
				if neighDir != cur.node.dir {
					nextCost += 1000
				}
				alt := cur.cost + nextCost
				queue = append(queue,
					step{
						node: nextNode,
						cost: alt,
					},
				)
				dbg("Adding previous[%v] -> %v", nextNode, cur.node)

				previous[nextNode] = cur.node

			}
		}

	}

	return distances, previous
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

func printMap(m []string, dn directedNode, path map[directedNode]directedNode) {
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

	for y, l := range m {
		for x, c := range l {
			p := Point{x, y}
			if dn.pos == p {
				fmt.Printf("%c", ch)
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Println()
	}

	// fmt.Println()
}
func solve2(s []string) int {
	res := 0

	return res
}
