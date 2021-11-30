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
				log.Printf("Part one: %v", nums[i]*nums[j])
			}
		}
	}

	// This is not ideal, I think we could calculate both parts within the same loop,
	// but I can't be bothered...
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] > 2020 {
					break
				}

				if nums[i]+nums[j]+nums[k] == 2020 {
					log.Printf("Part two: %v", nums[i]*nums[j]*nums[k])
				}
			}
		}
	}
}
