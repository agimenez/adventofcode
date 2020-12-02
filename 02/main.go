package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const (
	debug = false
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	nums := []int{}
	for s.Scan() {
		l := s.Text()
		n, _ := strconv.Atoi(l)
		nums = append(nums, n)
	}

}
