// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneRulesetPhaseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneRulesetPhaseService] method
// instead.
type ZoneRulesetPhaseService struct {
	Options                     []option.RequestOption
	HTTPConfigSettings          *ZoneRulesetPhaseHTTPConfigSettingService
	HTTPCustomErrors            *ZoneRulesetPhaseHTTPCustomErrorService
	HTTPRequestCacheSettings    *ZoneRulesetPhaseHTTPRequestCacheSettingService
	HTTPRequestDynamicRedirects *ZoneRulesetPhaseHTTPRequestDynamicRedirectService
	HTTPRequestOrigins          *ZoneRulesetPhaseHTTPRequestOriginService
	Entrypoints                 *ZoneRulesetPhaseEntrypointService
}

// NewZoneRulesetPhaseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneRulesetPhaseService(opts ...option.RequestOption) (r *ZoneRulesetPhaseService) {
	r = &ZoneRulesetPhaseService{}
	r.Options = opts
	r.HTTPConfigSettings = NewZoneRulesetPhaseHTTPConfigSettingService(opts...)
	r.HTTPCustomErrors = NewZoneRulesetPhaseHTTPCustomErrorService(opts...)
	r.HTTPRequestCacheSettings = NewZoneRulesetPhaseHTTPRequestCacheSettingService(opts...)
	r.HTTPRequestDynamicRedirects = NewZoneRulesetPhaseHTTPRequestDynamicRedirectService(opts...)
	r.HTTPRequestOrigins = NewZoneRulesetPhaseHTTPRequestOriginService(opts...)
	r.Entrypoints = NewZoneRulesetPhaseEntrypointService(opts...)
	return
}
