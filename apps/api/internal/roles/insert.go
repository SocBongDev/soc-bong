package roles

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @InsertRole godoc
// @Summary Create role api
// @Description Insert role
// @Tags Role
// @Accept json
// @Param post body WriteRoleRequest true "Create role body"
// @Success 200 {object} Role
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /roles [post]
func (h *RoleHandler) Insert(c *fiber.Ctx) error {
	ctx, body := c.UserContext(), new(WriteRoleRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("InsertRole.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}
	logger.InfoContext(ctx, "InsertRole request", "req", body)
	req := &Role{WriteRoleRequest: *body}
	role, err := h.createRole(ctx, req)
	if err != nil {
		logger.ErrorContext(ctx, "InsertRole.CreateAuth0Role err", "err", err)
		return fiber.ErrConflict
	}

	logger.InfoContext(ctx, "InsertRole.CreateAuth0Role Success Response", "res", role)
	return c.JSON(role)
}

func (h *RoleHandler) createRole(ctx context.Context, role *Role) (map[string]interface{}, error) {
	auth0Role := map[string]interface{}{
		"name":        role.Name,
		"description": role.Description,
	}

	payload, err := json.Marshal(auth0Role)
	if err != nil {
		logger.ErrorContext(ctx, "InsertRole.CreateAuth0Role Marshal err", "err", err)
		return nil, err
	}

	var js json.RawMessage

	if err := json.Unmarshal(payload, &js); err != nil {
		logger.ErrorContext(ctx, "InsertRole.CreateAuth0Role Unmarshal err", "err", err)
		return nil, fmt.Errorf("invalid JSON payload: %v", err)
	}

	logger.InfoContext(ctx, "InsertRole.CreateAuth0Role Validated Auth0 request payload", "payload", string(payload))

	auth0Domain := h.config.Domain
	if !strings.HasPrefix(auth0Domain, "https://") {
		auth0Domain = "https://" + auth0Domain
	}
	if !strings.HasSuffix(auth0Domain, "/") {
		auth0Domain += "/"
	}

	url := auth0Domain + "api/v2/roles"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		logger.ErrorContext(ctx, "InsertRole.CreateAuth0Role req err", "err", err)
		return nil, err
	}

	//Get token using the token manager
	token, err := h.tokenManager.GetToken()
	if err != nil {
		logger.ErrorContext(ctx, "InsertRole.GetManagementAPIToken err", "err", err)
		return nil, err
	}
	// Set up the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	logger.InfoContext(ctx, "InsertRole.CreateAuth0Role Auth0 request", "URL", fmt.Sprintf("%sapi/v2/roles", auth0Domain), "payload", string(payload))

	//HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 30, //appropriate timeout
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.ErrorContext(ctx, "InsertRole.CreateAuth0Role client.Do err", "err", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorContext(ctx, "InsertRole.CreateAuth0Role ReadAll err", "err", err)
		return nil, err
	}

	// Check the response
	if resp.StatusCode != 201 && resp.StatusCode != 200 && resp.StatusCode != 204 {
		var errorResponse struct {
			StatusCode int    `json:"statusCode"`
			Error      string `json:"error"`
			Message    string `json:"message"`
		}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			logger.ErrorContext(ctx, "InsertRole.CreateAuth0Role Unmarshal Response err", "err", err)
			return nil, fmt.Errorf("failed to create Auth0 role. Status: %d, Body: %s", resp.StatusCode, string(body))
		}
		logger.ErrorContext(ctx, "InsertRole.CreateAuth0Role error: %+v", errorResponse)
		return nil, fmt.Errorf("failed to create Auth0 role: %s", errorResponse.Message)
	}

	// Parse the response
	var createRole map[string]interface{}
	if err := json.Unmarshal(body, &createRole); err != nil {
		return nil, fmt.Errorf("error decoding Auth0 response: %v", err)
	}

	if createRole == nil {
		return nil, fmt.Errorf("no role created in Auth0 response")
	}

	return createRole, nil
}
