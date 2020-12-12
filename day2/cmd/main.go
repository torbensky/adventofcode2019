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

func main() {
	fmt.Printf("Part 1: %d\n", part1(common.OpenArgsFile()))
	fmt.Printf("Part 2: %d\n", part2(common.OpenArgsFile()))
}

func part1(reader io.Reader) int {
	program := loadProgram(reader)
	program[1] = 12
	program[2] = 2
	return executePart1(program)
}

func executePart1(program []int) int {
	i := 0
Loop:
	for {
		switch codeType(program[i]) {
		case add:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
			i += 4
		case multiply:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
			i += 4
		case finished:
			break Loop
		default:
			log.Fatalf("unknown code encountered: %d\n", i)
		}
	}

	return program[0]
}

func part2(reader io.Reader) int {
	// TODO: implement me
	return -1
}

func loadProgram(reader io.Reader) []int {
	fields := bytes.Split(common.ReadAll(reader), []byte(","))
	var intCodes []int
	for _, val := range fields {
		intCodes = append(intCodes, common.Atoi(string(val)))
	}
	return intCodes
}
