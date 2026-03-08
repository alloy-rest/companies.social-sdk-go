// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package companiessocial

import (
	"github.com/stainless-sdks/companies.social-go/internal/apijson"
	"github.com/stainless-sdks/companies.social-go/packages/respjson"
)

type LookupCompanyResponse struct {
	Cached    bool              `json:"cached" api:"required"`
	CompanyID string            `json:"companyId" api:"required"`
	Domain    string            `json:"domain" api:"required"`
	ScrapedAt string            `json:"scrapedAt" api:"required"`
	Socials   map[string]string `json:"socials" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cached      respjson.Field
		CompanyID   respjson.Field
		Domain      respjson.Field
		ScrapedAt   respjson.Field
		Socials     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupCompanyResponse) RawJSON() string { return r.JSON.raw }
func (r *LookupCompanyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupCompanyParams struct {
	// Identifier type
	//
	// Any of "domain", "twitter", "linkedin", "github", "tiktok", "instagram",
	// "facebook", "youtube".
	IDType LookupCompanyParamsIDType `path:"idType,omitzero" api:"required" json:"-"`
	paramObj
}

// Identifier type
type LookupCompanyParamsIDType string

const (
	LookupCompanyParamsIDTypeDomain    LookupCompanyParamsIDType = "domain"
	LookupCompanyParamsIDTypeTwitter   LookupCompanyParamsIDType = "twitter"
	LookupCompanyParamsIDTypeLinkedin  LookupCompanyParamsIDType = "linkedin"
	LookupCompanyParamsIDTypeGitHub    LookupCompanyParamsIDType = "github"
	LookupCompanyParamsIDTypeTiktok    LookupCompanyParamsIDType = "tiktok"
	LookupCompanyParamsIDTypeInstagram LookupCompanyParamsIDType = "instagram"
	LookupCompanyParamsIDTypeFacebook  LookupCompanyParamsIDType = "facebook"
	LookupCompanyParamsIDTypeYoutube   LookupCompanyParamsIDType = "youtube"
)
