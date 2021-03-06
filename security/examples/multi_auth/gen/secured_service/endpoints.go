// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// secured_service endpoints
//
// Command:
// $ goa gen goa.design/plugins/security/examples/multi_auth/design

package securedservice

import (
	"context"

	goa "goa.design/goa"
)

// Endpoints wraps the "secured_service" service endpoints.
type Endpoints struct {
	Signin           goa.Endpoint
	Secure           goa.Endpoint
	DoublySecure     goa.Endpoint
	AlsoDoublySecure goa.Endpoint
}

// NewEndpoints wraps the methods of the "secured_service" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Signin:           NewSigninEndpoint(s),
		Secure:           NewSecureEndpoint(s),
		DoublySecure:     NewDoublySecureEndpoint(s),
		AlsoDoublySecure: NewAlsoDoublySecureEndpoint(s),
	}
}

// Use applies the given middleware to all the "secured_service" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Signin = m(e.Signin)
	e.Secure = m(e.Secure)
	e.DoublySecure = m(e.DoublySecure)
	e.AlsoDoublySecure = m(e.AlsoDoublySecure)
}

// NewSigninEndpoint returns an endpoint function that calls the method
// "signin" of service "secured_service".
func NewSigninEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SigninPayload)
		return s.Signin(ctx, p)
	}
}

// NewSecureEndpoint returns an endpoint function that calls the method
// "secure" of service "secured_service".
func NewSecureEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SecurePayload)
		return s.Secure(ctx, p)
	}
}

// NewDoublySecureEndpoint returns an endpoint function that calls the method
// "doubly_secure" of service "secured_service".
func NewDoublySecureEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DoublySecurePayload)
		return s.DoublySecure(ctx, p)
	}
}

// NewAlsoDoublySecureEndpoint returns an endpoint function that calls the
// method "also_doubly_secure" of service "secured_service".
func NewAlsoDoublySecureEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AlsoDoublySecurePayload)
		return s.AlsoDoublySecure(ctx, p)
	}
}
