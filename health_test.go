// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package companiessocial_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/companies.social-go"
	"github.com/stainless-sdks/companies.social-go/internal/testutil"
	"github.com/stainless-sdks/companies.social-go/option"
)

func TestHealthCheck(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := companiessocial.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Health.Check(context.TODO())
	if err != nil {
		var apierr *companiessocial.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
