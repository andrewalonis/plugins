// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// fetcher HTTP server
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/fetcher/fetcher/design

package server

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
	fetchersvc "goa.design/plugins/goakit/examples/fetcher/fetcher/gen/fetcher"
)

// Server lists the fetcher service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Fetch  http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the fetcher service endpoints.
func New(
	e *fetchersvc.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Fetch", "GET", "/fetch/{*url}"},
		},
		Fetch: NewFetchHandler(e.Fetch, mux, dec, enc),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "fetcher" }

// Mount configures the mux to serve the fetcher endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountFetchHandler(mux, h.Fetch)
}

// MountFetchHandler configures the mux to serve the "fetcher" service "fetch"
// endpoint.
func MountFetchHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/fetch/{*url}", f)
}

// NewFetchHandler creates a HTTP handler which loads the HTTP request and
// calls the "fetcher" service "fetch" endpoint.
func NewFetchHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
) http.Handler {
	var (
		decodeRequest  = DecodeFetchRequest(mux, dec)
		encodeResponse = EncodeFetchResponse(enc)
		encodeError    = EncodeFetchError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept")
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, accept)
		ctx = context.WithValue(ctx, goa.MethodKey, "fetch")
		ctx = context.WithValue(ctx, goa.ServiceKey, "fetcher")
		payload, err := decodeRequest(r)
		if err != nil {
			encodeError(ctx, w, err)
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			encodeError(ctx, w, err)
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			encodeError(ctx, w, err)
		}
	})
}
