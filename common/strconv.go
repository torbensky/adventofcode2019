package common

import "strconv"

// Atoi is equivalent to "strconv.Atoi" except it fatally errors on error
func Atoi(s string) int {
	result, err := strconv.Atoi(s)
	MustNotError(err)
	return result
}
