package main

import (
	"testing"

	common "github.com/torbensky/adventofcode-common"
)

func TestPart1(t *testing.T) {
	got := part1(common.OpenFile("../test-input.txt"))
	want := 3434390
	if got != want {
		t.Fatalf("expected %d got %d\n", want, got)
	}
}

func TestPart2(t *testing.T) {
	got := part2(common.OpenFile("../test-input.txt"))
	want := 5148724
	if got != want {
		t.Fatalf("expected %d got %d\n", want, got)
	}
}
