package client_test

import (
	"context"
	"log"
	"testing"

	"github.com/paulbailey/mstr-rest-client/client"
)

func TestDataSourceList(t *testing.T) {
	client := client.NewStandardMstrRestClient(username, password, baseURL)
	client.Login(context.Background())
	datasources, err := client.DataSource.GetDataSources(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(datasources) == 0 {
		t.Errorf("expected datasources")
	}
	for _, d := range datasources {
		log.Printf("datasource: %s (%s) - %s", d.Name, d.ID, d.DBMS.Name)
	}
}
