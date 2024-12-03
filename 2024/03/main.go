package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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

func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	part1, part2 = getMuls(lines)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

var pattern = regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

func getMuls(lines []string) (int, int) {
	part1 := 0
	part2 := 0
	enabled := true
	for _, s := range lines {
		matches := pattern.FindAllStringSubmatch(s, -1)
		dbg("%v", matches)
		for _, match := range matches {
			dbg(" - %#v", match)
			switch {
			case strings.HasPrefix(match[0], "don't("):
				enabled = false
				dbg("  -> disabling!")
			case strings.HasPrefix(match[0], "do("):
				dbg("  -> enabling!")
				enabled = true
			}

			res := utils.ToInt(match[1]) * utils.ToInt(match[2])
			part1 += res
			if enabled {
				dbg("   * PART 2 (enabled): %v (%v)", part2, res)
				part2 += res
			}
		}
	}

	return part1, part2
}
