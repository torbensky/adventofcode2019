package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/torbensky/adventofcode2019/common"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(common.OpenArgsFile()))
	fmt.Printf("Part 2: %d\n", part2(common.OpenArgsFile()))
}

func processData(reader io.Reader, calcFn func(mass int)) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		mass, err := strconv.Atoi(line)
		common.MustNotError(err)
		calcFn(mass)
	}

	err := scanner.Err()
	common.MustNotError(err)

}

func part1(reader io.Reader) int {
	fuelNeeded := 0
	processData(reader, func(mass int) {
		fuelNeeded += calcModuleFuel(mass)
	})

	return fuelNeeded
}

func part2(reader io.Reader) int {
	fuelNeeded := 0
	processData(reader, func(mass int) {
		fuelNeeded += calcFuelFuel(mass)
	})

	return fuelNeeded
}

func calcFuelFuel(mass int) int {
	if mass > 0 {
		fuelFuel := calcModuleFuel(mass)
		return fuelFuel + calcFuelFuel(fuelFuel)
	}

	return 0
}

func calcModuleFuel(mass int) int {
	fuel := mass/3 - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}
