package main

import "math/rand"

var Extensions = []string{
	"",
	"html",
	"gif",
	"png",
	"js",
	"jpg",
}

func FakeExtension(r *rand.Rand) string {
	return Extensions[r.Intn(len(Extensions))]
}
