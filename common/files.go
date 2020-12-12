package common

import (
	"os"
)

// OpenArgsFile opens the file passed by CLI arg, or fatally errors if unable
func OpenArgsFile() *os.File {
	return OpenFile(os.Args[1])
}

// OpenFile opens the file or fatally errors if unable to
func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	MustNotError(err)
	return file
}
