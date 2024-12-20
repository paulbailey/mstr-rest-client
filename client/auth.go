package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/paulbailey/mstr-rest-client/types"
)

type AuthService service

func (c *AuthService) Authenticate() error {
	return nil
}

func (c *AuthService) Login(ctx context.Context) error {
	if c.client.Authentication == nil {
		return fmt.Errorf("no authentication method set")
	}
	authReq := c.client.Authentication.AuthenticationRequest(c.client.ApplicationType)
	resp, respErr := c.client.DoAPIRequest(ctx, types.APIRequestInput{
		Method:  http.MethodPost,
		APIPath: "/auth/login",
		Body:    authReq,
	})
	if respErr != nil {
		return fmt.Errorf("couldn't login: %v", respErr)
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("couldn't login: %v", resp.Status)
	} else {
		// checkov:skip=CKV_SECRET_6: Not a secret
		token := resp.Header.Get("X-MSTR-AuthToken")
		c.client.AuthToken = &token
	}
	return nil
}

func (c *AuthService) Logout(ctx context.Context) error {
	resp, respErr := c.client.DoAPIRequest(ctx, types.APIRequestInput{
		Method:  http.MethodPost,
		APIPath: "/auth/logout"})
	if respErr != nil {
		return fmt.Errorf("couldn't logout: %v", respErr)
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("couldn't logout: %v", resp.Status)
	} else {
		c.client.AuthToken = nil
		return nil
	}
}

// Implement API token request
func (c *AuthService) CreateAPIToken(ctx context.Context, userID string) (*string, error) {
	body := map[string]string{
		"userId": userID,
	}
	var respBody map[string]string
	resp, respErr := c.client.DoAPIRequest(ctx, types.APIRequestInput{
		Method:       http.MethodPost,
		APIPath:      "/auth/apiTokens",
		Body:         body,
		ResponseJSON: &respBody,
	})
	if respErr != nil {
		return nil, fmt.Errorf("couldn't create API token: %v", respErr)
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("couldn't create API token: %v", resp.Status)
	} else {
		token := respBody["apiToken"]
		return &token, nil
	}
}

// Delegate session
func (c *AuthService) DelegateSession(ctx context.Context, identityToken string, codeVerifier *string) error {
	body := map[string]interface{}{
		"loginMode":     -1,
		"identityToken": identityToken,
	}
	if codeVerifier != nil {
		body["codeVerifier"] = *codeVerifier
	}
	resp, respErr := c.client.DoAPIRequest(ctx, types.APIRequestInput{
		Method:  http.MethodPost,
		APIPath: "/auth/delegate",
		Body:    body,
	})
	if respErr != nil {
		return fmt.Errorf("couldn't delegate session: %v", respErr)
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("couldn't delegate session: %v", resp.Status)
	} else {
		newToken := resp.Header.Get("X-MSTR-AuthToken")
		c.client.AuthToken = &newToken
		return nil
	}
}

// validate identity token
func (c *AuthService) ValidateIdentityToken(ctx context.Context, identityToken string, codeVerifier *string) (*string, error) {

	var queryParams *map[string]string
	if codeVerifier != nil {
		queryParams = &map[string]string{
			"codeVerifier": *codeVerifier,
		}
	}
	request, reqErr := c.client.CreateAPIRequest(ctx, http.MethodPost, "/auth/identityToken", queryParams, nil)
	if reqErr != nil {
		return nil, fmt.Errorf("couldn't validate identity token: %v", reqErr)
	}
	request.Header.Set("X-MSTR-IdentityToken", identityToken)
	resp, respErr := c.client.Client.Do(request)
	if respErr != nil {
		return nil, fmt.Errorf("couldn't validate identity token: %v", respErr)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("couldn't validate identity token: %v", resp.Status)
	} else {
		// parse response body from json to manp[string]string
		var respBody map[string]string
		decoder := json.NewDecoder(resp.Body)
		decodeErr := decoder.Decode(&respBody)
		if decodeErr != nil {
			return nil, fmt.Errorf("couldn't validate identity token: %v", decodeErr)
		}
		userID := respBody["userId"]
		return &userID, nil
	}
}
