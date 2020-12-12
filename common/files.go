package common

import (
	"io"
	"io/ioutil"
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

// ReadAll reads all data from the reader, or fatally errors if unable
func ReadAll(reader io.Reader) []byte {
	data, err := ioutil.ReadAll(reader)
	MustNotError(err)
	return data
}
