// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// health HTTP server
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
	health "goa.design/plugins/goakit/examples/fetcher/fetcher/gen/health"
)

// Server lists the health service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Show   http.Handler
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

// New instantiates HTTP handlers for all the health service endpoints.
func New(
	e *health.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Show", "GET", "/health"},
		},
		Show: NewShowHandler(e.Show, mux, dec, enc),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "health" }

// Mount configures the mux to serve the health endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountShowHandler(mux, h.Show)
}

// MountShowHandler configures the mux to serve the "health" service "show"
// endpoint.
func MountShowHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/health", f)
}

// NewShowHandler creates a HTTP handler which loads the HTTP request and calls
// the "health" service "show" endpoint.
func NewShowHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
) http.Handler {
	var (
		encodeResponse = EncodeShowResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept")
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, accept)
		ctx = context.WithValue(ctx, goa.MethodKey, "show")
		ctx = context.WithValue(ctx, goa.ServiceKey, "health")
		res, err := endpoint(ctx, nil)

		if err != nil {
			encodeError(ctx, w, err)
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			encodeError(ctx, w, err)
		}
	})
}
