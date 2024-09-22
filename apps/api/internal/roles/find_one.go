package roles

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @FindOneRole godoc
// @Summary Get role details api
// @Description Get one role
// @Tags Role
// @Accept json
// @Param id path string true "Role ID"
// @Success 200 {object} Role
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /roles/{id} [get]
func (h *RoleHandler) FindOne(c *fiber.Ctx) error {
	ctx, id := c.UserContext(), c.Params("id")
	if id == "" {
		logger.ErrorContext(ctx, "GetRoleDetails.Param err", "err")
		return fiber.ErrBadRequest
	}
	logger.InfoContext(ctx, "GetRoleDetails request", "id", id)

	resp, err := h.findOneRole(ctx, id)

	if err != nil {
		logger.ErrorContext(ctx, "GetRoleDetails.Query err", "err", err)
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}
	logger.InfoContext(ctx, "GetRoleDetails.Success Response", "resp", resp)
	return c.JSON(resp)
}

func (h *RoleHandler) findOneRole(ctx context.Context, id string) (*Role, error) {
	auth0Domain := h.config.Domain
	if !strings.HasPrefix(auth0Domain, "https://") {
		auth0Domain = "https://" + auth0Domain
	}
	if !strings.HasSuffix(auth0Domain, "/") {
		auth0Domain += "/"
	}

	url := auth0Domain + "api/v2/roles/" + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.ErrorContext(ctx, "GetRoleDetails.Auth0RoleDetails err", "err", err)
		return nil, err
	}

	token, err := h.tokenManager.GetToken()
	if err != nil {
		logger.ErrorContext(ctx, "GetRoleDetails.GetManagementAPIToken err", "err", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+token)

	logger.InfoContext(ctx, "GetRoleDetails.Auth0RoleDetails Auth0 request", "URL", fmt.Sprintf("%sapi/v2/roles/:id", auth0Domain))

	//HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 30, //appropriate timeout
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.ErrorContext(ctx, "GetRoleDetails.Auth0RoleDetails client.Do err", "err", err)
		return nil, err
	}

	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode >= 400 {
		// Read the error response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.ErrorContext(ctx, "GetRoleDetails.Auth0RoleDetails ReadAll err", "err", err)
			return nil, err
		}

		// Parse the error response
		var errorResp struct {
			StatusCode int    `json:"statusCode"`
			Error      string `json:"error"`
			Message    string `json:"message"`
			ErrorCode  string `json:"errorCode"`
		}
		err = json.Unmarshal(body, &errorResp)
		if err != nil {
			logger.ErrorContext(ctx, "GetRoleDetails.Auth0RoleDetails Unmarshal Response err", "err", err)
			return nil, err
		}
		logger.ErrorContext(ctx, "GetRoleDetails.Auth0RoleDetails error", "err", errorResp)
		// Return the error
		return nil, fmt.Errorf("Auth0 API error: [%d] %s: %s", errorResp.StatusCode, errorResp.Error, errorResp.Message)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorContext(ctx, "GetRoleDetails.Auth0RoleDetails ReadAll err", "err", err)
		return nil, err
	}

	// Parse the JSON response
	var role Role
	err = json.Unmarshal(body, &role)
	if err != nil {
		logger.ErrorContext(ctx, "GetRoleDetails.Auth0RoleDetails Unmarshal Response err", "err", err)
		return nil, err
	}

	return &role, nil
}
