package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"slices"
	"strings"

	. "github.com/agimenez/adventofcode/utils"
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

type Block struct {
	id   int
	size int
}

func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	line := strings.Trim(string(p), "\n")
	dbg("lines: %#v", line)
	inFile := true
	fileID := 0
	out := []int{}
	disk := []Block{}
	for _, c := range line {
		id := fileID
		if !inFile {
			id = -1
		}
		blocks := ToInt(string(c))
		for b := blocks; b > 0; b-- {
			out = append(out, id)
		}

		b := Block{
			id:   id,
			size: blocks,
		}
		disk = append(disk, b)

		if inFile {
			fileID++
		}

		inFile = !inFile
	}
	dbg("out: %#v", out)
	compact := compact(out)
	dbg("compact: %v", compact)
	part1 = checksum(compact)

	dbg("disk: %v", disk2Int(disk))
	defrag := defrag(disk)
	d := disk2Int(defrag)
	dbg("defrag: %v", d)
	part2 = checksum(d)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

func disk2Int(d []Block) []int {
	res := make([]int, 0, len(d))
	for _, b := range d {
		for ; b.size > 0; b.size-- {
			res = append(res, b.id)
		}
	}

	return res
}

func compact(d []int) []int {
	disk := slices.Clone(d)
	for cur, last := 0, len(disk)-1; cur < len(disk) && cur < last; cur++ {
		if disk[cur] != -1 {
			continue
		}

		for ; disk[last] == -1 && last > 0; last-- {
		}
		disk[cur] = disk[last]
		disk[last] = -1
		last--
		dbg("disk: %v", disk[:cur])
	}

	return disk
}

func defrag(disk []Block) []Block {
	res := slices.Clone(disk)
	for cur := len(res) - 1; cur > 0; cur-- {
		dbg("Checking block %d: %v", cur, res[cur])
		b := res[cur]
		if b.id == -1 {
			dbg(" -> FREE")
			continue
		}

		for free := 0; free < cur; free++ {
			f := res[free]
			dbg("  -> BLOCK %d (size %d)", f.id, f.size)
			if f.id != -1 || f.size < b.size {
				continue
			}

			// Fits in the free block, just overwrite the free block
			if f.size == b.size {
				dbg("  ---> same size")
				res[free].id = b.id
				res[free].size = b.size

				res[cur].id = -1
				break
			}

			// Fits but there's some remaining space: move block to new position
			if f.size > b.size {
				dbg("  ---> remaining space")
				newfree := Block{
					id:   -1,
					size: f.size - b.size,
				}
				res[free] = b
				res[cur].id = -1
				slices.Insert(res, free+1, newfree)
				break
			}
		}
		dbg("STATE: %v", disk2Int(res))
	}

	dbg("FINAL STATE: %v", disk2Int(res))
	return res
}

func checksum(disk []int) int {
	res := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			continue
		}

		res += i * disk[i]
	}

	return res
}
