package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @InsertUser godoc
// @Summary Create user api
// @Description Insert user
// @Tags User
// @Accept json
// @Param post body UserInput true "Create user body"
// @Success 200 {object} User
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /users [post]
func (h *UserHandler) Insert(c *fiber.Ctx) error {
	// expected data from the user
	ctx, input := c.UserContext(), new(UserInput)
	if err := c.BodyParser(&input); err != nil {
		logger.ErrorContext(ctx, "InsertUser.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "User request", "req", input)
	pwd := password{}
	hash, err := pwd.Set(input.Password)
	if err != nil {
		logger.ErrorContext(ctx, "Password hashing err", "err", err)
		return fiber.ErrInternalServerError
	}

	logger.InfoContext(ctx, "User password hash", "req", hash)
	req := &User{UserInput: UserInput{
		Email:       input.Email,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Password:    hash,
		IsActive:    input.IsActive,
		VerifyEmail: input.VerifyEmail,
		Connection:  input.Connection,
		PhoneNumber: input.PhoneNumber,
		BirthDate:   input.BirthDate,
		AgencyId:    input.AgencyId,
	}}
	// validator

	v := common.New()
	if ValidateUser(v, req); !v.Valid() {
		logger.ErrorContext(ctx, "Validate User error. Response", "req", v)
		return fiber.ErrBadRequest
	}

	// Create user in Auth0
	auth0User, err := h.createAuth0User(ctx, req, input.Password)
	if err != nil {
		logger.ErrorContext(ctx, "Auth0 user creation err", "err", err)
		// delete user in database if createAuth0User error
		return fiber.ErrConflict
	}
	auth0UserID, ok := auth0User["user_id"].(string)

	if !ok {
		logger.ErrorContext(ctx, "Failed to get Auth0 user ID from response")
		return fiber.ErrInternalServerError
	} else {
		req.Auth0UserId = auth0UserID
		if err := h.repo.Insert(ctx, req); err != nil {
			logger.ErrorContext(ctx, "User insertion err", "err", err)
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				return fiber.ErrUnprocessableEntity
			}
			return fiber.ErrInternalServerError
		}
	}
	// Respond with success
	return c.JSON(fiber.Map{
		"user":          auth0User,
		"auth0_user_id": auth0UserID,
	})
}

func (h *UserHandler) createAuth0User(ctx context.Context, user *User, password string) (map[string]interface{}, error) {
	auth0User := map[string]any{
		"email":       user.Email,
		"given_name":  user.FirstName,
		"family_name": user.LastName,
		"name":        user.FirstName + " " + user.LastName,
		"connection":  "Username-Password-Authentication",
		"password":    password,
	}

	payload, err := json.Marshal(auth0User)
	if err != nil {
		logger.ErrorContext(ctx, "createAuth0User Marshal err", "err", err)
		return nil, err
	}

	var js json.RawMessage
	if err := json.Unmarshal(payload, &js); err != nil {
		logger.ErrorContext(ctx, "createAuth0User Unmarshal err", "err", err)
		return nil, fmt.Errorf("invalid JSON payload: %v", err)
	}

	logger.InfoContext(ctx, "createAuth0User Validated Auth0 request payload", "payload", string(payload))
	auth0Domain := h.config.Domain
	if !strings.HasPrefix(auth0Domain, "https://") {
		auth0Domain = "https://" + auth0Domain
	}
	if !strings.HasSuffix(auth0Domain, "/") {
		auth0Domain += "/"
	}

	url := auth0Domain + "api/v2/users"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		logger.ErrorContext(ctx, "NewRequest err", "err", err)
		return nil, err
	}

	// Get token using the token manager
	token, err := h.tokenManager.GetToken()
	if err != nil {
		logger.ErrorContext(ctx, "GetToken err", "err", err)
		return nil, err
	}
	// Set up the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	logger.InfoContext(ctx, "createAuth0User Auth0 request", "URL", fmt.Sprintf("%sapi/v2/users", auth0Domain), "payload", string(payload))
	// HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 30, // appropriate timeout
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.ErrorContext(ctx, "createAuth0User client.Do err", "err", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorContext(ctx, "createAuth0User ReadAll err", "err", err)
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
			logger.ErrorContext(ctx, "createAuth0User Unmarshal err", "err", err)
			return nil, fmt.Errorf("failed to create Auth0 user. Status: %d, Body: %s", resp.StatusCode, string(body))
		}
		logger.ErrorContext(ctx, "Auth0 error: %+v", errorResponse)
		return nil, fmt.Errorf("failed to create Auth0 user: %s", errorResponse.Message)
	}

	// Parse the response
	var createdUser map[string]any
	if err := json.Unmarshal(body, &createdUser); err != nil {
		return nil, fmt.Errorf("error decoding Auth0 response: %v", err)
	}

	if createdUser == nil {
		return nil, fmt.Errorf("no user created in Auth0 response")
	}

	return createdUser, nil
}
