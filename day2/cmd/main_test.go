package main

import (
	"testing"

	"github.com/torbensky/adventofcode2019/common"
)

const (
	expectedPart1 = 3500
	expectedPart2 = -999999
)

func TestPart1(t *testing.T) {
	t.Parallel()
	got := executePart1(loadProgram(common.OpenFile("../test-input.txt")))
	want := expectedPart1
	if got != want {
		t.Fatalf("expected %d got %d\n", want, got)
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()
	got := part2(common.OpenFile("../test-input.txt"))
	want := expectedPart2
	if got != want {
		t.Fatalf("expected %d got %d\n", want, got)
	}
}
