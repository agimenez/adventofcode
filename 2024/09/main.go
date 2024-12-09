package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
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
	for _, c := range line {
		id := fileID
		if !inFile {
			id = -1
		}
		for blocks := ToInt(string(c)); blocks > 0; blocks-- {
			out = append(out, id)
		}

		if inFile {
			fileID++
		}

		inFile = !inFile
	}
	dbg("out: %#v", out)
	defrag := defrag(out)
	dbg("defrag: %v", defrag)
	part1 = checksum(defrag)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

func defrag(disk []int) []int {
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

func checksum(disk []int) int {
	res := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			break
		}

		res += i * disk[i]
	}

	return res
}
