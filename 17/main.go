package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	flag.Parse()
}

type Cube struct {
	x, y, z int
}

func (c Cube) Add(c2 Cube) Cube {
	return Cube{
		c.x + c2.x,
		c.y + c2.y,
		c.z + c2.z,
	}
}

func (c Cube) Neighbours() []Cube {
	n := []Cube{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}

				n = append(n, c.Add(Cube{x, y, z}))
			}
		}
	}

	return n
}

func (c Cube) ActiveNeighbours(active map[Cube]bool) int {
	count := 0

	for _, n := range c.Neighbours() {
		if active[n] {
			count++
		}
	}

	return count
}

func mapCopy(src map[Cube]bool) map[Cube]bool {
	dst := map[Cube]bool{}
	for k, v := range src {
		dst[k] = v
	}

	return dst
}

func CountActive(active map[Cube]bool) int {
	count := 0

	for _, a := range active {
		if a {
			count++
		}
	}

	return count
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	active := make(map[Cube]bool)
	for y := range lines {
		for x, c := range lines[y] {
			if c == '#' {
				active[Cube{x, y, 0}] = true
			}
		}
	}

	for i := 0; i < 6; i++ {
		prev := mapCopy(active)
		for c := range active {
			for _, n := range c.Neighbours() {
				// make them explicitly present in the map to calculate their
				// possible values
				prev[n] = prev[n]
			}
		}

		for cube, act := range prev {
			n := cube.ActiveNeighbours(prev)
			if act && (n != 2 && n != 3) {
				delete(active, cube)
			} else if n == 3 {
				active[cube] = true
			}
		}

	}

	part1 = CountActive(active)
	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
