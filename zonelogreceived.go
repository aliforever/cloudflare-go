// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneLogReceivedService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneLogReceivedService] method
// instead.
type ZoneLogReceivedService struct {
	Options []option.RequestOption
	Fields  *ZoneLogReceivedFieldService
}

// NewZoneLogReceivedService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneLogReceivedService(opts ...option.RequestOption) (r *ZoneLogReceivedService) {
	r = &ZoneLogReceivedService{}
	r.Options = opts
	r.Fields = NewZoneLogReceivedFieldService(opts...)
	return
}

// The `/received` api route allows customers to retrieve their edge HTTP logs. The
// basic access pattern is "give me all the logs for zone Z for minute M", where
// the minute M refers to the time records were received at Cloudflare's central
// data center. `start` is inclusive, and `end` is exclusive. Because of that, to
// get all data, at minutely cadence, starting at 10AM, the proper values are:
// `start=2018-05-20T10:00:00Z&end=2018-05-20T10:01:00Z`, then
// `start=2018-05-20T10:01:00Z&end=2018-05-20T10:02:00Z` and so on; the overlap
// will be handled properly.
func (r *ZoneLogReceivedService) ReceivedGetLogsReceived(ctx context.Context, zoneIdentifier string, query ZoneLogReceivedReceivedGetLogsReceivedParams, opts ...option.RequestOption) (res *Log, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logs/received", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneLogReceivedReceivedGetLogsReceivedParams struct {
	// Sets the (exclusive) end of the requested time frame. This can be a unix
	// timestamp (in seconds or nanoseconds), or an absolute timestamp that conforms to
	// RFC 3339. `end` must be at least five minutes earlier than now and must be later
	// than `start`. Difference between `start` and `end` must be not greater than one
	// hour.
	End param.Field[ZoneLogReceivedReceivedGetLogsReceivedParamsEnd] `query:"end,required"`
	// When `?count=` is provided, the response will contain up to `count` results.
	// Since results are not sorted, you are likely to get different data for repeated
	// requests. `count` must be an integer > 0.
	Count param.Field[int64] `query:"count"`
	// The `/received` route by default returns a limited set of fields, and allows
	// customers to override the default field set by specifying individual fields. The
	// reasons for this are: 1. Most customers require only a small subset of fields,
	// but that subset varies from customer to customer; 2. Flat schema is much easier
	// to work with downstream (importing into BigTable etc); 3. Performance (time to
	// process, file size). If `?fields=` is not specified, default field set is
	// returned. This default field set may change at any time. When `?fields=` is
	// provided, each record is returned with the specified fields. `fields` must be
	// specified as a comma separated list without any whitespaces, and all fields must
	// exist. The order in which fields are specified does not matter, and the order of
	// fields in the response is not specified.
	Fields param.Field[string] `query:"fields"`
	// When `?sample=` is provided, a sample of matching records is returned. If
	// `sample=0.1` then 10% of records will be returned. Sampling is random: repeated
	// calls will not only return different records, but likely will also vary slightly
	// in number of returned records. When `?count=` is also specified, `count` is
	// applied to the number of returned records, not the sampled records. So, with
	// `sample=0.05` and `count=7`, when there is a total of 100 records available,
	// approximately five will be returned. When there are 1000 records, seven will be
	// returned. When there are 10,000 records, seven will be returned.
	Sample param.Field[float64] `query:"sample"`
	// Sets the (inclusive) beginning of the requested time frame. This can be a unix
	// timestamp (in seconds or nanoseconds), or an absolute timestamp that conforms to
	// RFC 3339. At this point in time, it cannot exceed a time in the past greater
	// than seven days.
	Start param.Field[ZoneLogReceivedReceivedGetLogsReceivedParamsStart] `query:"start"`
	// By default, timestamps in responses are returned as Unix nanosecond integers.
	// The `?timestamps=` argument can be set to change the format in which response
	// timestamps are returned. Possible values are: `unix`, `unixnano`, `rfc3339`.
	// Note that `unix` and `unixnano` return timestamps as integers; `rfc3339` returns
	// timestamps as strings.
	Timestamps param.Field[ZoneLogReceivedReceivedGetLogsReceivedParamsTimestamps] `query:"timestamps"`
}

// URLQuery serializes [ZoneLogReceivedReceivedGetLogsReceivedParams]'s query
// parameters as `url.Values`.
func (r ZoneLogReceivedReceivedGetLogsReceivedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sets the (exclusive) end of the requested time frame. This can be a unix
// timestamp (in seconds or nanoseconds), or an absolute timestamp that conforms to
// RFC 3339. `end` must be at least five minutes earlier than now and must be later
// than `start`. Difference between `start` and `end` must be not greater than one
// hour.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type ZoneLogReceivedReceivedGetLogsReceivedParamsEnd interface {
	ImplementsZoneLogReceivedReceivedGetLogsReceivedParamsEnd()
}

// Sets the (inclusive) beginning of the requested time frame. This can be a unix
// timestamp (in seconds or nanoseconds), or an absolute timestamp that conforms to
// RFC 3339. At this point in time, it cannot exceed a time in the past greater
// than seven days.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type ZoneLogReceivedReceivedGetLogsReceivedParamsStart interface {
	ImplementsZoneLogReceivedReceivedGetLogsReceivedParamsStart()
}

// By default, timestamps in responses are returned as Unix nanosecond integers.
// The `?timestamps=` argument can be set to change the format in which response
// timestamps are returned. Possible values are: `unix`, `unixnano`, `rfc3339`.
// Note that `unix` and `unixnano` return timestamps as integers; `rfc3339` returns
// timestamps as strings.
type ZoneLogReceivedReceivedGetLogsReceivedParamsTimestamps string

const (
	ZoneLogReceivedReceivedGetLogsReceivedParamsTimestampsUnix     ZoneLogReceivedReceivedGetLogsReceivedParamsTimestamps = "unix"
	ZoneLogReceivedReceivedGetLogsReceivedParamsTimestampsUnixnano ZoneLogReceivedReceivedGetLogsReceivedParamsTimestamps = "unixnano"
	ZoneLogReceivedReceivedGetLogsReceivedParamsTimestampsRfc3339  ZoneLogReceivedReceivedGetLogsReceivedParamsTimestamps = "rfc3339"
)
