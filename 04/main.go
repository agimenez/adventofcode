package main

import (
	"bufio"
	"log"
	"os"
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

type passport map[string]string

func validPassport(p passport) bool {

	if len(p) == 8 {
		return true
	}

	if _, hasCid := p["cid"]; len(p) == 7 && !hasCid {
		return true
	}

	return false
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	list := []passport{}

	p := passport{}
	for s.Scan() {
		l := s.Text()
		dbg("Line: %v\n", l)
		if l == "" {
			list = append(list, p)
			p = passport{}
			continue
		}

		fields := strings.Split(l, " ")
		for _, field := range fields {
			f := strings.Split(field, ":")
			p[f[0]] = f[1]
		}

		dbg("List: %v\n", list)

	}
	list = append(list, p)

	part1 := 0
	for _, p := range list {
		if validPassport(p) {
			part1++
		}
	}

	log.Printf("Part 1: %v", part1)
	log.Printf("Part 2: \n")

}
