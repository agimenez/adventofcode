package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	debug = true
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	part1, part2 := 0, 0
	contains := make(map[string]map[string]int)
	containedBy := make(map[string][]string)
	for s.Scan() {
		l := s.Text()

		s1 := strings.Split(l, " bags contain ")
		if s1[1] == "no other bags." {
			continue
		}

		container := s1[0]
		contains[container] = make(map[string]int)
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

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
