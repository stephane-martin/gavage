package main

import "math/rand"

var ExceptionIDCorpus = []string{
	"authentication_failed",
	"policy_denied",
	"tcp_error",
	"configuration_error",
	"internal_error",
	"dns_server_failure",
	"dns_unresolved_hostname",
	"authentication_failed_password_expired",
	"invalid_request",
	"authentication_redirect_from_virtual_host",
}

func FakeExceptionID(r *rand.Rand) string {
	return ExceptionIDCorpus[r.Intn(len(ExceptionIDCorpus))]
}
