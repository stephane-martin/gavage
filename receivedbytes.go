package main

import "math/rand"

func FakeReceivedBytes(r *rand.Rand) int64 {
	received := r.NormFloat64()*900000 + 1000000
	if received <= 0 {
		received = 0
	}
	return int64(received)
}
