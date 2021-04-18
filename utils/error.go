package utils

import "log"

// Check - checks for errors
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}