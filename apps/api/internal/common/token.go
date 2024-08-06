package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/SocBongDev/soc-bong/internal/config"
)

type Auth0TokenManager struct {
	token        string
	expiredAt    time.Time
	mutex        sync.RWMutex
	config       *config.Config
	clientID     string
	clientSecret string
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func NewAuth0TokenManager(config *config.Config, clientId, clientSecret string) *Auth0TokenManager {
	return &Auth0TokenManager{
		config:       config,
		clientID:     clientId,
		clientSecret: clientSecret,
	}
}

func (tm *Auth0TokenManager) GetToken() (string, error) {
	tm.mutex.RLock()
	if tm.token != "" && time.Now().Before(tm.expiredAt) {
		defer tm.mutex.RUnlock()
		return tm.token, nil
	}
	tm.mutex.RUnlock()
	return tm.refreshToken()
}

func (tm *Auth0TokenManager) refreshToken() (string, error) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	//check again in case another goroutine has already refreshed the token
	if tm.token != "" && time.Now().Before(tm.expiredAt) {
		return tm.token, nil
	}

	token, err := getManagementToken(tm.config.Domain, tm.config.ClientId, tm.clientSecret)
	if err != nil {
		return "", err
	}

	tm.token = token
	tm.expiredAt = time.Now().Add(24 * time.Hour)

	return tm.token, nil
}

func getManagementToken(domain, clientID, clientSecret string) (string, error) {
	audience := fmt.Sprintf("https://%s/api/v2/", domain)

	payload := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     clientID,
		"client_secret": clientSecret,
		"audience":      audience,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("err marshalling payload %v", err)
	}

	url := fmt.Sprintf("https://%s/oauth/token", domain)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("error reading resp body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var tokenResp TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	return tokenResp.AccessToken, nil
}
