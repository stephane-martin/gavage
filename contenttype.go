package main

import "math/rand"

var ContentType = []string{
	"image/jpeg",
	"image/gif",
	"text/html",
	"image/png",
	"text/plain",
	"text/xml",
	"application/xml",
	"application/javascript",
	"text/css",
}

func FakeContentType(r *rand.Rand) string {
	return ContentType[r.Intn(len(ContentType))]
}
