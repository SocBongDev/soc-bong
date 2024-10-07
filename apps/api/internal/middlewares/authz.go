package middlewares

import (
	"context"
	"net/http"
	"slices"

	"github.com/SocBongDev/soc-bong/internal/logger"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gofiber/fiber/v2"
)

const (
	missingJWTErrorMessage       = "Requires authentication"
	invalidJWTErrorMessage       = "Bad credentials"
	permissionDeniedErrorMessage = "Permission denied"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type EnhancedValidatedClaims struct {
	*validator.ValidatedClaims
	Role []string `json:"user/roles,omitempty"`
}

type CustomClaims struct {
	Permissions []string `json:"permissions"`
	UserRoles   []string `json:"user/roles,omitempty"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func (c CustomClaims) HasPermissions(expectedClaims []string) bool {
	if len(expectedClaims) == 0 {
		return false
	}
	for _, scope := range expectedClaims {
		if !slices.Contains(c.Permissions, scope) {
			return false
		}
	}
	return true
}

func ValidatePermissions(expectedClaims []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		token, ok := c.Locals(jwtmiddleware.ContextKey{}).(*EnhancedValidatedClaims)
		if !ok {
			logger.ErrorContext(ctx, "authz.ValidatePermissions invalid token err", "err", "Invalid token")
			return c.Status(http.StatusUnauthorized).JSON(ErrorMessage{Message: "Invalid token"})
		}
		claims := token.CustomClaims.(*CustomClaims)
		if !claims.HasPermissions(expectedClaims) {
			errorMessage := ErrorMessage{Message: permissionDeniedErrorMessage}
			if err := c.Status(http.StatusForbidden).JSON(errorMessage); err != nil {
				logger.ErrorContext(ctx, "authz.ValidatePermissions err", "err", err)
			}
			return nil
		}
		c.Locals("role", token.Role)
		c.Locals("userId", token.RegisteredClaims.Subject)
		return c.Next()
	}
}
