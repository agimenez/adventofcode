package main

import (
	"container/ring"
	"encoding/hex"
	"flag"
	"fmt"
	. "github.com/agimenez/adventofcode/utils"
	"io"
	"iter"
	"log"
	"math/bits"
	"os"
	"slices"
	"strings"
	"time"
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
	reverseStart := kh.cur
	for i := 0; i < halfway; i++ {
		if reverseStart == reverseLast {
			break
		}

		reverseStart.Value, reverseLast.Value = reverseLast.Value, reverseStart.Value
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

func solve1(s []string) int {
	res := 0

	for i := range 128 {
		input := fmt.Sprintf("%s-%d", s[0], i)
		kh := NewKnotHash(256)
		kh.HashText(input, 64)
		r := kh.DenseHash()
		dbg("%s -> %v", input, r)
		values, _ := hex.DecodeString(r)
		for _, byte := range values {
			res += bits.OnesCount8(byte)
		}

	}

	dbg("")
	return res
}

func hash2bits(hex []byte) [128]bool {
	ret := [128]bool{}
	for i, v := range hex {
		lo := v & 0x0f
		hi := v >> 4

		base := 8 * i
		ret[base] = hi&8 != 0
		ret[base+1] = hi&4 != 0
		ret[base+2] = hi&2 != 0
		ret[base+3] = hi&1 != 0

		ret[base+4] = lo&8 != 0
		ret[base+5] = lo&4 != 0
		ret[base+6] = lo&2 != 0
		ret[base+7] = lo&1 != 0
	}

	return ret
}

func dbgGrid(g [128][128]bool) {
	if !debug {
		return
	}

	var sb strings.Builder
	for _, row := range g {
		for _, col := range row {
			if col {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	dbg("%s", sb.String())
}

func processRegion(g [128][128]bool, x, y int, r [128][128]bool) [128][128]bool {
	p := NewPoint(x, y)
	r[x][y] = true
	for n := range p.Adjacent(false) {
		if n.X < 0 || n.X > 127 || n.Y < 0 || n.Y > 127 {
			continue
		}

		if g[n.X][n.Y] && !r[n.X][n.Y] {
			r = processRegion(g, n.X, n.Y, r)
		}
	}

	return r
}

func gridRegions(g [128][128]bool) int {
	count := 0
	r := [128][128]bool{}

	for i, row := range g {
		for j, v := range row {
			if v && !r[i][j] {
				r = processRegion(g, i, j, r)
				count++
			}
		}
	}

	return count
}

func solve2(s []string) int {
	res := 0

	dbg("========== PART 2 ===========")
	grid := [128][128]bool{}

	for i := range 128 {
		input := fmt.Sprintf("%s-%d", s[0], i)
		kh := NewKnotHash(256)
		kh.HashText(input, 64)
		r := kh.DenseHash()
		values, _ := hex.DecodeString(r)
		for _, byte := range values {
			res += bits.OnesCount8(byte)
		}
		grid[i] = hash2bits(values)
	}
	dbgGrid(grid)

	res = gridRegions(grid)

	return res
}
