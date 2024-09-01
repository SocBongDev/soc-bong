package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @InsertUserBySignUp godoc
// @Summary Create user by sign up api
// @Description Insert user by sign up
// @Tags SignUpInTheUserRoute
// @Accept json
// @Param post body UserInput true "Create user sign up body"
// @Success 200 {object} User
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Router /sign-up [post]
func (h *UserHandler) SignUp(c *fiber.Ctx) error {
	// expected data from the user
	ctx, input := c.UserContext(), new(UserInput)

	if err := c.BodyParser(&input); err != nil {
		log.Println("InsertUser.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("User request: %+v\n", input)

	pwd := password{}
	hash, err := pwd.Set(input.Password)
	if err != nil {
		log.Println("Password hashing error: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("User password hash: %+v\n", hash)

	req := &User{UserInput: UserInput{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  hash,

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
		log.Printf("Validate User error. Response: %+v\n", v)
		return fiber.ErrBadRequest
	}

	// Create user in Auth0

	auth0User, err := h.createSignUpAuth0User(req, input.Password)
	if err != nil {
		log.Println("Auth0 user creation error: ", err)
		// delete user in database if createAuth0User error
		return fiber.ErrConflict
	}

	auth0UserID, ok := auth0User["user_id"].(string)
	log.Printf("auth0UserID: %+v\n", auth0UserID)

	if !ok {
		log.Println("Failed to get Auth0 user ID from response")
		return fiber.ErrBadRequest
	} else {
		req.Auth0UserId = auth0UserID
		log.Printf("check this req Insert: %+v\n", req)
		if err := h.repo.Insert(ctx, req); err != nil {
			log.Println("User insertion error: ", err)
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				return fiber.ErrUnprocessableEntity
			}
			return fiber.ErrInternalServerError
		}

		// Respond with success
		return c.JSON(fiber.Map{
			"user":          auth0User,
			"auth0_user_id": auth0UserID,
		})
	}
}

func (h *UserHandler) createSignUpAuth0User(user *User, password string) (map[string]interface{}, error) {
	auth0User := map[string]interface{}{
		"email":       user.Email,
		"given_name":  user.FirstName,
		"family_name": user.LastName,
		"name":        user.FirstName + " " + user.LastName,
		"connection":  "Username-Password-Authentication",
		"password":    password,
	}

	payload, err := json.Marshal(auth0User)
	if err != nil {
		log.Printf("error marshaling Auth0 user data: %+v\n", err)
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

	url := auth0Domain + "api/v2/users"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	// Get token using the token manager
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

	// HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 30, // appropriate timeout
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
			return nil, fmt.Errorf("failed to create Auth0 user. Status: %d, Body: %s", resp.StatusCode, string(body))
		}
		log.Printf("Auth0 error: %+v", errorResponse)
		return nil, fmt.Errorf("failed to create Auth0 user: %s", errorResponse.Message)
	}

	// Parse the response
	var createdUser map[string]interface{}
	if err := json.Unmarshal(body, &createdUser); err != nil {
		return nil, fmt.Errorf("error decoding Auth0 response: %v", err)
	}

	if createdUser == nil {
		return nil, fmt.Errorf("no user created in Auth0 response")
	}

	return createdUser, nil
}
