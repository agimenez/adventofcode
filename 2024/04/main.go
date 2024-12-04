package main

import (
	"flag"
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
	//dbg("lines: %#v", lines)

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
		match := findXMASDir(s, start, dir)
		if match {
			found++
		}
	}

	return found
}

func findXMASDir(s []string, start, dir Point) bool {
	str := "XMAS"
	dbg("Start: %v, dir: %v", start, dir)
	for p, i := start, 0; i < len(str); i++ {
		if c, ok := GetChInPoint(s, p); !ok || c != str[i] {
			return false
		}

		p = p.Sum(dir)
	}

	return true
}
