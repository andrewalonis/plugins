// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// calc client
//
// Command:
// $ goa gen goa.design/plugins/security/examples/calc/calc/design

package calcsvc

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "calc" service client.
type Client struct {
	LoginEndpoint goa.Endpoint
	AddEndpoint   goa.Endpoint
}

// NewClient initializes a "calc" service client given the endpoints.
func NewClient(login, add goa.Endpoint) *Client {
	return &Client{
		LoginEndpoint: login,
		AddEndpoint:   add,
	}
}

// Login calls the "login" endpoint of the "calc" service.
// Login can return the following error types:
//	- *Unauthorized
//	- error: generic transport error.
func (c *Client) Login(ctx context.Context, p *LoginPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.LoginEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Add calls the "add" endpoint of the "calc" service.
func (c *Client) Add(ctx context.Context, p *AddPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}
