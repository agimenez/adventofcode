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
	part1 = findAllXMAS(lines)
	part2 = findAllX_MAS(lines)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

func findAllXMAS(s []string) int {
	total := 0
	for i, l := range s {
		for j := strings.Index(l, "X"); j < len(l); j++ {
			total += checkXMAS(s, Point{j, i})
		}
	}

	return total
}

func checkXMAS(s []string, start Point) int {
	found := 0
	dir0 := Point{0, 0}
	directions := []Point{
		dir0.Up(),
		dir0.Up().Right(),
		dir0.Right(),
		dir0.Right().Down(),
		dir0.Down(),
		dir0.Down().Left(),
		dir0.Left(),
		dir0.Left().Up(),
	}

	for _, dir := range directions {
		match := findStrDir(s, "XMAS", start, dir)
		if match {
			found++
		}
	}

	return found
}

func findStrDir(s []string, str string, start, dir Point) bool {
	//dbg("Start: %v, dir: %v", start, dir)
	for p, i := start, 0; i < len(str); i++ {
		if c, ok := GetChInPoint(s, p); !ok || c != str[i] {
			return false
		}

		p = p.Sum(dir)
	}

	return true
}

func findAllX_MAS(s []string) int {
	total := 0
	for i, l := range s {
		for j := 0; j < len(l); j++ {
			if s[i][j] == 'A' {
				//printBox(s, Point{j, i})
				total += checkMAS(s, Point{j, i})
			}
		}
	}

	return total
}
func checkMAS(s []string, start Point) int {
	found := 0
	dir0 := Point{0, 0}
	choices := []struct {
		start Point
		dir   Point
	}{
		{start.Up().Left(), dir0.Down().Right()},
		{start.Down().Left(), dir0.Up().Right()},
		{start.Up().Right(), dir0.Down().Left()},
		{start.Down().Right(), dir0.Up().Left()},
	}

	for _, dir := range choices {
		match := findStrDir(s, "MAS", dir.start, dir.dir)
		if match {
			dbg("Found MAS in %v", dir.dir)
			found++
		}
	}

	// We're matching 2 strings per each correct X-MAS
	return found / 2
}

func printBox(s []string, start Point) {
	dbg("Box sorrounding %v", start)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			c, ok := GetChInPoint(s, start.Sum(Point{j, i}))
			if !ok {
				fmt.Printf("%c", '-')
			} else {
				fmt.Printf("%c", c)
			}
		}
		println()
	}
}
