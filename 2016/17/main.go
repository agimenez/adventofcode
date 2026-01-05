package main

import (
	"bytes"
	"cmp"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"iter"
	"log"
	"math"
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

type state struct {
	node Point
	cost int
	path []byte
}

func (s state) Path() string {
	return string(s.path)
}

func (s state) String() string {
	return fmt.Sprintf("(Point: %v | Cost: %v | Path: %q)\n", s.node, s.cost, s.Path())
}

func OpenDoors(hash [16]byte) iter.Seq[byte] {
	dirs := [4]byte{'U', 'D', 'L', 'R'}
	return func(yield func(r byte) bool) {
		for i := range hash[:2] {
			if hash[i]>>4 >= 0xb {
				if !yield(dirs[2*i]) {
					return
				}
			}

			if hash[i]&0x0F >= 0x0b {
				if !yield(dirs[2*i+1]) {
					return
				}
			}

		}
	}
}

func NextPoint(p Point, d byte) Point {
	switch d {
	case 'U':
		return p.Up()
	case 'D':
		return p.Down()
	case 'L':
		return p.Left()
	case 'R':
		return p.Right()
	}

	panic("Unknown direction: " + string(d))
}

func findShortestPath(pass string, start, end Point) (state, state) {
	queue := []state{{node: start, cost: start.ManhattanDistance(end), path: []byte{}}}
	distances := map[string]state{
		"": {node: P0, cost: 0, path: []byte{}},
	}
	g := NewGrid(3, 3)
	b := bytes.NewBufferString(pass)

	shortest := state{cost: math.MaxInt}
	var longest state

	for len(queue) > 0 {
		// Poor man's priority queue
		slices.SortFunc(queue, func(i, j state) int {
			return cmp.Compare(i.cost, j.cost)
		})
		dbg("Q: %v", queue)

		cur := queue[0]
		queue = queue[1:]

		dbg("CUR: %v", cur)
		if cur.node == end {
			// fmt.Printf("FOUND: %v\n", cur)
			if cur.cost > longest.cost {
				longest = cur
			}

			if cur.cost < shortest.cost {
				shortest = cur
			}

			continue
		}

		b.Truncate(len(pass))
		dbg("CURRENT pass: %v", b.String())

		b.Write(cur.path)
		dbg("CURRENT BUFFER: %v", b.String())

		hash := md5.Sum(b.Bytes())
		dbg("CURRENT HASH: %q", hex.EncodeToString(hash[:2]))

		for nextDir := range OpenDoors(hash) {
			dbg(" -> %c", nextDir)
			nextPoint := NextPoint(cur.node, nextDir)
			if g.OutOfBounds(nextPoint) {
				dbg("   -> %c INVALID", nextDir)
				continue
			}

			nextState := state{
				node: nextPoint,
				cost: distances[cur.Path()].cost + 1,
				path: bytes.Clone(append(distances[cur.Path()].path, nextDir)),
			}
			dbg("NEXT state to consider: %v", nextState)

			if _, ok := distances[nextState.Path()]; !ok {
				distances[nextState.Path()] = nextState
				dbg(" -> Added to %v", nextState.Path())

				queue = append(queue, state{
					node: nextPoint,
					path: bytes.Clone(nextState.path),
					cost: nextState.cost + nextPoint.ManhattanDistance(end),
				})
			}
		}
		dbg("==================\nDIST:\n%v", distances)
		dbg("")

	}

	return shortest, longest
}

func solve1(s []string) int {
	res := 0

	start := NewPoint(0, 0)
	end := NewPoint(3, 3)
	short, long := findShortestPath(s[0], start, end)
	fmt.Printf("SHORTEST: %v\n", short)
	fmt.Printf("LONGEST: %v\n", long)

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
