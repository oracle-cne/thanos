// Copyright (c) 2023, Oracle and/or its affiliates.

package basicauth

import (
	"context"
	"encoding/base64"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type basicAuth struct {
	user string
	pass string
}

// basicAuth implements the credentials.PerRPCCredentials interface
var _ credentials.PerRPCCredentials = basicAuth{}

// GetRequestMetadata returns a map containing an authorization key and base64-encoded user and password
func (ba basicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	auth := ba.user + ":" + ba.pass
	return map[string]string{"authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))}, nil
}

func (ba basicAuth) RequireTransportSecurity() bool {
	// returning false here allows basic auth to be set for non-TLS endpoints, when TLS is being originated
	// by a proxy
	return false
}

// AddBasicAuthDialOption adds a basic auth HTTP header to the list of DialOption if user and pass environment variables are set
func AddBasicAuthDialOption(opts []grpc.DialOption) []grpc.DialOption {
	user := os.Getenv("CLIENT_BASIC_AUTH_USER")
	pass := os.Getenv("CLIENT_BASIC_AUTH_PASS")

	if user != "" && pass != "" {
		opt := grpc.WithPerRPCCredentials(basicAuth{user: user, pass: pass})
		opts = append(opts, opt)
	}

	return opts
}
