package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	. "github.com/agimenez/adventofcode/utils"
)

var (
	debug bool
)

func dbg(f string, v ...interface{}) {
	if debug {
		fmt.Printf(f+"\n", v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}

func main() {
	flag.Parse()

	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	part1, part2, dur1, dur2 := solve(lines)
	log.Printf("Part 1 (%v): %v\n", dur1, part1)
	log.Printf("Part 2 (%v): %v\n", dur2, part2)

}

func solve(lines []string) (int, int, time.Duration, time.Duration) {
	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 := solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 := solve2(lines)
	dur[1] = time.Since(now)

	return part1, part2, dur[0], dur[1]

}

type Memory struct {
	banks []int
}

func (m Memory) Equals(m2 Memory) bool {
	return slices.Equal(m.banks, m2.banks)
}

func (m Memory) Clone() Memory {
	m.banks = slices.Clone(m.banks)

	return m
}

func (m Memory) Serial() string {
	return fmt.Sprintf("%v", m.banks)
}

func parseBanks(s string) Memory {
	banks := strings.Fields(s)
	m := Memory{
		banks: make([]int, len(banks)),
	}

	for i, bank := range banks {
		m.banks[i] = ToInt(bank)
	}

	return m
}

func (m Memory) Highest() (int, int) {
	idx, count := 0, 0

	for i, c := range m.banks {
		if c > count {
			idx = i
			count = m.banks[i]
		}
	}

	return idx, count
}

func (m Memory) Redistribute() Memory {
	m = m.Clone()

	dbg("REDISTRIBUTING %v", m)
	maxIdx, blocks := m.Highest()
	dbg("  >> IDX: %d, BLOCKS: %d", maxIdx, blocks)

	m.banks[maxIdx] = 0
	for i := (maxIdx + 1) % len(m.banks); blocks > 0; i = (i + 1) % len(m.banks) {
		dbg(" >> [%d] %v", blocks, m)
		m.banks[i]++
		blocks--
	}

	dbg("FINAL: %v (%s)", m, m.Serial())
	dbg("")

	return m
}

func redistribute(m Memory) int {
	steps := 0
	seen := map[string]bool{m.Serial(): true}
	for {
		m = m.Redistribute()
		steps++
		if ok := seen[m.Serial()]; ok {
			break
		}
		seen[m.Serial()] = true
	}

	return steps
}

func solve1(s []string) int {
	res := 0

	mem := parseBanks(s[0])
	dbg("%v", mem)
	res = redistribute(mem)

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")
	m := parseBanks(s[0])
	steps := 0
	seen := map[string]int{m.Serial(): 1}
	for {
		m = m.Redistribute()
		steps++
		if _, ok := seen[m.Serial()]; ok {
			break
		}
		seen[m.Serial()] = steps
	}
	res = steps - seen[m.Serial()]

	return res
}
