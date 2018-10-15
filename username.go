package main

import "math/rand"

func FakeUserName(r *rand.Rand) string {
	return FakeFirstName(r) + "." + FakeLastName(r)
}
