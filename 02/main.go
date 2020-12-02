package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
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

func parsePolicy(l string) (int, int, string, string) {
	re := regexp.MustCompile(`(\d*)-(\d*)\s(\w*):\s(\w*)`)
	v := re.FindStringSubmatch(l)[1:]
	min, _ := strconv.Atoi(v[0])
	max, _ := strconv.Atoi(v[1])
	char := v[2]
	pw := v[3]

	return min, max, char, pw
}

func checkPolicy(l string) bool {
	min, max, char, pw := parsePolicy(l)

	c := strings.Count(pw, char)
	dbg("%v-%v %v: %v (count: %v)\n", min, max, char, pw, c)
	return c >= min && c <= max
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	correct := 0
	for s.Scan() {
		l := s.Text()

		if checkPolicy(l) {
			correct++
			dbg("%v is correct!\n", l)
		}
	}

	log.Printf("Part 1: %v\n", correct)

}
