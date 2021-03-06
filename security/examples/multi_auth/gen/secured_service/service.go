// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// secured_service service
//
// Command:
// $ goa gen goa.design/plugins/security/examples/multi_auth/design

package securedservice

import (
	"context"
)

// The secured service exposes endpoints that require valid authorization
// credentials.
type Service interface {
	// Creates a valid JWT
	Signin(context.Context, *SigninPayload) (string, error)
	// This action is secured with the jwt scheme
	Secure(context.Context, *SecurePayload) (string, error)
	// This action is secured with the jwt scheme and also requires an API key
	// query string.
	DoublySecure(context.Context, *DoublySecurePayload) (string, error)
	// This action is secured with the jwt scheme and also requires an API key
	// header.
	AlsoDoublySecure(context.Context, *AlsoDoublySecurePayload) (string, error)
}

// Credentials used to authenticate to retrieve JWT token
type SigninPayload struct {
	// Username used to perform signin
	Username *string
	// Username used to perform signin
	Password *string
}

// SecurePayload is the payload type of the secured_service service secure
// method.
type SecurePayload struct {
	// Whether to force auth failure even with a valid JWT
	Fail *bool
	// JWT used for authentication
	Token *string
}

// DoublySecurePayload is the payload type of the secured_service service
// doubly_secure method.
type DoublySecurePayload struct {
	// API key
	Key *string
	// JWT used for authentication
	Token *string
}

// AlsoDoublySecurePayload is the payload type of the secured_service service
// also_doubly_secure method.
type AlsoDoublySecurePayload struct {
	// Username used to perform signin
	Username *string
	// Username used to perform signin
	Password *string
	// API key
	Key *string
	// JWT used for authentication
	Token      *string
	OauthToken *string
}

type Unauthorized struct {
	// Credentials are invalid
	Value string
}

// Error returns "unauthorized".
func (e *Unauthorized) Error() string {
	return "unauthorized"
}
