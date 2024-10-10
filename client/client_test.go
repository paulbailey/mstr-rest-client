package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paulbailey/mstr-rest-client/client"
	"github.com/paulbailey/mstr-rest-client/types"
)

func TestAnonymousLogin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/auth/login" {
			w.Header().Set("X-MSTR-AuthToken", "test-token")
			w.WriteHeader(http.StatusNoContent)
		} else {
			t.Errorf("unexpected request: %s", r.URL.Path)
		}
	}))
	defer ts.Close()

	client := client.NewAnonymousMstrRestClient(ts.URL)
	custom := types.CustomApp
	err := client.Login(context.Background(), &custom)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(*client.AuthToken) == 0 {
		t.Errorf("expected auth token")
	}
}

func TestLogout(t *testing.T) {
	client := client.NewAnonymousMstrRestClient("https://demo.microstrategy.com/MicroStrategyLibrary/api")
	custom := types.CustomApp
	client.Login(context.Background(), &custom)
	err := client.Logout(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
