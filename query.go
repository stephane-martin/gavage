package main

import (
	"math/rand"
	"strings"

	"github.com/icrowley/fake"
)

func FakeQuery(r *rand.Rand) string {
	nbWords := int(r.ExpFloat64())
	if nbWords == 0 {
		return ""
	}
	pairs := make([]string, 0, nbWords)
	var i int
	for i = 0; i < nbWords; i++ {
		pairs = append(pairs, fake.Word()+"="+fake.Word())
	}
	return strings.Join(pairs, "&")
}
