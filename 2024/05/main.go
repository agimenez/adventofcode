package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"slices"
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

	inRules := true
	rules := make(map[int][]int)
	for _, l := range lines {
		if l == "" {
			inRules = false
			dbg("Rules: %v", rules)
			continue
		}

		if inRules {
			parts := strings.Split(l, "|")
			before := utils.ToInt(parts[0])
			after := utils.ToInt(parts[1])
			rules[before] = append(rules[before], after)
		} else {
			pages := utils.CSVToIntSlice(l, ",")
			dbg("Pages: %v", pages)

			// This should return -1 when page "a" goes before page "b", 1 otherwise
			cmp := func(a, b int) int {
				dbg("Comparing %v and %v (%v)", a, b, rules[b])
				for _, v := range rules[b] {
					dbg(" - %v", v)
					if a == v {
						dbg(" -    -> FOUND (%v before %v)", a, b)
						return 1
					}

				}
				return -1

			}

			if slices.IsSortedFunc(pages, cmp) {
				dbg("%v is SORTED, returning %d", l, pages[len(pages)/2])
				part1 += pages[len(pages)/2]
			} else {
				dbg("Pages UNSORTED: %v", pages)
				slices.SortFunc(pages, cmp)
				dbg("Pages   SORTED: %v", pages)
				part2 += pages[len(pages)/2]
			}

		}

	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
