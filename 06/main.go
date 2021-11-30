package main

import (
	"bufio"
	"log"
	"os"
)

const (
	debug = false
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

type questions map[rune]int

func main() {

	s := bufio.NewScanner(os.Stdin)
	list := []questions{}

	q := questions{}
	groupSize := []int{0}
	groupNum := 0
	for s.Scan() {
		l := s.Text()
		dbg("Line: %v\n", l)
		if l == "" {
			list = append(list, q)
			q = questions{}
			groupSize = append(groupSize, 0)
			groupNum++
			continue
		}

		for _, question := range l {
			q[question]++
		}
		groupSize[groupNum]++

	}
	list = append(list, q)
	dbg("List: %v\n", list)
	dbg("groupsize: %v\n", groupSize)

	part1, part2 := 0, 0
	for i, q := range list {
		part1 += len(q)
		for _, p := range q {
			if p == groupSize[i] {
				part2++
			}
		}
	}

	log.Printf("Part 1: %v", part1)
	log.Printf("Part 2: %v", part2)

}
