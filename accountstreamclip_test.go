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

func TestAccountStreamClipStreamVideoClippingClipVideosGivenAStartAndEndTimeWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Streams.Clips.StreamVideoClippingClipVideosGivenAStartAndEndTime(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountStreamClipStreamVideoClippingClipVideosGivenAStartAndEndTimeParams{
			AllowedOrigins:        cloudflare.F([]string{"example.com"}),
			Creator:               cloudflare.F("creator-id_abcde12345"),
			MaxDurationSeconds:    cloudflare.F(int64(1)),
			ThumbnailTimestampPct: cloudflare.F(0.529241),
			Watermark: cloudflare.F(cloudflare.AccountStreamClipStreamVideoClippingClipVideosGivenAStartAndEndTimeParamsWatermark{
				Uid: cloudflare.F("ea95132c15732412d22c1476fa83f27a"),
			}),
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
