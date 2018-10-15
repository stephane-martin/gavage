package main

import "math/rand"

func FakeScheme(r *rand.Rand) string {
	if r.Intn(2) == 0 {
		return "http"
	}
	return "https"
}
