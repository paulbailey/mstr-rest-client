package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/paulbailey/mstr-rest-client/types"
)

func (c *MstrRestClient) GetProjects(ctx context.Context) ([]types.MstrProject, error) {
	if !c.LoggedIn() {
		return nil, fmt.Errorf("not logged in")
	}
	var projects []types.MstrProject
	_, err := c.DoAPIRequest(ctx, http.MethodGet, "/projects/", nil, nil, &projects)
	if err != nil {
		return nil, fmt.Errorf("couldn't get projects: %v", err)
	}
	return projects, nil
}

func (c *MstrRestClient) GetProject(ctx context.Context, id string) (*types.MstrProject, error) {
	if !c.LoggedIn() {
		return nil, fmt.Errorf("not logged in")
	}
	var project types.MstrProject
	_, err := c.DoAPIRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s", id), nil, nil, &project)
	if err != nil {
		return nil, fmt.Errorf("couldn't get project: %v", err)
	}
	return &project, nil
}

func (c *MstrRestClient) GetProjectSettings(ctx context.Context, id string) (types.MstrProjectSettings, error) {
	if !c.LoggedIn() {
		return nil, fmt.Errorf("not logged in")
	}
	var settings types.MstrProjectSettings
	_, err := c.DoAPIRequest(ctx, http.MethodGet, fmt.Sprintf("/v2/projects/%s/settings", id), nil, nil, &settings)
	if err != nil {
		return nil, fmt.Errorf("couldn't get project settings: %v", err)
	}
	return settings, nil
}

func (c *MstrRestClient) SetProjectSettings(ctx context.Context, id string, settings types.MstrProjectSettings) error {
	if !c.LoggedIn() {
		return fmt.Errorf("not logged in")
	}
	resp, err := c.DoAPIRequest(ctx, http.MethodPatch, fmt.Sprintf("/v2/projects/%s/settings", id), settings, nil, nil)
	if err != nil {
		return fmt.Errorf("couldn't patch project settings: %v", err)
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	if resp.StatusCode == 207 {
		return nil
	}
	if resp.StatusCode >= 400 {
		return ParseMicroStrategyError(resp)
	}
	return nil
}
