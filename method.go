package main

import "math/rand"

var Method = []string{
	"CONNECT",
	"GET",
	"POST",
	"HEAD",
	"OPTIONS",
	"PUT",
	"TUNNEL",
	"LIST",
}

func FakeMethod(r *rand.Rand) string {
	methodIdx := int64(r.ExpFloat64())
	if methodIdx >= int64(len(Method)) {
		methodIdx = int64(len(Method) - 1)
	}
	return Method[methodIdx]
}
