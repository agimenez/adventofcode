package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] > 2020 {
				break
			}

			if nums[i]+nums[j] == 2020 {
				log.Printf("%v", nums[i]*nums[j])
				os.Exit(0)
			}
		}
	}
}
