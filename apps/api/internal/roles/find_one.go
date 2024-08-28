package roles

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

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
	id := c.Params("id")
	if id == "" {
		log.Println("GetRoleDetails.Param empty")
		return fiber.ErrBadRequest
	}

	log.Println("GetRoleDetails id: ", id)

	resp, err := h.findOneRole(id)

	if err != nil {
		log.Println("GetRoleDetails err: ", err)
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	log.Printf("GetRoleDetails success. Response: %+v\n", resp)

	return c.JSON(resp)

}

func (h *RoleHandler) findOneRole(id string) (*Role, error) {
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
		log.Printf("Error get request: %+v\n", err)
		return nil, err
	}

	token, err := h.tokenManager.GetToken()
	if err != nil {
		log.Printf("err getting Auth0 token: %v", err)
		return nil, err
	}

	log.Printf("check token: %+v\n", token)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	log.Printf("Auth0 request URL: %s", auth0Domain+"api/v2/roles/:id")

	//HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 30, //appropriate timeout
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	log.Printf("check this resp: %+v\n", resp)
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode >= 400 {
		// Read the error response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
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
			log.Printf("Error parsing error response: %v", err)
			return nil, err
		}

		// Return the error
		return nil, fmt.Errorf("Auth0 API error: [%d] %s: %s", errorResp.StatusCode, errorResp.Error, errorResp.Message)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}

	log.Printf("Check body after get a role from Auth0: %+v\n", body)

	// Parse the JSON response
	var role Role
	err = json.Unmarshal(body, &role)
	if err != nil {
		log.Println("Can't unmarshal json: ", err)
		return nil, err
	}

	return &role, nil
}
