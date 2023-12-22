// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestRadarEmailTimeserySpamListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.Emails.Timeseries.Spams.List(context.TODO(), cloudflare.RadarEmailTimeserySpamListParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailTimeserySpamListParamsAggInterval15m),
		Arc:         cloudflare.F([]cloudflare.RadarEmailTimeserySpamListParamsArc{cloudflare.RadarEmailTimeserySpamListParamsArcPass, cloudflare.RadarEmailTimeserySpamListParamsArcNone, cloudflare.RadarEmailTimeserySpamListParamsArcFail}),
		ASN:         cloudflare.F([]string{"15169", "15169", "15169"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailTimeserySpamListParamsDateRange{cloudflare.RadarEmailTimeserySpamListParamsDateRange1d, cloudflare.RadarEmailTimeserySpamListParamsDateRange7d, cloudflare.RadarEmailTimeserySpamListParamsDateRange14d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dkim:        cloudflare.F([]cloudflare.RadarEmailTimeserySpamListParamsDkim{cloudflare.RadarEmailTimeserySpamListParamsDkimPass, cloudflare.RadarEmailTimeserySpamListParamsDkimNone, cloudflare.RadarEmailTimeserySpamListParamsDkimFail}),
		Dmarc:       cloudflare.F([]cloudflare.RadarEmailTimeserySpamListParamsDmarc{cloudflare.RadarEmailTimeserySpamListParamsDmarcPass, cloudflare.RadarEmailTimeserySpamListParamsDmarcNone, cloudflare.RadarEmailTimeserySpamListParamsDmarcFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailTimeserySpamListParamsFormatJson),
		Location:    cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:        cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Spf:         cloudflare.F([]cloudflare.RadarEmailTimeserySpamListParamsSpf{cloudflare.RadarEmailTimeserySpamListParamsSpfPass, cloudflare.RadarEmailTimeserySpamListParamsSpfNone, cloudflare.RadarEmailTimeserySpamListParamsSpfFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
