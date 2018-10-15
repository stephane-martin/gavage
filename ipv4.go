package main

import (
	"math/rand"
	"net"
)

func FakeClientIP(r *rand.Rand) string {
	size := 4
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(r.Intn(256))
	}
	return net.IP(ip).To4().String()
}
