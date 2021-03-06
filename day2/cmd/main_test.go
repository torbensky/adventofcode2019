package main

import (
	"testing"

	common "github.com/torbensky/adventofcode-common"
)

const (
	expectedPart1 = 3500
	expectedPart2 = 4925
)

func TestPart1(t *testing.T) {
	t.Parallel()
	prog := loadProgram(loadIntCodes(common.OpenFile("../test-input.txt")))
	got := prog.execute()
	want := expectedPart1
	if got != want {
		t.Fatalf("expected %d got %d\n", want, got)
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()
	got := part2(common.OpenFile("../input.txt"))
	want := expectedPart2
	if got != want {
		t.Fatalf("expected %d got %d\n", want, got)
	}
}
