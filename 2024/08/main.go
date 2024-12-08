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
	//dbg("lines: %#v", lines)
	antennas := make(map[rune][]Point)
	for y, l := range lines {
		for x, c := range l {
			if c != '.' {
				antennas[c] = append(antennas[c], Point{x, y})
			}
		}
	}
	dbg("antennas: %v", antennas)
	antinodes := getAntinodes(lines, antennas)
	dbg("antinodes: %v", antinodes)
	part1 = len(antinodes)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

func getAntinodes(m []string, a map[rune][]Point) map[Point]bool {
	antinodes := map[Point]bool{}
	for freq, antennas := range a {
		dbg("Checking freq %c", freq)
		for _, antenna := range antennas {
			for _, pair := range antennas {
				if pair == antenna {
					continue
				}

				dir := pair.Sub(antenna)
				a1 := antenna.Sub(dir)
				a2 := pair.Sum(dir)
				dbg(" Antenna: %v, pair: %v, antinodes: %v, %v", antenna, pair, a1, a2)
				if _, ok := GetChInPoint(m, a1); ok {
					antinodes[a1] = true
				}
				if _, ok := GetChInPoint(m, a2); ok {
					antinodes[a2] = true
				}
			}
		}
	}

	return antinodes
}
