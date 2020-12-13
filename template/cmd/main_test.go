package main

import (
	"testing"

	common "github.com/torbensky/adventofcode-common"
)

const (
	expectedPart1 = -999999
	expectedPart2 = -999999
)

func TestPart1(t *testing.T) {
	t.Parallel()
	got := part1(common.OpenFile("../test-input.txt"))
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
