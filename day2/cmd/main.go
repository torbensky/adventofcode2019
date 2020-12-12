package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/torbensky/adventofcode2019/common"
)

type codeType int

const (
	add      codeType = 1
	multiply codeType = 2
	finished codeType = 99
)

type instruction struct {
	code codeType
	// parameters is expected to be length 3 for the three params
	parameters []int
}

type program struct {
	data         []int
	instructions []instruction
}

func main() {
	fmt.Printf("Part 1: %d\n", part1(common.OpenArgsFile()))
	fmt.Printf("Part 2: %d\n", part2(common.OpenArgsFile()))
}

func part1(reader io.Reader) int {
	intCodes := loadIntCodes(reader)
	intCodes[1] = 12
	intCodes[2] = 2
	return execute(intCodes)
}

func execute(intCodes []int) int {
	i := 0
Loop:
	for {
		switch codeType(intCodes[i]) {
		case add:
			intCodes[intCodes[i+3]] = intCodes[intCodes[i+1]] + intCodes[intCodes[i+2]]
			i += 4
		case multiply:
			intCodes[intCodes[i+3]] = intCodes[intCodes[i+1]] * intCodes[intCodes[i+2]]
			i += 4
		case finished:
			break Loop
		default:
			log.Fatalf("unknown code encountered: %d\n", i)
		}
	}

	return intCodes[0]
}

func part2(reader io.Reader) int {
	codes := loadIntCodes(reader)

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			codesCopy := make([]int, len(codes))
			copy(codesCopy, codes)
			codesCopy[1] = noun
			codesCopy[2] = verb
			result := execute(codesCopy)
			if result == 19690720 {
				return 100*noun + verb
			}
		}
	}

	return -1
}

func loadProgram(ints []int) program {
	var instructions []instruction
	i := 0
	for i < len(ints) {

		code := codeType(ints[i])
		next := instruction{
			code: code,
		}
		instructions = append(instructions, next)

		if code == finished {
			break
		}

		// any non-finished codes have parameters
		next.parameters = ints[i+1 : i+4]
		i += 4
	}

	return program{
		data:         ints,
		instructions: instructions,
	}
}

func loadIntCodes(reader io.Reader) []int {
	fields := bytes.Split(common.ReadAll(reader), []byte(","))
	var intCodes []int
	for _, val := range fields {
		intCodes = append(intCodes, common.Atoi(string(val)))
	}
	return intCodes
}
