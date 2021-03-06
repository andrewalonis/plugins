package security

import (
	"context"
	"fmt"
	"strings"
)

type (
	// BasicAuthScheme represents the BasicAuth`security scheme.
	// It consists of a simple username and password.
	BasicAuthScheme struct {
		// Name is the scheme name defined in the design.
		Name string
	}

	// APIKeyScheme represents the API key security scheme.
	// It consists of a key which is used in authentication.
	APIKeyScheme struct {
		// Name is the scheme name defined in the design.
		Name string
	}

	// JWTScheme represents an API key based scheme with support
	// for scopes.
	JWTScheme struct {
		// Name is the scheme name defined in the design.
		Name string
		// Scopes holds a list of scopes for the scheme.
		Scopes []string
		// RequiredScopes holds a list of scopes which are required
		// by the scheme. It is a subset of Scopes field.
		RequiredScopes []string
	}

	// OAuth2Scheme represents the oauth2 security scheme.
	OAuth2Scheme struct {
		// Name is the scheme name defined in the design.
		Name string
		// Scopes holds a list of scopes for the scheme.
		Scopes []string
		// RequiredScopes holds a list of scopes which are required
		// by the scheme. It is a subset of Scopes field.
		RequiredScopes []string
		// Flows determine the oauth2 flows.
		Flows []*OAuthFlow
	}

	// OAuthFlow represents the OAuth2 flow defined by the scheme.
	OAuthFlow struct {
		// Type is the type of grant.
		Type string
		// AuthorizationURL to be used for implicit or authorizationCode flows.
		AuthorizationURL string
		// TokenURL to be used for password, clientCredentials or authorizationCode flows.
		TokenURL string
		// RefreshURL to be used for obtaining refresh token.
		RefreshURL string
	}

	// AuthorizeBasicAuthFunc is the function type that implements the basic auth
	// scheme of using username and password.
	AuthorizeBasicAuthFunc func(ctx context.Context, user, pass string, s *BasicAuthScheme) (context.Context, error)

	// AuthorizeAPIKeyFunc is the function type that implements the API key
	// scheme of using an API key.
	AuthorizeAPIKeyFunc func(ctx context.Context, key string, s *APIKeyScheme) (context.Context, error)

	// AuthorizeOAuth2Func is the function type that implements the OAuth2
	// scheme of using an OAuth2 token.
	AuthorizeOAuth2Func func(ctx context.Context, token string, s *OAuth2Scheme) (context.Context, error)

	// AuthorizeJWTFunc is the function type that implements the JWT
	// scheme of using a JWT token.
	AuthorizeJWTFunc func(ctx context.Context, token string, s *JWTScheme) (context.Context, error)
)

// Validate returns a non-nil error if scopes does not contain all of
// OAuth2 scheme's required scopes.
func (s *OAuth2Scheme) Validate(scopes []string) error {
	return validateScopes(s.RequiredScopes, scopes)
}

// Validate returns a non-nil error if scopes does not contain all of
// JWT scheme's required scopes.
func (s *JWTScheme) Validate(scopes []string) error {
	return validateScopes(s.RequiredScopes, scopes)
}

func validateScopes(reqdScopes, scopes []string) error {
	var missing []string
	fmt.Println(reqdScopes, scopes)
	for _, r := range reqdScopes {
		found := false
		for _, s := range scopes {
			if s == r {
				found = true
				break
			}
		}
		if !found {
			missing = append(missing, r)
		}
	}
	if len(missing) == 0 {
		return nil
	}
	return fmt.Errorf("missing sopes: %s", strings.Join(missing, ", "))
}
