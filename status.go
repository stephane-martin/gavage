package main

import "math/rand"

var Status = []int64{
	201, 204, 206, 301, 302, 304, 400, 401, 403, 404, 500, 502, 503,
}

func FakeStatus(r *rand.Rand) int64 {
	if r.Float64() <= 0.8 {
		return 200
	}
	return Status[r.Intn(len(Status))]
}
