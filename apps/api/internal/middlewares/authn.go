package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gofiber/fiber/v2"
)

type JWTMiddlewareConfig struct {
	Audience string
	Domain   string
	Next     func(*fiber.Ctx) bool
}

type JWTMiddlewareOption func(*JWTMiddlewareConfig)

func WithNext(next func(*fiber.Ctx) bool) JWTMiddlewareOption {
	return func(config *JWTMiddlewareConfig) {
		config.Next = next
	}
}

func ValidateJWT(audience, domain string, opts ...JWTMiddlewareOption) fiber.Handler {
	config := &JWTMiddlewareConfig{
		Audience: audience,
		Domain:   domain,
		Next:     nil,
	}

	for _, opt := range opts {
		opt(config)
	}

	issuerURL, err := url.Parse("https://" + config.Domain + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{config.Audience},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return new(CustomClaims)
		}),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	return func(c *fiber.Ctx) error {
		if config.Next != nil && config.Next(c) {
			return c.Next()
		}

		authHeader := c.Get("Authorization")
		if authHeaderParts := strings.Fields(authHeader); len(authHeaderParts) > 0 && strings.ToLower(authHeaderParts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": invalidJWTErrorMessage})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": missingJWTErrorMessage})
		}

		token, err := jwtValidator.ValidateToken(c.Context(), tokenString)
		if err != nil {
			log.Printf("Encountered error while validating JWT: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": invalidJWTErrorMessage})
		}

		validatedClaims, ok := token.(*validator.ValidatedClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token claims"})
		}

		enhancedClaims := &EnhancedValidatedClaims{
			ValidatedClaims: validatedClaims,
		}

		if customClaims, ok := validatedClaims.CustomClaims.(*CustomClaims); ok {
			enhancedClaims.Role = customClaims.UserRoles
		}

		c.Locals("user", enhancedClaims)

		return c.Next()
	}
}

func validateJWT(config *JWTMiddlewareConfig, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		audience, domain := config.Audience, config.Domain
		// if config.Next != nil && config.Next(adaptor.)

		issuerURL, err := url.Parse("https://" + domain + "/")
		if err != nil {
			log.Fatalf("Failed to parse the issuer url: %v", err)
		}

		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{audience},
			validator.WithCustomClaims(func() validator.CustomClaims {
				return new(CustomClaims)
			}),
		)
		if err != nil {
			log.Fatalf("Failed to set up the jwt validator")
		}

		if authHeaderParts := strings.Fields(r.Header.Get("Authorization")); len(
			authHeaderParts,
		) > 0 &&
			strings.ToLower(authHeaderParts[0]) != "bearer" {
			errorMessage := ErrorMessage{Message: invalidJWTErrorMessage}
			if err := WriteJSON(w, http.StatusUnauthorized, errorMessage); err != nil {
				log.Printf("Failed to write error message: %v", err)
			}
			return
		}

		errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("Encountered error while validating JWT: %v", err)
			if errors.Is(err, jwtmiddleware.ErrJWTMissing) {
				errorMessage := ErrorMessage{Message: missingJWTErrorMessage}
				if err := WriteJSON(w, http.StatusUnauthorized, errorMessage); err != nil {
					log.Printf("Failed to write error message: %v", err)
				}
				return
			}
			if errors.Is(err, jwtmiddleware.ErrJWTInvalid) {
				errorMessage := ErrorMessage{Message: invalidJWTErrorMessage}
				if err := WriteJSON(w, http.StatusUnauthorized, errorMessage); err != nil {
					log.Printf("Failed to write error message: %v", err)
				}
				return
			}
			ServerError(w, err)
		}

		middleware := jwtmiddleware.New(
			func(ctx context.Context, tokenString string) (interface{}, error) {
				token, err := jwtValidator.ValidateToken(ctx, tokenString)
				if err != nil {
					return nil, err
				}

				validatedClaims, ok := token.(*validator.ValidatedClaims)

				if !ok {
					return nil, errors.New("invalid token claims")
				}

				enhancedClaims := &EnhancedValidatedClaims{
					ValidatedClaims: validatedClaims,
				}

				// Extract role from token code
				if customClaims, ok := validatedClaims.CustomClaims.(*CustomClaims); ok {
					enhancedClaims.Role = customClaims.UserRoles
				}
				return enhancedClaims, nil
			},
			jwtmiddleware.WithErrorHandler(errorHandler),
		)

		middleware.CheckJWT(next).ServeHTTP(w, r)
	})
}
