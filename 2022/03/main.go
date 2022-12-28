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
}

type Set map[rune]bool

type rucksack [2]Set

func NewRuckSack(in string) rucksack {
	r := rucksack{
		Set{},
		Set{},
	}

	r[0] = StrToSet(in[:len(in)/2])
	r[1] = StrToSet(in[len(in)/2:])

	return r
}

func StrToSet(in string) Set {
	s := Set{}

	for _, v := range in {
		s[v] = true
	}

	return s
}

func (r rucksack) getDuplicate() rune {
	for k := range r[0] {
		if _, ok := r[1][k]; ok {
			return k
		}
	}

	return 0
}

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	} else {
		return int(r - 'A' + 27)
	}

}

func findCommon(in []string) rune {
	sa := StrToSet(in[0])
	sb := StrToSet(in[1])
	sc := StrToSet(in[2])

	for k := range sa {
		_, okb := sb[k]
		_, okc := sc[k]
		if okb && okc {
			return k
		}
	}

	return 0
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
	dbg("lines: %v", lines)
	for i, line := range lines {
		r := NewRuckSack(line)
		d := r.getDuplicate()
		part1 += priority(d)

		if i%3 == 0 {
			c := findCommon(lines[i : i+3])
			part2 += priority(c)
		}
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
