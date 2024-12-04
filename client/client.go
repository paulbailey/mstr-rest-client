// Package client implements an HTTP client for the MicroStrategy REST API.
//
// The client provides methods for interacting with the MicroStrategy REST API, handles authentication,
// and where appropriate returns strongly typed responses.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/paulbailey/mstr-rest-client/types"
)

type MstrRestClient struct {
	http.Client
	AuthToken       *string
	Authentication  types.MstrAuthentication
	BaseURL         string
	ApplicationType *types.MstrApplicationType

	shared     service
	Auth       *AuthService
	DataSource *DataSourceService
	Project    *ProjectService
}

// NewMstrRestClient creates a new MstrRestClient with the given authentication and base URL.
func NewMstrRestClient(auth types.MstrAuthentication, baseURL string, configs ...func(*types.ClientConfig)) *MstrRestClient {
	config := types.ClientConfig{
		ApplicationType: nil,
	}
	for _, cfgFunc := range configs {
		cfgFunc(&config)
	}
	client := &MstrRestClient{
		Authentication: auth,
		BaseURL:        baseURL,
	}

	client.ApplicationType = config.ApplicationType

	jar, cookieJarErr := cookiejar.New(nil)
	if cookieJarErr != nil {
		panic(cookieJarErr)
	}
	client.Jar = jar

	client.initialise()

	return client
}

// NewAnonymousMstrRestClient creates a new MstrRestClient with anonymous authentication and the given base URL.
func NewAnonymousMstrRestClient(baseURL string, configs ...func(*types.ClientConfig)) *MstrRestClient {
	return NewMstrRestClient(&types.AnonymousAuthentication{}, baseURL, configs...)
}

// NewStandardMstrRestClient creates a new MstrRestClient with standard authentication and the given username, password, and base URL.
func NewStandardMstrRestClient(username, password, baseURL string, configs ...func(*types.ClientConfig)) *MstrRestClient {
	return NewMstrRestClient(&types.StandardAuthentication{Username: username, Password: password}, baseURL, configs...)
}

func (c *MstrRestClient) composeURL(apiPath string, queryParams *map[string]string) (*string, error) {

	pth, err := url.JoinPath(c.BaseURL, apiPath)
	if queryParams != nil {
		pth += "?"
		for key, value := range *queryParams {
			pth += fmt.Sprintf("%s=%s&", url.QueryEscape(key), url.QueryEscape(value))
		}
	}
	if err != nil {
		return nil, err
	}
	return &pth, nil
}

// CreateAPIRequest creates an HTTP request for the given method, API path, query parameters, and body.
func (c *MstrRestClient) CreateAPIRequest(ctx context.Context, method string, apiPath string, queryParams *map[string]string, body interface{}) (*http.Request, error) {
	url, urlErr := c.composeURL(apiPath, queryParams)
	if urlErr != nil {
		return nil, urlErr
	}
	var (
		req    *http.Request
		reqErr error
	)
	if body == nil {
		req, reqErr = http.NewRequestWithContext(ctx, method, *url, nil)
	} else {

		jsonData, jsonErr := json.Marshal(body)

		if jsonErr != nil {
			return nil, jsonErr
		}
		req, reqErr = http.NewRequestWithContext(ctx, method, *url, bytes.NewBuffer(jsonData))
	}
	if reqErr != nil {
		return nil, reqErr
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.AuthToken != nil {
		req.Header.Set("X-MSTR-AuthToken", *c.AuthToken)
	}
	if method == http.MethodDelete {
		req.Header.Set("Accept", "*/*")

	} else {
		req.Header.Set("Accept", "application/json")
	}
	return req, nil
}

// DoAPIRequest performs an API request with the given method, API path, body, query parameters, and parsed response.
func (c *MstrRestClient) DoAPIRequest(ctx context.Context, requestOptions types.APIRequestInput) (*http.Response, error) {
	req, reqErr := c.CreateAPIRequest(ctx, requestOptions.Method, requestOptions.APIPath, requestOptions.QueryParams, requestOptions.Body)
	if reqErr != nil {
		return nil, reqErr
	}
	slog.Debug("Request", slog.String("method", req.Method), slog.Any("url", req.URL))
	resp, err := c.Do(req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		return resp, ParseMicroStrategyError(resp)
	}
	if requestOptions.ResponseJSON != nil {
		parseErr := json.NewDecoder(resp.Body).Decode(requestOptions.ResponseJSON)
		if parseErr != nil {
			return resp, parseErr
		}
	}
	return resp, nil
}

// ParseMicroStrategyError parses a MicroStrategy REST API error from the given response.
func ParseMicroStrategyError(resp *http.Response) *types.MstrRestError {
	var err types.MstrRestError
	json.NewDecoder(resp.Body).Decode(&err)
	return &err
}

// LoggedIn returns true if the client is logged in.
func (c *MstrRestClient) LoggedIn() bool {
	return c.AuthToken != nil
}

func (c *MstrRestClient) initialise() {
	c.shared.client = c
	c.Auth = (*AuthService)(&c.shared)
	c.DataSource = (*DataSourceService)(&c.shared)
	c.Project = (*ProjectService)(&c.shared)
}

func (c *MstrRestClient) Login(ctx context.Context) error {
	return c.Auth.Login(ctx)
}

func (c *MstrRestClient) Logout(ctx context.Context) error {
	return c.Auth.Logout(ctx)
}

type service struct {
	client *MstrRestClient
}
