package roles

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// @FindRole godoc
// @Summary Get list role details api
// @Description Get list role
// @Tags Role
// @Accept json
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param  sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param  search query string false "Search term"
// @Success 200 {object} FindRoleResp
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /roles [get]
func (h *RoleHandler) Find(c *fiber.Ctx) error {
	query := &RoleQuery{}
	if err := c.QueryParser(query); err != nil {
		log.Println("FindRole.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("FindRoles request: %+v\n", query)

	data, err := h.findRole(query)
	if err != nil {
		log.Println("FindRole.All err: ", err)
		return fiber.ErrInternalServerError
	}

	resp := FindRoleResp{Data: data, Page: query.GetPage(), PageSize: query.GetPageSize()}
	log.Printf("FindRoles success. Response: %+v\n", resp)
	return c.JSON(resp)
}

func (h *RoleHandler) findRole(query *RoleQuery) ([]Role, error) {
	auth0Domain := h.config.Domain
	if !strings.HasPrefix(auth0Domain, "https://") {
		auth0Domain = "https://" + auth0Domain
	}
	if !strings.HasSuffix(auth0Domain, "/") {
		auth0Domain += "/"
	}

	url := auth0Domain + "api/v2/roles"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error get request: %+v\n", err)
		return nil, err
	}

	// Add query parameters
	q := req.URL.Query()
	if query.Search != "" {
		q.Add("name_filter", query.Search)
	}
	if query.Page > 0 {
		q.Add("page", strconv.Itoa(int(query.Page)))
	}
	if query.PageSize > 0 {
		q.Add("per_page", strconv.Itoa(int(query.PageSize)))
	}
	req.URL.RawQuery = q.Encode()

	token, err := h.tokenManager.GetToken()
	if err != nil {
		log.Printf("err getting Auth0 token: %v", err)
		return nil, err
	}

	log.Printf("check token: %+v\n", token)

	req.Header.Set("Content-Type", "application/json charset=utf-8")
	// req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	log.Printf("Auth0 request URL: %s", auth0Domain+"api/v2/roles")

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

	log.Printf("Check body after get Roles from Auth0: %+v\n", body)

	// Parse the JSON response
	var roles []Role
	err = json.Unmarshal(body, &roles)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
