package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"iter"
	"log"
	"os"
	"strings"
	"time"
	// . "github.com/agimenez/adventofcode/utils"
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

type Crypt struct {
	keys     []string
	validIdx []int

	salt string // salt to generate next keys
	seq  int    // sequence to generate next keys
	idx  int    // Current next key read index

	sfactor int // Stretch factor
}

func NewCrypt(salt string, stretch int) Crypt {
	return Crypt{
		keys:     []string{},
		validIdx: []int{},

		salt:    salt,
		sfactor: stretch,
		seq:     0,
		idx:     0,
	}
}

func (c Crypt) Stats() string {
	return fmt.Sprintf("Keys: %v, valid: %v, seq: %v, idx: %v, salt: %v", len(c.keys), len(c.validIdx), c.seq, c.idx, c.salt)
}

func (c *Crypt) genNext() (string, int) {
	var key string
	var idx int

	d := md5.Sum([]byte(fmt.Sprintf("%s%d", c.salt, c.seq)))

	for range c.sfactor {
		d = md5.Sum([]byte(hex.EncodeToString(d[:])))
	}

	key = fmt.Sprintf("%x", d)
	idx = c.seq

	c.keys = append(c.keys, key)
	c.seq++

	return key, idx
}

// next returns the next key and its index
// The returned key might not be valid
func (c *Crypt) nextKey() (string, int) {
	var key string
	var idx int
	if c.idx < c.seq {
		key = c.keys[c.idx]
		idx = c.idx
		c.idx++
		return key, idx
	}

	key, idx = c.genNext()
	c.idx++

	return key, idx
}

// findThree finds triplets of the same character, and returns
// the index, or -1 if not found
func findThree(s string) int {
	for i := range s[:len(s)-2] {
		if s[i] == s[i+1] && s[i+1] == s[i+2] {
			return i
		}
	}

	return -1
}

func (c *Crypt) readKey(offset int) string {
	var key string

	if offset < c.seq {
		return c.keys[offset]
	}

	key, _ = c.genNext()

	return key
}

func (c *Crypt) getNextKeys(max int) iter.Seq2[string, int] {
	return func(yield func(k string, idx int) bool) {
		for i := 0; i < max; i++ {
			offset := i + c.idx
			if !yield(c.readKey(offset), offset) {
				return
			}

		}
	}
}

func (c *Crypt) hasFiveOf(b byte) bool {
	fiveOf := strings.Repeat(string(b), 5)

	for key := range c.getNextKeys(1000) {
		if strings.Contains(key, fiveOf) {
			return true
		}
	}

	return false
}

func (c *Crypt) IsValid(key string) bool {
	idx := findThree(key)
	if idx == -1 {
		dbg("  >> NO THREE")
		return false
	}
	dbg("  >> HAS THREE")

	if !c.hasFiveOf(key[idx]) {
		dbg("  >> NO FIVES")
		return false
	}
	dbg("  >> VALID!")

	return true
}

// GetKey returns a string representing the next valid key, and the index of that key.
func (c *Crypt) GetKey() (string, int) {

	var key string
	var keyIdx int

	for {
		key, keyIdx = c.nextKey()
		dbg("Got Key: %q (idx: %v)", key, keyIdx)
		if c.IsValid(key) {
			c.validIdx = append(c.validIdx, keyIdx)
			break
		}

		dbg("%v", c.Stats())
	}

	return key, keyIdx

}

func solve1(s []string) int {
	res := 0

	c := NewCrypt(s[0], 0)
	for valid := 1; valid <= 64; valid++ {
		var k string
		k, res = c.GetKey()
		dbg("== VALID KEY: %q (%d)", k, res)
	}

	dbg("==== VALID KEYS ====")
	for idx := range c.validIdx {
		keyIdx := c.validIdx[idx]
		// dbg("[%d][%d] %q", idx, keyIdx, c.keys[keyIdx])
		dbg("%d", keyIdx)
	}
	dbg("%v", c.Stats())

	return res
}

func solve2(s []string) int {
	res := 0
	dbg("====== PART 2 =====")
	c := NewCrypt(s[0], 2016)
	for valid := 1; valid <= 64; valid++ {
		var k string
		k, res = c.GetKey()
		dbg("== VALID KEY: %q (%d)", k, res)
	}
	dbg("==== VALID KEYS ====")
	for idx := range c.validIdx {
		keyIdx := c.validIdx[idx]
		// dbg("[%d][%d] %q", idx, keyIdx, c.keys[keyIdx])
		fmt.Printf("%d\n", keyIdx)
	}
	dbg("%v", c.Stats())

	return res
}
