package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/paulbailey/mstr-rest-client/types"
)

func (c *MstrRestClient) GetDataSources(ctx context.Context) ([]types.MstrRestDataSource, error) {
	if !c.LoggedIn() {
		return nil, fmt.Errorf("not logged in")
	}
	var response struct {
		Datasources []types.MstrRestDataSource `json:"datasources"`
	}
	_, err := c.DoAPIRequest(ctx, http.MethodGet, "/datasources", nil, nil, &response)
	if err != nil {
		return nil, err
	}
	return response.Datasources, nil
}

func (c *MstrRestClient) GetDataSource(ctx context.Context, id string) (*types.MstrRestDataSource, error) {
	if !c.LoggedIn() {
		return nil, fmt.Errorf("not logged in")
	}
	var datasource types.MstrRestDataSource

	_, err := c.DoAPIRequest(ctx, http.MethodGet, fmt.Sprintf("/datasources/%s", id), nil, nil, &datasource)
	if err != nil {
		return nil, err
	}
	return &datasource, nil
}

func (c *MstrRestClient) DeleteDataSource(ctx context.Context, id string) error {
	if !c.LoggedIn() {
		return fmt.Errorf("not logged in")
	}
	resp, err := c.DoAPIRequest(ctx, http.MethodDelete, fmt.Sprintf("/datasources/%s", id), nil, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP error deleting datasource: %s", resp.Status)
	}
	return nil
}

func (c *MstrRestClient) CreateDataSource(ctx context.Context, datasource types.MstrRestDataSource) (*types.MstrRestDataSource, error) {
	if !c.LoggedIn() {
		return nil, fmt.Errorf("not logged in")
	}
	var response struct {
		DataSource types.MstrRestDataSource `json:"datasource"`
	}
	_, err := c.DoAPIRequest(ctx, http.MethodPost, "/datasources", datasource, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response.DataSource, nil
}
