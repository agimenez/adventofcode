package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
}

var config = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
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
	for game, l := range lines {
		dbg("====== NEW LINE ======")
		dbg(l)
		subsets := strings.Split(strings.Split(l, ": ")[1], "; ")
		valid := true
		minCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, s := range subsets {
			dbg("Subset: %q", s)
			cubes := strings.Split(s, ", ")
			for _, c := range cubes {
				dbg("  -> cube: '%q", c)
				parts := strings.Split(c, " ")
				num := parts[0]
				color := parts[1]
				n, _ := strconv.Atoi(num)
				if n > config[color] {
					valid = false
				}

				if n > minCubes[color] {
					minCubes[color] = n
				}
			}
		}

		if valid {
			part1 += game + 1
		}

		mult := 1
		for _, v := range minCubes {
			mult *= v
		}

		part2 += mult
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
