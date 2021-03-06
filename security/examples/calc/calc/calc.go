package calc

import (
	"context"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	goahttp "goa.design/goa/http"
	adder "goa.design/plugins/security/examples/calc/adder/gen/adder"
	addercli "goa.design/plugins/security/examples/calc/adder/gen/http/adder/client"
	"goa.design/plugins/security/examples/calc/calc/gen/calc"
)

// calcs is a calc service implementation.
type calcs struct {
	logger *log.Logger
	adderc *adder.Client
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger, scheme, host string) calcsvc.Service {
	// Create adder service client using default HTTP client, encoder and
	// decoder.
	adderc := addercli.NewClient(scheme, host, http.DefaultClient, goahttp.RequestEncoder, goahttp.ResponseDecoder, false)
	return &calcs{logger: logger, adderc: adder.NewClient(adderc.Add())}
}

// Login creates a valid JWT given valid credentials. Login returns an error of
// type calcsvc.Unauthorized if the credentials are invalid.
func (s *calcs) Login(ctx context.Context, p *calcsvc.LoginPayload) (string, error) {
	// create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"iat":    time.Now().Unix(),
		"scopes": []string{"calc:add"},
	})

	s.logger.Printf("user '%s' logged in", p.User)

	// note that if "SignedString" returns an error then it is returned as
	// an internal error to the client
	return token.SignedString(Key)
}

// Add calls the adder service 'Add' endpoint. This endpoint is secured with the
// JWT scheme.
func (s *calcs) Add(ctx context.Context, p *calcsvc.AddPayload) (int, error) {
	return s.adderc.Add(ctx, &adder.AddPayload{A: p.A, B: p.B, Key: p.Token})
}
