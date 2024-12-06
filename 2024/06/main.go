package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	//dbg("lines: %#v", lines)
	var start Point
	for y := 0; y < len(lines); y++ {
		idx := strings.IndexByte(lines[y], '^')
		if idx != -1 {
			start = Point{idx, y}
		}
	}
	visited := walkMap(lines, start, Point{0, 0}.Up(), nil)
	part1 = len(visited)
	for v := range visited {
		res := walkMap(lines, start, Point{0, 0}.Up(), &v)

		// We have a cycle
		if res == nil {
			part2++
		}
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

func printMap(m []string, cur, dir Point) {
	if !debug {
		return
	}

	for y, l := range m {
		for x, c := range l {
			p := Point{x, y}
			if cur == p {
				fmt.Printf("x")
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Println()
	}
}

type movement struct {
	pos Point
	dir Point
}

func walkMap(m []string, start, dir Point, injected *Point) map[Point]bool {

	visited := map[Point]bool{}
	cycle := map[movement]bool{}
	for {
		visited[start] = true
		cur := movement{start, dir}
		if cycle[cur] {
			return nil
		}
		cycle[cur] = true
		next := start.Sum(dir)
		dbg("* Point: %v, dir: %v, next: %v", start, dir, next)
		c, inside := GetChInPoint(m, next)
		if !inside {
			break
		}

		if c == '#' || (injected != nil && *injected == next) {
			dir = dir.Rotate90CW()
			continue
		} else {

			start = next
		}
		printMap(m, start, dir)
	}

	return visited
}
