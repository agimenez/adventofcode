package main

import (
	"container/ring"
	"flag"
	"fmt"
	"io"
	"iter"
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

type KnotHash struct {
	buf   *ring.Ring
	start *ring.Ring

	cur  *ring.Ring
	skip int
}

func NewKnotHash(size int) KnotHash {
	r := ring.New(size)

	kh := KnotHash{
		buf:   r,
		start: r,
		cur:   r,
		skip:  0,
	}

	for i := range size {
		r.Value = i
		r = r.Next()
	}

	return kh
}

func (kh KnotHash) AllItems() iter.Seq[int] {
	return func(yield func(int) bool) {
		if !yield(kh.start.Value.(int)) {
			return
		}
		for p := kh.start.Next(); p != kh.start; p = p.Next() {
			if !yield(p.Value.(int)) {
				return
			}
		}
	}

}

func (kh KnotHash) String() string {

	var sb strings.Builder

	cur := kh.cur.Value
	for v := range kh.AllItems() {
		if cur == v {
			sb.WriteString(fmt.Sprintf("[%d]", cur))
		} else {
			sb.WriteString(fmt.Sprintf("%d", v))
		}

		sb.WriteRune(' ')
	}

	return sb.String()
}

func (kh *KnotHash) step(l int) {
	reverseLast := kh.cur.Move(l - 1)

	halfway := l / 2
	dbg("Step (len %d, halfway: %d): %v", l, halfway, kh)
	reverseStart := kh.cur
	for i := 0; i < halfway; i++ {
		if reverseStart == reverseLast {
			break
		}

		dbg("  >> (%d)| %v: Swapping %d <-> %d", i, kh, reverseStart.Value, reverseLast.Value)
		reverseStart.Value, reverseLast.Value = reverseLast.Value, reverseStart.Value
		dbg("  >> (%d)| %v", i, kh)
		reverseStart = reverseStart.Next()
		reverseLast = reverseLast.Prev()
	}

	kh.cur = kh.cur.Move(l + kh.skip)
	kh.skip++

}

func (kh *KnotHash) Hash(lens []int, rounds int) {
	for range rounds {
		for _, l := range lens {
			kh.step(l)
			dbg("After step: %v", kh)
			dbg("")
		}
	}
}

func (kh *KnotHash) HashText(in string, rounds int) {
	lens := make([]int, 0, len(in))

	for _, c := range in {
		lens = append(lens, int(c))
	}

	lens = slices.Concat(lens, []int{17, 31, 73, 47, 23})
	kh.Hash(lens, rounds)
}

func (kh KnotHash) SparseHash() []int {
	return slices.Collect(kh.AllItems())
}

func (kh KnotHash) DenseHash() string {
	var b strings.Builder
	sparse := kh.SparseHash()

	for block := 0; block < 16; block++ {
		res := 0
		for j := 0; j < 16; j++ {
			res ^= sparse[block*16+j]
		}
		b.WriteString(fmt.Sprintf("%02x", res))
	}

	return b.String()
}

func TestRun() {
	kh := NewKnotHash(5)
	lens := []int{3, 4, 1, 5}
	kh.Hash(lens, 1)

	v := kh.SparseHash()
	println(v[0] * v[1])
}

func solve1(s []string) int {
	res := 0

	kh := NewKnotHash(256)
	dbg("%v", kh)
	lens := CSVToIntSlice(s[0], ",")
	kh.Hash(lens, 1)
	v := kh.SparseHash()
	res = v[0] * v[1]

	TestRun()

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")
	kh := NewKnotHash(256)
	kh.HashText(s[0], 64)

	r := kh.DenseHash()
	fmt.Println(r)
	dbg("%v", r)

	return res
}
