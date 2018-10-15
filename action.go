package main

import "math/rand"

var Action = []string{
	"TCP_HIT",
	"TCP_TUNNELED",
	"TCP_DENIED",
	"TCP_NC_MISS",
	"TCP_ERR_MISS",
	"TCP_MISS",
	"TCP_REFRESH_MISS",
	"TCP_NC_MISS_RST",
	"TCP_CLIENT_REFRESH",
	"TCP_AUTH_MISS",
	"TCP_PARTIAL_MISS",
	"TCP_AUTH_HIT",
	"TCP_AUTH_REDIRECT",
}

func FakeAction(r *rand.Rand) string {
	actionIdx := int64(r.ExpFloat64())
	if actionIdx >= int64(len(Action)) {
		actionIdx = int64(len(Action) - 1)
	}
	return Action[actionIdx]
}
