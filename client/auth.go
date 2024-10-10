package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/paulbailey/mstr-rest-client/types"
)

func (c *MstrRestClient) Authenticate() error {
	return nil
}

func (c *MstrRestClient) Login(ctx context.Context, applicationType *types.MstrApplicationType) error {
	if c.Authentication == nil {
		return fmt.Errorf("no authentication method set")
	}
	authReq := c.Authentication.AuthenticationRequest(applicationType)
	resp, respErr := c.DoAPIRequest(ctx, http.MethodPost, "/auth/login", authReq, nil, nil)
	if respErr != nil {
		return fmt.Errorf("couldn't login: %v", respErr)
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("couldn't login: %v", resp.Status)
	} else {
		// checkov:skip=CKV_SECRET_6: Not a secret
		token := resp.Header.Get("X-MSTR-AuthToken")
		c.AuthToken = &token
	}
	return nil
}

func (c *MstrRestClient) Logout(ctx context.Context) error {
	resp, respErr := c.DoAPIRequest(ctx, http.MethodPost, "/auth/logout", nil, nil, nil)
	if respErr != nil {
		return fmt.Errorf("couldn't logout: %v", respErr)
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("couldn't logout: %v", resp.Status)
	} else {
		c.AuthToken = nil
		return nil
	}
}
