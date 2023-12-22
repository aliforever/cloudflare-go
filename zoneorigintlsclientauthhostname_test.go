// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestZoneOriginTlsClientAuthHostnameGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Zones.OriginTlsClientAuths.Hostnames.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"app.example.com",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthentication(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Zones.OriginTlsClientAuths.Hostnames.PerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthentication(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationParams{
			Config: cloudflare.F([]cloudflare.ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationParamsConfig{{
				CertID:   cloudflare.F("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60"),
				Enabled:  cloudflare.F(true),
				Hostname: cloudflare.F("app.example.com"),
			}, {
				CertID:   cloudflare.F("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60"),
				Enabled:  cloudflare.F(true),
				Hostname: cloudflare.F("app.example.com"),
			}, {
				CertID:   cloudflare.F("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60"),
				Enabled:  cloudflare.F(true),
				Hostname: cloudflare.F("app.example.com"),
			}}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
