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
	AuthToken      *string
	Authentication types.MstrAuthentication
	BaseURL        string
}

func NewMstrRestClient(auth types.MstrAuthentication, baseURL string) *MstrRestClient {
	client := &MstrRestClient{
		Authentication: auth,
		BaseURL:        baseURL,
	}

	jar, cookieJarErr := cookiejar.New(nil)
	if cookieJarErr != nil {
		panic(cookieJarErr)
	}
	client.Jar = jar

	return client
}

func NewAnonymousMstrRestClient(baseURL string) *MstrRestClient {
	return NewMstrRestClient(&types.AnonymousAuthentication{}, baseURL)
}

func NewStandardMstrRestClient(username, password, baseURL string) *MstrRestClient {
	return NewMstrRestClient(&types.StandardAuthentication{Username: username, Password: password}, baseURL)
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

func (c *MstrRestClient) DoAPIRequest(ctx context.Context, method string, apiPath string, body interface{}, queryParams *map[string]string, parsedResponse interface{}) (*http.Response, error) {
	req, reqErr := c.CreateAPIRequest(ctx, method, apiPath, queryParams, body)
	if reqErr != nil {
		return nil, reqErr
	}
	slog.Debug("Request", slog.String("method", req.Method), slog.Any("url", req.URL))
	resp, err := c.Do(req)
	slog.Debug("Response:", slog.String("responseCode", resp.Status))
	if err != nil {
		return resp, err
	}
	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		return resp, ParseMicroStrategyError(resp)
	}
	if parsedResponse != nil {
		parseErr := json.NewDecoder(resp.Body).Decode(parsedResponse)
		if parseErr != nil {
			return resp, parseErr
		}
	}
	return resp, nil
}

func ParseMicroStrategyError(resp *http.Response) *types.MstrRestError {
	var err types.MstrRestError
	json.NewDecoder(resp.Body).Decode(&err)
	return &err
}

func (c *MstrRestClient) LoggedIn() bool {
	return c.AuthToken != nil
}
