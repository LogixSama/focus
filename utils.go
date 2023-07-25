package main

import (
	"math/rand"
)

func isAvailable(elem string, arr []string) bool {
	for _, val := range arr {
		if elem == val {
			return true
		}
	}
	return false
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
