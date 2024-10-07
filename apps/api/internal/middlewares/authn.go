package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/SocBongDev/soc-bong/internal/logger"
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

	issuerURL, err := url.Parse(fmt.Sprintf("https://%s/", config.Domain))
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
		ctx := c.UserContext()
		if config.Next != nil && config.Next(c) {
			return c.Next()
		}

		authHeader := c.Get("Authorization")
		if authHeaderParts := strings.Fields(authHeader); len(authHeaderParts) > 0 && strings.ToLower(authHeaderParts[0]) != "bearer" {
			logger.ErrorContext(ctx, "authn.ValidateJWT err", "err", invalidJWTErrorMessage)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": invalidJWTErrorMessage})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			logger.ErrorContext(ctx, "authn.ValidateJWT err", "err", missingJWTErrorMessage, "authHeader", authHeader)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": missingJWTErrorMessage})
		}

		token, err := jwtValidator.ValidateToken(ctx, tokenString)
		if err != nil {
			logger.ErrorContext(ctx, "Encountered error while validating JWT", "err", err, "tokenString", tokenString)
			if errors.Is(err, jwtmiddleware.ErrJWTMissing) {
				logger.ErrorContext(ctx, "authn.ValidateJWT err", "err", missingJWTErrorMessage, "tokenString", tokenString)
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": missingJWTErrorMessage})
			}
			if errors.Is(err, jwtmiddleware.ErrJWTInvalid) {
				logger.ErrorContext(ctx, "authn.ValidateJWT err", "err", invalidJWTErrorMessage, "token", token)
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": invalidJWTErrorMessage})
			}
			return fiber.ErrInternalServerError
		}

		validatedClaims, ok := token.(*validator.ValidatedClaims)
		if !ok {
			logger.ErrorContext(ctx, "authn.ValidateJWT err", "err", "Invalid token claims", "token", token)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token claims"})
		}

		enhancedClaims := &EnhancedValidatedClaims{
			ValidatedClaims: validatedClaims,
		}

		if customClaims, ok := validatedClaims.CustomClaims.(*CustomClaims); ok {
			enhancedClaims.Role = customClaims.UserRoles
		}

		c.Locals(jwtmiddleware.ContextKey{}, enhancedClaims)

		return c.Next()
	}
}
