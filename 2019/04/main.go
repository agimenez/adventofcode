package main

import (
	"log"
)

const (
	Start = 134564
	End   = 585159

	debug = false
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func main() {

	compliant := 0
	for candidate := Start; candidate <= End; candidate++ {
		if checkCompliant(candidate) {
			dbg("%d is compliant", candidate)
			compliant++
		}
	}

	log.Printf("Compliant passwords: %d", compliant)

}

func checkCompliant(c int) bool {

	// this is a sentinel value that can't exist modulo 10
	last := 10

	adjacent := false
	for i := 5; i >= 0; i-- {
		curDigit := c % 10

		// check if it decreases (backwards increasing)
		if curDigit > last {
			return false
		}

		if curDigit == last {
			adjacent = true
		}

		last = curDigit
		c = c / 10
	}

	return adjacent
}
