package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/agimenez/adventofcode/utils"
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

func expand(in []string, factor int) []utils.Point {
	galaxies := []utils.Point{}
	shiftY := map[int]int{}
	colHasGalaxy := map[int]bool{}
	shift := 0
	for y, l := range in {
		shiftY[y] = shift
		for x, c := range l {
			if c == '#' {
				colHasGalaxy[x] = true
				galaxies = append(galaxies, utils.Point{x, y})
			}
		}
		if strings.Index(l, "#") == -1 {
			shift += factor
		}
	}

	shiftX := map[int]int{}
	shift = 0
	for x, _ := range in[0] {
		shiftX[x] = shift
		if !colHasGalaxy[x] {
			shift += factor
		}
	}
	dbg("shiftY: %v", shiftY)
	dbg("shiftX: %v", shiftX)

	expandedGalaxies := []utils.Point{}
	for _, g := range galaxies {
		g.X += shiftX[g.X] + factor
		g.Y += shiftY[g.Y] + factor

		expandedGalaxies = append(expandedGalaxies, g)
	}

	dbg("galaxies: %v", galaxies)
	dbg("expanded: %v", expandedGalaxies)

	return expandedGalaxies
}

func distancePairs(g []utils.Point) []int {
	d := []int{}
	for i := 0; i < len(g); i++ {
		for j := i + 1; j < len(g); j++ {
			dbg("Distance %d %v -> %d %v = %d", i, g[i], j, g[j], g[i].ManhattanDistance(g[j]))
			d = append(d, g[i].ManhattanDistance(g[j]))
		}
	}

	return d
}

func solve1(in []string, factor int) int {
	var res int
	galaxies := expand(in, factor)
	distances := distancePairs(galaxies)

	for _, d := range distances {
		res += d
	}

	return res
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
	part1 = solve1(lines, 1)
	part2 = solve1(lines, 999_999)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
