package common

import "log"

// MustNotError will log a fatal error if the error is not nil
func MustNotError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
