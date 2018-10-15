package main

import (
	"math/rand"
	"strings"

	"github.com/icrowley/fake"
)

func FakePath(r *rand.Rand) string {
	nbWords := int(r.ExpFloat64())
	if nbWords == 0 {
		return "/"
	}
	return "/" + strings.Replace(fake.WordsN(nbWords), " ", "/", -1)
}
