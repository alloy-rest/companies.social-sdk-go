// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package companiessocial

import (
	"github.com/alloy-rest/companies.social-sdk-go/internal/apijson"
	"github.com/alloy-rest/companies.social-sdk-go/packages/respjson"
)

type LookupCompanyResponse struct {
	CompanyID string                        `json:"companyId" api:"required"`
	Domain    string                        `json:"domain" api:"required"`
	GitHub    LookupCompanyResponseGitHub   `json:"github" api:"required"`
	Linkedin  LookupCompanyResponseLinkedin `json:"linkedin" api:"required"`
	Socials   map[string]string             `json:"socials" api:"required"`
	Sources   []LookupCompanyResponseSource `json:"sources" api:"required"`
	Tiktok    LookupCompanyResponseTiktok   `json:"tiktok" api:"required"`
	Youtube   LookupCompanyResponseYoutube  `json:"youtube" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyID   respjson.Field
		Domain      respjson.Field
		GitHub      respjson.Field
		Linkedin    respjson.Field
		Socials     respjson.Field
		Sources     respjson.Field
		Tiktok      respjson.Field
		Youtube     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupCompanyResponse) RawJSON() string { return r.JSON.raw }
func (r *LookupCompanyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupCompanyResponseGitHub struct {
	Blog            string                               `json:"blog" api:"required"`
	CreatedAt       string                               `json:"createdAt" api:"required"`
	Description     string                               `json:"description" api:"required"`
	Email           string                               `json:"email" api:"required"`
	Followers       float64                              `json:"followers" api:"required"`
	IsVerified      bool                                 `json:"isVerified" api:"required"`
	Location        string                               `json:"location" api:"required"`
	Name            string                               `json:"name" api:"required"`
	PublicRepos     float64                              `json:"publicRepos" api:"required"`
	Sponsoring      []string                             `json:"sponsoring" api:"required"`
	TopLanguages    []string                             `json:"topLanguages" api:"required"`
	TopRepos        []LookupCompanyResponseGitHubTopRepo `json:"topRepos" api:"required"`
	TopTopics       []string                             `json:"topTopics" api:"required"`
	TotalForks      float64                              `json:"totalForks" api:"required"`
	TotalStars      float64                              `json:"totalStars" api:"required"`
	TwitterUsername string                               `json:"twitterUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blog            respjson.Field
		CreatedAt       respjson.Field
		Description     respjson.Field
		Email           respjson.Field
		Followers       respjson.Field
		IsVerified      respjson.Field
		Location        respjson.Field
		Name            respjson.Field
		PublicRepos     respjson.Field
		Sponsoring      respjson.Field
		TopLanguages    respjson.Field
		TopRepos        respjson.Field
		TopTopics       respjson.Field
		TotalForks      respjson.Field
		TotalStars      respjson.Field
		TwitterUsername respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupCompanyResponseGitHub) RawJSON() string { return r.JSON.raw }
func (r *LookupCompanyResponseGitHub) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupCompanyResponseGitHubTopRepo struct {
	Description string  `json:"description" api:"required"`
	Forks       float64 `json:"forks" api:"required"`
	Language    string  `json:"language" api:"required"`
	Name        string  `json:"name" api:"required"`
	Stars       float64 `json:"stars" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Forks       respjson.Field
		Language    respjson.Field
		Name        respjson.Field
		Stars       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupCompanyResponseGitHubTopRepo) RawJSON() string { return r.JSON.raw }
func (r *LookupCompanyResponseGitHubTopRepo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupCompanyResponseLinkedin struct {
	CrunchbaseURL string                                  `json:"crunchbaseUrl" api:"required"`
	Description   string                                  `json:"description" api:"required"`
	EmployeeCount float64                                 `json:"employeeCount" api:"required"`
	FollowerCount float64                                 `json:"followerCount" api:"required"`
	Founded       float64                                 `json:"founded" api:"required"`
	Funding       LookupCompanyResponseLinkedinFunding    `json:"funding" api:"required"`
	Industry      string                                  `json:"industry" api:"required"`
	Locations     []LookupCompanyResponseLinkedinLocation `json:"locations" api:"required"`
	Type          string                                  `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrunchbaseURL respjson.Field
		Description   respjson.Field
		EmployeeCount respjson.Field
		FollowerCount respjson.Field
		Founded       respjson.Field
		Funding       respjson.Field
		Industry      respjson.Field
		Locations     respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupCompanyResponseLinkedin) RawJSON() string { return r.JSON.raw }
func (r *LookupCompanyResponseLinkedin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupCompanyResponseLinkedinFunding struct {
	LastRoundAmount float64 `json:"lastRoundAmount" api:"required"`
	LastRoundDate   string  `json:"lastRoundDate" api:"required"`
	LastRoundType   string  `json:"lastRoundType" api:"required"`
	TotalRounds     float64 `json:"totalRounds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastRoundAmount respjson.Field
		LastRoundDate   respjson.Field
		LastRoundType   respjson.Field
		TotalRounds     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupCompanyResponseLinkedinFunding) RawJSON() string { return r.JSON.raw }
func (r *LookupCompanyResponseLinkedinFunding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupCompanyResponseLinkedinLocation struct {
	City       string `json:"city" api:"required"`
	Country    string `json:"country" api:"required"`
	Geographic string `json:"geographic" api:"required"`
	IsHq       bool   `json:"isHq" api:"required"`
	IsPrimary  bool   `json:"isPrimary" api:"required"`
	Line1      string `json:"line1" api:"required"`
	Line2      string `json:"line2" api:"required"`
	PostalCode string `json:"postalCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Geographic  respjson.Field
		IsHq        respjson.Field
		IsPrimary   respjson.Field
		Line1       respjson.Field
		Line2       respjson.Field
		PostalCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupCompanyResponseLinkedinLocation) RawJSON() string { return r.JSON.raw }
func (r *LookupCompanyResponseLinkedinLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupCompanyResponseSource struct {
	Cached    bool   `json:"cached" api:"required"`
	Name      string `json:"name" api:"required"`
	ScrapedAt string `json:"scrapedAt" api:"required"`
	Version   string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cached      respjson.Field
		Name        respjson.Field
		ScrapedAt   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupCompanyResponseSource) RawJSON() string { return r.JSON.raw }
func (r *LookupCompanyResponseSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupCompanyResponseTiktok struct {
	AppLinks        LookupCompanyResponseTiktokAppLinks `json:"appLinks" api:"required"`
	Bio             string                              `json:"bio" api:"required"`
	BioLink         string                              `json:"bioLink" api:"required"`
	Category        string                              `json:"category" api:"required"`
	CommentSetting  float64                             `json:"commentSetting" api:"required"`
	CreateTime      float64                             `json:"createTime" api:"required"`
	DiggCount       float64                             `json:"diggCount" api:"required"`
	DownloadSetting float64                             `json:"downloadSetting" api:"required"`
	DuetSetting     float64                             `json:"duetSetting" api:"required"`
	FollowerCount   float64                             `json:"followerCount" api:"required"`
	FollowingCount  float64                             `json:"followingCount" api:"required"`
	FriendCount     float64                             `json:"friendCount" api:"required"`
	IsOrganization  bool                                `json:"isOrganization" api:"required"`
	LikeCount       float64                             `json:"likeCount" api:"required"`
	StitchSetting   float64                             `json:"stitchSetting" api:"required"`
	TtSeller        bool                                `json:"ttSeller" api:"required"`
	Verified        bool                                `json:"verified" api:"required"`
	VideoCount      float64                             `json:"videoCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppLinks        respjson.Field
		Bio             respjson.Field
		BioLink         respjson.Field
		Category        respjson.Field
		CommentSetting  respjson.Field
		CreateTime      respjson.Field
		DiggCount       respjson.Field
		DownloadSetting respjson.Field
		DuetSetting     respjson.Field
		FollowerCount   respjson.Field
		FollowingCount  respjson.Field
		FriendCount     respjson.Field
		IsOrganization  respjson.Field
		LikeCount       respjson.Field
		StitchSetting   respjson.Field
		TtSeller        respjson.Field
		Verified        respjson.Field
		VideoCount      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupCompanyResponseTiktok) RawJSON() string { return r.JSON.raw }
func (r *LookupCompanyResponseTiktok) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupCompanyResponseTiktokAppLinks struct {
	Android string `json:"android" api:"required"`
	Ios     string `json:"ios" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Android     respjson.Field
		Ios         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupCompanyResponseTiktokAppLinks) RawJSON() string { return r.JSON.raw }
func (r *LookupCompanyResponseTiktokAppLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupCompanyResponseYoutube struct {
	ChannelID       string   `json:"channelId" api:"required"`
	Description     string   `json:"description" api:"required"`
	Handle          string   `json:"handle" api:"required"`
	Keywords        []string `json:"keywords" api:"required"`
	Name            string   `json:"name" api:"required"`
	SubscriberCount float64  `json:"subscriberCount" api:"required"`
	Verified        bool     `json:"verified" api:"required"`
	VideoCount      float64  `json:"videoCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelID       respjson.Field
		Description     respjson.Field
		Handle          respjson.Field
		Keywords        respjson.Field
		Name            respjson.Field
		SubscriberCount respjson.Field
		Verified        respjson.Field
		VideoCount      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupCompanyResponseYoutube) RawJSON() string { return r.JSON.raw }
func (r *LookupCompanyResponseYoutube) UnmarshalJSON(data []byte) error {
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
