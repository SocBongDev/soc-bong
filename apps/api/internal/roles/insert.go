package roles

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

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
	body := new(WriteRoleRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("InsertRole.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("InsertRole request: %+v\n", body)
	req := &Role{WriteRoleRequest: *body}
	role, err := h.createRole(req)
	if err != nil {
		log.Println("Auth0 role creation error: ", err)
		return fiber.ErrConflict
	}

	log.Printf("create role successfully: %+v\n", role)

	return c.JSON(role)
}

func (h *RoleHandler) createRole(role *Role) (map[string]interface{}, error) {
	auth0Role := map[string]interface{}{
		"name":        role.Name,
		"description": role.Description,
	}

	payload, err := json.Marshal(auth0Role)
	if err != nil {
		log.Printf("error marshaling Auth0 role data: %+v\n", err)
		return nil, err
	}

	var js json.RawMessage

	if err := json.Unmarshal(payload, &js); err != nil {
		log.Printf("Invalid JSON payload %+v\n", err)
		return nil, fmt.Errorf("invalid JSON payload: %v", err)
	}

	log.Printf("Validated Auth0 request payload: %s", string(payload))

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
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	//Get token using the token manager
	token, err := h.tokenManager.GetToken()
	if err != nil {
		log.Printf("err getting Auth0 token: %v", err)
		return nil, err
	}
	// Set up the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	log.Printf("Auth0 request URL: %s", auth0Domain+"api/v2/users")
	log.Printf("Auth0 request payload: %s", string(payload))

	//HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 30, //appropriate timeout
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}

	fmt.Printf("Auth0 response status code: %d\n", resp.StatusCode)
	fmt.Printf("Auth0 response body: %s\n", string(body))

	// Check the response
	if resp.StatusCode != 201 && resp.StatusCode != 200 && resp.StatusCode != 204 {
		var errorResponse struct {
			StatusCode int    `json:"statusCode"`
			Error      string `json:"error"`
			Message    string `json:"message"`
		}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.Printf("Failed to parse error response: %v", err)
			return nil, fmt.Errorf("failed to create Auth0 role. Status: %d, Body: %s", resp.StatusCode, string(body))
		}
		log.Printf("Auth0 error: %+v", errorResponse)
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
