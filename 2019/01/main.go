package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("** Calculating mass for fuel %d\n", num)
		total += fuelForMass(num)

	}
	log.Printf("total: %d\n", total)
}

func fuelForMass(mass int) int {
	log.Printf("  -> fuelforMass(%d)\n", mass)
	if mass <= 0 {
		log.Printf("  ----> mass <= 0, returning 0")
		return 0
	}

	fuel := mass/3 - 2
	log.Printf("  ----> fuel = %d\n", fuel)
	fuel4fuel := fuelForMass(fuel)
	if fuel4fuel < 0 {
		fuel4fuel = 0
	}
	log.Printf("  ----> fuel4fuel = %d\n", fuel4fuel)
	return fuel + fuel4fuel
}
