package client_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/paulbailey/mstr-rest-client/client"
)

var username = os.Getenv("MSTR_USERNAME")
var password = os.Getenv("MSTR_PASSWORD")
var baseURL = os.Getenv("MSTR_BASEURL")

func TestProjects(t *testing.T) {
	client := client.NewAnonymousMstrRestClient("https://demo.microstrategy.com/MicroStrategyLibrary/api/")
	client.Login(context.Background())
	projects, err := client.Project.GetProjects(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(projects) == 0 {
		t.Errorf("expected projects")
	}
}

func TestProjectSettings(t *testing.T) {
	// client := client.NewAnonymousMstrRestClient("https://demo.microstrategy.com/MicroStrategyLibrary/api/")
	client := client.NewStandardMstrRestClient(username, password, baseURL)
	client.Login(context.Background())
	projects, err := client.Project.GetProjects(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(projects) == 0 {
		t.Errorf("expected projects")
	} else {
		project := projects[0]
		settings, err := client.Project.GetProjectSettings(context.Background(), project.ID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if settings == nil {
			t.Errorf("expected project settings")
		}
		for k, v := range settings {
			log.Printf("%s is set to; %v", k, v)
		}
	}
}
