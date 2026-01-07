package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"maps"
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

type Node struct {
	size, used, avail int
}

func ReadCluster(df []string) map[Point]Node {
	cluster := map[Point]Node{}
	for _, line := range df {
		// Filesystem              Size  Used  Avail  Use%
		// /dev/grid/node-x0-y0     85T   67T    18T   78%
		parts := strings.Fields(line)
		coords := strings.Split(parts[0], "-")
		p := NewPoint(ToInt(coords[1][1:]), ToInt(coords[2][1:]))

		cluster[p] = Node{
			size:  ToInt(parts[1][:len(parts[1])-1]),
			used:  ToInt(parts[2][:len(parts[2])-1]),
			avail: ToInt(parts[3][:len(parts[3])-1]),
		}

	}

	return cluster
}

func (n Node) FitsInto(n2 Node) bool {
	return n.used != 0 && n.used <= n2.avail
}

func solve1(s []string) int {
	res := 0

	cluster := ReadCluster(s[2:])
	nodes := slices.Collect(maps.Values(cluster))
	for pair := range Combinations(nodes, 2) {
		dbg("%+v -> %+v", pair[0], pair[1])
		if pair[0].FitsInto(pair[1]) || pair[1].FitsInto(pair[0]) {
			res++
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
