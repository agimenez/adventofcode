package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	debug = false
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

type passport map[string]string

func validPassport(p passport) bool {

	if len(p) == 8 {
		return true
	}

	if _, hasCid := p["cid"]; len(p) == 7 && !hasCid {
		return true
	}

	return false
}

type validateFunc func(s string) bool

var validators map[string]validateFunc = map[string]validateFunc{
	"byr": validateByr,
	"iyr": validateIyr,
	"eyr": validateEyr,
	"hgt": validateHgt,
	"hcl": validateHcl,
	"ecl": validateEcl,
	"pid": validatePid,
	"cid": validateCid,
}

func validateIntRange(s string, min, max int) bool {
	if len(s) != 4 {
		return false
	}

	y, _ := strconv.Atoi(s)
	if y < min || y > max {
		return false
	}

	return true
}

func validateByr(s string) bool {
	return validateIntRange(s, 1920, 2002)
}
func validateIyr(s string) bool {
	return validateIntRange(s, 2010, 2020)
}
func validateEyr(s string) bool {
	return validateIntRange(s, 2020, 2030)
}

func validateHgt(s string) bool {
	re := regexp.MustCompile(`(\d*)(\w*)`)
	v := re.FindStringSubmatch(s)[1:]

	switch v[1] {
	case "cm":
		return validateIntRange(v[0], 150, 193)
	case "in":
		return validateIntRange(v[0], 59, 76)
	}

	return false
}
func validateHcl(s string) bool {
	m, _ := regexp.MatchString(`^#[[:xdigit:]]{6}$`, s)
	return m
}
func validateEcl(s string) bool {
	switch s {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}

	return false
}
func validatePid(s string) bool {
	m, _ := regexp.MatchString(`^\d{9}$`, s)
	return m
}

func validateCid(s string) bool {
	return true
}

func validFields(p passport) bool {
	if !validPassport(p) {
		return false
	}

	for f, v := range p {
		if validators[f](v) == false {
			return false
		}
	}

	return true
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	list := []passport{}

	p := passport{}
	for s.Scan() {
		l := s.Text()
		dbg("Line: %v\n", l)
		if l == "" {
			list = append(list, p)
			p = passport{}
			continue
		}

		fields := strings.Split(l, " ")
		for _, field := range fields {
			f := strings.Split(field, ":")
			p[f[0]] = f[1]
		}

		dbg("List: %v\n", list)

	}
	list = append(list, p)

	part1, part2 := 0, 0
	for _, p := range list {
		if validPassport(p) {
			part1++
		}

		if validFields(p) {
			part2++
		}
	}

	log.Printf("Part 1: %v", part1)
	log.Printf("Part 2: %v", part2)

}
