package main

import "math/rand"

var FilterResultCorpus = []string{
	"PROXIED",
	"OBSERVED",
	"DENIED",
}

func FakeFilterResult(r *rand.Rand) string {
	return FilterResultCorpus[r.Intn(len(FilterResultCorpus))]
}
