package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strings"
)

var (
	debug int
)

const (
	black = '0'
	white = '1'
	alpha = '2'
)

func dbg(level int, fmt string, v ...interface{}) {
	if debug >= level {
		log.Printf(fmt, v...)
	}
}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

func checkSum(in string, wide, tall int) int {
	minz := math.MaxInt32
	res := 0
	for i, start, end := 0, 0, wide*tall; start < len(in); i, start, end = i+1, end, end+(wide*tall) {
		dbg(1, "%d: start = %d, end = %d", i, start, end)
		str := in[start:end]
		z := strings.Count(str, "0")
		dbg(2, "str = %s, z = %d (%d)", str, z, minz)
		if z < minz {
			o := strings.Count(str, "1")
			t := strings.Count(str, "2")
			dbg(2, "new max! o = %d, t = %d", o, t)

			minz = z
			res = o * t
		}

	}

	return res
}

func isTransparent(c rune) bool {
	return c == alpha
}

func isWhite(c rune) bool {
	return c == white
}

func isBlack(c rune) bool {
	return c == black
}

func processLayers(in string, wide, tall int) string {

	img := make([]rune, wide*tall)
	for start, end := 0, wide*tall; start < len(in); start, end = end, end+(wide*tall) {
		layer := in[start:end]
		for i, c := range layer {

			// initialize output to transparent
			if !isTransparent(img[i]) && !isWhite(img[i]) && !isBlack(img[i]) {
				img[i] = alpha
			}

			if isTransparent(img[i]) {
				img[i] = c
			}

		}
	}

	return string(img)
}

func renderImage(data string, wide, tall int) {

	img := strings.Builder{}
	for i, c := range data {
		col := i % wide
		if isWhite(c) {
			img.WriteRune('#')
		} else {
			img.WriteRune(' ')
		}
		if col == wide-1 {
			img.WriteRune('\n')
		}
	}

	fmt.Println(img.String())
}

func main() {

	var in string
	fmt.Scan(&in)
	wide := 25
	tall := 6

	sum := checkSum(in, wide, tall)
	image := processLayers(in, wide, tall)
	renderImage(image, wide, tall)

	log.Printf("Max output: %v", sum)
	log.Printf("image: %v", image)

}
