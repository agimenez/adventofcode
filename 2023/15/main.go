package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"slices"
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

func hash(s string) int {
	current := 0
	for _, c := range s {
		current += int(c)
		current *= 17
		current %= 256
	}
	//dbg("hash(%s) = %d", s, current)

	return current
}

type lens struct {
	label       string
	focalLength int
}

type hashmap [256][]lens

func toInt(s string) int {
	v, _ := strconv.Atoi(s)

	return v
}

func (hm *hashmap) removeLens(label string) {
	box := hash(label)

	dbg("removeLens %q before (box %d): %v", label, box, hm[box])
	hm[box] = slices.DeleteFunc(hm[box], func(l lens) bool {
		if l.label == label {
			return true
		}

		return false
	})
	dbg("removeLens after: %v", hm[box])
}

func (hm *hashmap) addLens(label string, fl int) {
	box := hash(label)

	idx := slices.IndexFunc(hm[box], func(l lens) bool {
		if l.label == label {
			return true
		}

		return false
	})
	dbg("Found index: %d", idx)
	if idx == -1 {
		hm[box] = append(hm[box], lens{label, fl})
	} else {
		hm[box][idx].focalLength = fl
	}
}

func (hm *hashmap) step(s string) {
	dbg("Step: %s", s)
	if s[len(s)-1] == '-' {
		hm.removeLens(s[:len(s)-1])
		return
	}

	parts := strings.Split(s, "=")
	label, focalLength := parts[0], parts[1]
	hm.addLens(label, toInt(focalLength))

}

func (hm *hashmap) focusingPower() int {
	ret := 0
	for box := range hm {
		for lens := range hm[box] {
			ret += (box + 1) * ((lens + 1) * hm[box][lens].focalLength)
		}
	}

	return ret
}

func (hm hashmap) String() string {
	b := strings.Builder{}
	for i := range hm {
		if len(hm[i]) == 0 {
			continue
		}

		b.WriteString(fmt.Sprintf("Box %-3d: %v\n", i, hm[i]))
	}

	return b.String()
}

func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	hm := hashmap{}
	seqs := strings.Split(string(p)[:len(p)-1], ",")
	for _, s := range seqs {
		part1 += hash(s)
		hm.step(s)

		dbg("\n%v\n", hm)
	}

	part2 = hm.focusingPower()

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
