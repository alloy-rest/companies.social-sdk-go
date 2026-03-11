// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package companiessocial

import (
	"context"
	"net/http"
	"slices"

	"github.com/alloy-rest/-companies.social-sdk-go/internal/apijson"
	"github.com/alloy-rest/-companies.social-sdk-go/internal/requestconfig"
	"github.com/alloy-rest/-companies.social-sdk-go/option"
	"github.com/alloy-rest/-companies.social-sdk-go/packages/respjson"
)

// WhoamiService contains methods and other services that help with interacting
// with the companies.social API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhoamiService] method instead.
type WhoamiService struct {
	Options []option.RequestOption
}

// NewWhoamiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWhoamiService(opts ...option.RequestOption) (r WhoamiService) {
	r = WhoamiService{}
	r.Options = opts
	return
}

func (r *WhoamiService) Get(ctx context.Context, opts ...option.RequestOption) (res *WhoamiGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "whoami"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type WhoamiGetResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	Email     string `json:"email" api:"required"`
	Name      string `json:"name" api:"required"`
	Plan      string `json:"plan" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		Plan        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhoamiGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WhoamiGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
