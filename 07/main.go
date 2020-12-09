package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	debug = false
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func findContainers(containedBy map[string][]string, s string) []string {
	if _, ok := containedBy[s]; !ok {
		return []string{}
	}

	bags := containedBy[s]
	for _, b := range containedBy[s] {
		bags = append(bags, findContainers(containedBy, b)...)
	}

	return bags

}

var depth int

func findContains(contains map[string]map[string]int, s string) int {
	spc := strings.Repeat(" ", 4*depth)
	dbg("[%v]%sFinding %v", depth, spc, s)
	if len(contains[s]) == 0 {
		dbg("[%v]%s -> no containers!", depth, spc)
		return 0
	}

	acc := 0
	for bag, count := range contains[s] {

		depth++
		c := findContains(contains, bag)
		dbg("[%v]%s'%v' contains a total of %v bags, times %v", depth, spc, bag, c, count)

		acc += count + count*findContains(contains, bag)
	}

	depth--
	return acc
}

func deduplicate(list []string) []string {
	uniq := map[string]struct{}{}

	for _, b := range list {
		uniq[b] = struct{}{}
	}

	res := []string{}
	for b := range uniq {
		res = append(res, b)
	}

	return res
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	part1, part2 := 0, 0
	depth = 0
	contains := make(map[string]map[string]int)
	containedBy := make(map[string][]string)
	for s.Scan() {
		l := s.Text()

		s1 := strings.Split(l, " bags contain ")
		container := s1[0]
		contains[container] = make(map[string]int)
		if s1[1] == "no other bags." {
			continue
		}

		contents := strings.Split(s1[1], ", ")
		dbg("Container: %v", container)
		for _, bag := range contents {
			words := strings.Split(bag, " ")
			num, err := strconv.Atoi(words[0])
			if err != nil {
				panic("can't parse input!")
			}

			name := words[1] + " " + words[2]
			contains[container][name] = num
			containedBy[name] = append(containedBy[name], container)

		}

		dbg("Contains: %+v", contains)
		dbg("ContainedBy: %+v", containedBy)

	}

	bags := deduplicate(findContainers(containedBy, "shiny gold"))
	part1 = len(bags)

	part2 = findContains(contains, "shiny gold")

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
