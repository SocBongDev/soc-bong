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
	body := new(WriteRoleRequest)

	if err := c.BodyParser(body); err != nil {
		log.Println("UpdateRole.BodyParser error", err)
		return fiber.ErrBadRequest
	}

	id := c.Params("id")

	if id == "" {
		log.Println("UpdateRole.Params err: ", id)
		return fiber.ErrBadRequest
	}

	log.Printf("UpdateRole request: %+v\n", body)
	req := &Role{WriteRoleRequest: *body, BaseEntity: BaseEntity{Id: id}}

	resp, err := h.updateUserRole(req)
	if err != nil {
		log.Println("Auth0 role update error: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("update role successfully: %+v\n", resp)
	return c.JSON(resp)
}

func (h *RoleHandler) updateUserRole(req *Role) (map[string]interface{}, error) {
	auth0Role := map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
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

	url := auth0Domain + "api/v2/roles/" + req.Id

	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payload))

	log.Println("check request: ", request)

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
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	// request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	log.Printf("Auth0 request URL: %s", auth0Domain+"api/v2/roles/"+req.Id)
	log.Printf("Auth0 request payload: %s", string(payload))

	//HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 30, //appropriate timeout
	}

	resp, err := client.Do(request)
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
			return nil, fmt.Errorf("failed to update Auth0 role. Status: %d, Body: %s", resp.StatusCode, string(body))
		}
		log.Printf("Auth0 error: %+v", errorResponse)
		return nil, fmt.Errorf("failed to update Auth0 role: %s", errorResponse.Message)
	}

	// Parse the response
	var updateRole map[string]interface{}
	if err := json.Unmarshal(body, &updateRole); err != nil {
		return nil, fmt.Errorf("error decoding Auth0 response: %v", err)
	}

	if updateRole == nil {
		return nil, fmt.Errorf("no role created in Auth0 response")
	}

	return updateRole, nil
}
