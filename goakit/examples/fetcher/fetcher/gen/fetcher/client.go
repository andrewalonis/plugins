// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// fetcher client
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/fetcher/fetcher/design

package fetchersvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Client is the "fetcher" service client.
type Client struct {
	FetchEndpoint endpoint.Endpoint
}

// NewClient initializes a "fetcher" service client given the endpoints.
func NewClient(fetch endpoint.Endpoint) *Client {
	return &Client{
		FetchEndpoint: fetch,
	}
}

// Fetch calls the "fetch" endpoint of the "fetcher" service.
// Fetch can return the following error types:
//	- *Error
//	- *Error
//	- error: generic transport error.
func (c *Client) Fetch(ctx context.Context, p *FetchPayload) (res *FetchMedia, err error) {
	var ires interface{}
	ires, err = c.FetchEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*FetchMedia), nil
}
