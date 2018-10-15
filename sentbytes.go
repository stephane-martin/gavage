package main

import "math/rand"

func FakeSentBytes(r *rand.Rand) int64 {
	var sent float64
	if r.Intn(2) == 0 {
		sent = r.NormFloat64()*50 + 250
	} else {
		sent = r.NormFloat64()*20 + 1000
	}
	if sent <= 26 {
		sent = 26
	}
	return int64(sent)
}
