package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	goahttp "goa.design/goa/http"
	"goa.design/goa/http/middleware"
	calc "goa.design/plugins/security/examples/calc/calc"
	calcsvc "goa.design/plugins/security/examples/calc/calc/gen/calc"
	calcsvcsvr "goa.design/plugins/security/examples/calc/calc/gen/http/calc/server"
)

func main() {
	// Define command line flags, add any other flag required to configure
	// the service.
	var (
		adderAddr = flag.String("adder-addr", ":8000", "Adder service HTTP listen `address`")
		addr      = flag.String("listen", ":8080", "HTTP listen `address`")
		dbg       = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger and goa log adapter. Replace logger with your own using
	// your log package of choice. The goa.design/middleware/logging/...
	// packages define log adapters for common log packages.
	var (
		adapter middleware.Logger
		logger  *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[calc] ", log.Ltime)
		adapter = middleware.NewLogger(logger)
	}

	// Create the structs that implement the services.
	var (
		calcsvcSvc calcsvc.Service
	)
	{
		calcsvcSvc = calc.NewCalc(logger, "http", *adderAddr)
	}

	// Wrap the services in endpoints that can be invoked from other
	// services potentially running in different processes.
	var (
		calcsvcEndpoints *calcsvc.Endpoints
	)
	{
		calcsvcEndpoints = calcsvc.NewSecureEndpoints(calcsvcSvc, calc.BasicAuthFunc, calc.JWTAuthFunc)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		calcsvcServer *calcsvcsvr.Server
	)
	{
		calcsvcServer = calcsvcsvr.New(calcsvcEndpoints, mux, dec, enc)
	}

	// Configure the mux.
	calcsvcsvr.Mount(mux, calcsvcServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		if *dbg {
			handler = middleware.Debug(mux, os.Stdout)(handler)
		}
		handler = middleware.Log(adapter)(handler)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the service to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: *addr, Handler: handler}
	go func() {
		for _, m := range calcsvcServer.Mounts {
			logger.Printf("[calc] service %q method %q mounted on %s %s", calcsvcServer.Service(), m.Method, m.Verb, m.Pattern)
		}
		logger.Printf("[calc] listening on %s", *addr)
		errc <- srv.ListenAndServe()
	}()

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Shutdown gracefully with a 30s timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)

	logger.Println("exited")
}
