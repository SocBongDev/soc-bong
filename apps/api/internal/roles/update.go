package roles

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @UpdateRole godoc
// @Summary Update role api
// @Description Update role
// @Tags Role
// @Accept json
// @Param post body WriteRoleRequest true "Update role body"
// @Param id path string true "Role ID"
// @Success 200 {object} Role
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /roles/{id} [put]
func (h *RoleHandler) Update(c *fiber.Ctx) error {
	ctx, body := c.UserContext(), new(WriteRoleRequest)

	if err := c.BodyParser(body); err != nil {
		logger.ErrorContext(ctx, "UpdateRole.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	id := c.Params("id")

	if id == "" {
		logger.ErrorContext(ctx, "UpdateRole.Params empty")
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "UpdateRole.Request request", "req", body)
	req := &Role{WriteRoleRequest: *body, BaseEntity: BaseEntity{Id: id}}

	resp, err := h.updateUserRole(ctx, req)
	if err != nil {
		logger.InfoContext(ctx, "UpdateRole.Update err", "err", err)
		return fiber.ErrBadRequest
	}
	logger.ErrorContext(ctx, "UpdateRole.Success Response", "req", resp)
	return c.JSON(resp)
}

func (h *RoleHandler) updateUserRole(ctx context.Context, req *Role) (map[string]interface{}, error) {
	auth0Role := map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
	}

	payload, err := json.Marshal(auth0Role)
	if err != nil {
		logger.ErrorContext(ctx, "UpdateRole.UpdateAuth0Role marshaling data err", "err", err)
		return nil, err
	}

	var js json.RawMessage

	if err := json.Unmarshal(payload, &js); err != nil {
		logger.ErrorContext(ctx, "UpdateRole.UpdateAuth0Role invalid JSON payload err", "err", err)
		return nil, fmt.Errorf("invalid JSON payload: %v", err)
	}

	auth0Domain := h.config.Domain
	if !strings.HasPrefix(auth0Domain, "https://") {
		auth0Domain = "https://" + auth0Domain
	}
	if !strings.HasSuffix(auth0Domain, "/") {
		auth0Domain += "/"
	}

	url := auth0Domain + "api/v2/roles/" + req.Id

	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payload))

	if err != nil {
		logger.ErrorContext(ctx, "UpdateRole.UpdateAuth0Role Request err", "err", err)
		return nil, err
	}

	//Get token using the token manager
	token, err := h.tokenManager.GetToken()
	if err != nil {
		logger.ErrorContext(ctx, "UpdateRole.UpdateAuth0Role.GetManagementAPIToken err", "err", err)
		return nil, err
	}
	// Set up the request
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	// request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	logger.InfoContext(ctx, "UpdateRole.UpdateAuth0Role Auth0 request", "URL", fmt.Sprintf("%sapi/v2/roles/:id", auth0Domain), "payload", string(payload))
	//HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 30, //appropriate timeout
	}

	resp, err := client.Do(request)
	if err != nil {
		logger.ErrorContext(ctx, "UpdateRole.UpdateAuth0Role client.Do err", "err", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorContext(ctx, "UpdateRole.UpdateAuth0Role ReadAll Response Body err", "err", err)
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
			logger.ErrorContext(ctx, "UpdateRole.UpdateAuth0Role Unmarshal Response err", "err", err)
			return nil, fmt.Errorf("failed to update Auth0 role. Status: %d, Body: %s", resp.StatusCode, string(body))
		}
		logger.ErrorContext(ctx, "UpdateRole.UpdateAuth0Role err", "err", errorResponse)
		return nil, fmt.Errorf("failed to update Auth0 role: %s", errorResponse.Message)
	}

	// Parse the response
	var updateRole map[string]interface{}
	if err := json.Unmarshal(body, &updateRole); err != nil {
		logger.ErrorContext(ctx, "UpdateRole.UpdateAuth0Role Unmarshal Response err", "err", err)
		return nil, fmt.Errorf("error decoding Auth0 response: %v", err)
	}

	if updateRole == nil {
		logger.ErrorContext(ctx, "UpdateRole.UpdateAuth0Role no role created in Auth0 response")
		return nil, fmt.Errorf("no role created in Auth0 response")
	}

	return updateRole, nil
}
