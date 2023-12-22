// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAccessCertificateService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountAccessCertificateService] method instead.
type AccountAccessCertificateService struct {
	Options []option.RequestOption
}

// NewAccountAccessCertificateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessCertificateService(opts ...option.RequestOption) (r *AccountAccessCertificateService) {
	r = &AccountAccessCertificateService{}
	r.Options = opts
	return
}

// Fetches a single mTLS certificate.
func (r *AccountAccessCertificateService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *CertificatesSingleResponseBEbBQusK, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/certificates/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured mTLS certificate.
func (r *AccountAccessCertificateService) Update(ctx context.Context, identifier string, uuid string, body AccountAccessCertificateUpdateParams, opts ...option.RequestOption) (res *CertificatesSingleResponseBEbBQusK, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/certificates/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an mTLS certificate.
func (r *AccountAccessCertificateService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccountAccessCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/certificates/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new mTLS root certificate to Access.
func (r *AccountAccessCertificateService) AccessMTlsAuthenticationAddAnMTlsCertificate(ctx context.Context, identifier string, body AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateParams, opts ...option.RequestOption) (res *CertificatesSingleResponseBEbBQusK, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/certificates", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all mTLS root certificates.
func (r *AccountAccessCertificateService) AccessMTlsAuthenticationListMTlsCertificates(ctx context.Context, identifier string, opts ...option.RequestOption) (res *CertificatesResponseCollectionFGyg2vQi, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/certificates", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CertificatesResponseCollectionFGyg2vQi struct {
	Errors     []CertificatesResponseCollectionFGyg2vQiError    `json:"errors"`
	Messages   []CertificatesResponseCollectionFGyg2vQiMessage  `json:"messages"`
	Result     []CertificatesResponseCollectionFGyg2vQiResult   `json:"result"`
	ResultInfo CertificatesResponseCollectionFGyg2vQiResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success CertificatesResponseCollectionFGyg2vQiSuccess `json:"success"`
	JSON    certificatesResponseCollectionFGyg2vQiJSON    `json:"-"`
}

// certificatesResponseCollectionFGyg2vQiJSON contains the JSON metadata for the
// struct [CertificatesResponseCollectionFGyg2vQi]
type certificatesResponseCollectionFGyg2vQiJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatesResponseCollectionFGyg2vQi) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatesResponseCollectionFGyg2vQiError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    certificatesResponseCollectionFGyg2vQiErrorJSON `json:"-"`
}

// certificatesResponseCollectionFGyg2vQiErrorJSON contains the JSON metadata for
// the struct [CertificatesResponseCollectionFGyg2vQiError]
type certificatesResponseCollectionFGyg2vQiErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatesResponseCollectionFGyg2vQiError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatesResponseCollectionFGyg2vQiMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    certificatesResponseCollectionFGyg2vQiMessageJSON `json:"-"`
}

// certificatesResponseCollectionFGyg2vQiMessageJSON contains the JSON metadata for
// the struct [CertificatesResponseCollectionFGyg2vQiMessage]
type certificatesResponseCollectionFGyg2vQiMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatesResponseCollectionFGyg2vQiMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatesResponseCollectionFGyg2vQiResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                           `json:"name"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	JSON      certificatesResponseCollectionFGyg2vQiResultJSON `json:"-"`
}

// certificatesResponseCollectionFGyg2vQiResultJSON contains the JSON metadata for
// the struct [CertificatesResponseCollectionFGyg2vQiResult]
type certificatesResponseCollectionFGyg2vQiResultJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CertificatesResponseCollectionFGyg2vQiResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatesResponseCollectionFGyg2vQiResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       certificatesResponseCollectionFGyg2vQiResultInfoJSON `json:"-"`
}

// certificatesResponseCollectionFGyg2vQiResultInfoJSON contains the JSON metadata
// for the struct [CertificatesResponseCollectionFGyg2vQiResultInfo]
type certificatesResponseCollectionFGyg2vQiResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatesResponseCollectionFGyg2vQiResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificatesResponseCollectionFGyg2vQiSuccess bool

const (
	CertificatesResponseCollectionFGyg2vQiSuccessTrue CertificatesResponseCollectionFGyg2vQiSuccess = true
)

type CertificatesSingleResponseBEbBQusK struct {
	Errors   []CertificatesSingleResponseBEbBQusKError   `json:"errors"`
	Messages []CertificatesSingleResponseBEbBQusKMessage `json:"messages"`
	Result   CertificatesSingleResponseBEbBQusKResult    `json:"result"`
	// Whether the API call was successful
	Success CertificatesSingleResponseBEbBQusKSuccess `json:"success"`
	JSON    certificatesSingleResponseBEbBQusKJSON    `json:"-"`
}

// certificatesSingleResponseBEbBQusKJSON contains the JSON metadata for the struct
// [CertificatesSingleResponseBEbBQusK]
type certificatesSingleResponseBEbBQusKJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatesSingleResponseBEbBQusK) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatesSingleResponseBEbBQusKError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    certificatesSingleResponseBEbBQusKErrorJSON `json:"-"`
}

// certificatesSingleResponseBEbBQusKErrorJSON contains the JSON metadata for the
// struct [CertificatesSingleResponseBEbBQusKError]
type certificatesSingleResponseBEbBQusKErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatesSingleResponseBEbBQusKError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatesSingleResponseBEbBQusKMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    certificatesSingleResponseBEbBQusKMessageJSON `json:"-"`
}

// certificatesSingleResponseBEbBQusKMessageJSON contains the JSON metadata for the
// struct [CertificatesSingleResponseBEbBQusKMessage]
type certificatesSingleResponseBEbBQusKMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatesSingleResponseBEbBQusKMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatesSingleResponseBEbBQusKResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                       `json:"name"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      certificatesSingleResponseBEbBQusKResultJSON `json:"-"`
}

// certificatesSingleResponseBEbBQusKResultJSON contains the JSON metadata for the
// struct [CertificatesSingleResponseBEbBQusKResult]
type certificatesSingleResponseBEbBQusKResultJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CertificatesSingleResponseBEbBQusKResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificatesSingleResponseBEbBQusKSuccess bool

const (
	CertificatesSingleResponseBEbBQusKSuccessTrue CertificatesSingleResponseBEbBQusKSuccess = true
)

type AccountAccessCertificateDeleteResponse struct {
	Errors   []AccountAccessCertificateDeleteResponseError   `json:"errors"`
	Messages []AccountAccessCertificateDeleteResponseMessage `json:"messages"`
	Result   AccountAccessCertificateDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessCertificateDeleteResponseSuccess `json:"success"`
	JSON    accountAccessCertificateDeleteResponseJSON    `json:"-"`
}

// accountAccessCertificateDeleteResponseJSON contains the JSON metadata for the
// struct [AccountAccessCertificateDeleteResponse]
type accountAccessCertificateDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateDeleteResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountAccessCertificateDeleteResponseErrorJSON `json:"-"`
}

// accountAccessCertificateDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountAccessCertificateDeleteResponseError]
type accountAccessCertificateDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateDeleteResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountAccessCertificateDeleteResponseMessageJSON `json:"-"`
}

// accountAccessCertificateDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountAccessCertificateDeleteResponseMessage]
type accountAccessCertificateDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateDeleteResponseResult struct {
	// UUID
	ID   string                                           `json:"id"`
	JSON accountAccessCertificateDeleteResponseResultJSON `json:"-"`
}

// accountAccessCertificateDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountAccessCertificateDeleteResponseResult]
type accountAccessCertificateDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessCertificateDeleteResponseSuccess bool

const (
	AccountAccessCertificateDeleteResponseSuccessTrue AccountAccessCertificateDeleteResponseSuccess = true
)

type AccountAccessCertificateUpdateParams struct {
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateParams struct {
	// The certificate content.
	Certificate param.Field[string] `json:"certificate,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name,required"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames"`
}

func (r AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
