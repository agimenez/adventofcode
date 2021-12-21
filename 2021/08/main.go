package main

import (
	"flag"
	"fmt"
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
	flag.Parse()
}

func parseEntry(s string) ([]string, []string) {
	parts := strings.Split(s, " | ")
	signals := strings.Fields(parts[0])
	outputs := strings.Fields(parts[1])

	return signals, outputs
}

func solvePart1(in []string) int {
	count := 0
	for _, entry := range in {
		_, outputValues := parseEntry(entry)
		dbg("outputs: %v", outputValues)

		for _, v := range outputValues {
			switch len(v) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}

	return count
}

type signalMap struct {
	toNumber map[string]int
	toString map[int]string
}

func (sm signalMap) add(s string, v int) {
	sm.toNumber[s] = v
	sm.toString[v] = s
}

func charMap(signal string) map[byte]bool {
	segments := make(map[byte]bool)
	for _, c := range signal {
		segments[byte(c)] = true
	}

	return segments
}

func segmentDifference(signal string, segments string) string {
	segs := charMap(segments)
	out := ""

	for i := range signal {
		if _, ok := segs[signal[i]]; ok {
			delete(segs, signal[i])
		} else {
			out += string(signal[i])
		}
	}

	dbg("SegmentDifference: %s", out)
	return out

}

func containsSegments(signal string, segments string) bool {
	segs := charMap(segments)

	for i := range signal {
		if _, ok := segs[signal[i]]; ok {
			delete(segs, signal[i])
		}
	}

	return len(segs) == 0
}
func (sm signalMap) containsNumber(s string, n int) bool {
	nsignal, ok := sm.toString[n]
	if !ok {
		return false
	}

	return containsSegments(s, nsignal)
}

func (sm signalMap) segmentDifference(n1 int, n2 int) string {
	s1, ok := sm.toString[n1]
	if !ok {
		return s1
	}

	s2, ok := sm.toString[n2]
	if !ok {
		return s1
	}

	sd := segmentDifference(s1, s2)
	return sd
}

func signalMapping(in []string) signalMap {
	sm := signalMap{
		toNumber: make(map[string]int),
		toString: make(map[int]string),
	}

	// First, trivial cases that will be used in deducing the other ones
	for _, s := range in {
		if len(s) == 2 {
			sm.add(s, 1)
		} else if len(s) == 3 {
			sm.add(s, 7)
		} else if len(s) == 4 {
			sm.add(s, 4)
		} else if len(s) == 7 {
			sm.add(s, 8)
		}
	}
	for len(sm.toNumber) < 9 {
		for _, s := range in {
			if len(s) == 5 {
				if sm.containsNumber(s, 1) {
					sm.add(s, 3)
				} else if containsSegments(s, sm.segmentDifference(4, 1)) {
					sm.add(s, 5)
				} else if !containsSegments(s, sm.segmentDifference(4, 1)) {
					sm.add(s, 2)
				}
			} else if len(s) == 6 {
				if sm.containsNumber(s, 4) {
					sm.add(s, 9)
				} else if sm.containsNumber(s, 7) && !sm.containsNumber(s, 4) {
					sm.add(s, 0)
				} else {
					sm.add(s, 6)
				}
			}
		}

	}

	return sm
}

func (sm signalMap) toValue(digits []string) int {
	var s string
	for _, d := range digits {
		dbg("Digit: %s", d)
		for signal := range sm.toNumber {
			dbg(" -> checking signal %s", signal)
			if len(segmentDifference(signal, d)) == 0 && len(segmentDifference(d, signal)) == 0 {
				dbg("Digit: %s, number: %v", d, sm.toNumber[signal])
				s += fmt.Sprintf("%d", sm.toNumber[signal])
				break
			}
		}
	}
	dbg("Digits: %s", s)
	v, _ := strconv.Atoi(s)
	return v
}

func solvePart2(in []string) int {
	sum := 0
	for _, entry := range in {
		signals, digits := parseEntry(entry)
		dbg("Signals: %v, values: %s", signals, digits)
		sm := signalMapping(signals)
		dbg("Signal mapping: %v", sm)
		sum += sm.toValue(digits)
	}

	return sum
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	part1 = solvePart1(lines)
	part2 = solvePart2(lines)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
