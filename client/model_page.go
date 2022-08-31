/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
	"time"
)

// Page The configuration for GitHub Pages for a repository.
type Page struct {
	// The API address for accessing this Page resource.
	Url string `json:"url"`
	// The status of the most recent build of the Page.
	Status NullableString `json:"status"`
	// The Pages site's custom domain
	Cname NullableString `json:"cname"`
	// The state if the domain is verified
	ProtectedDomainState NullableString `json:"protected_domain_state,omitempty"`
	// The timestamp when a pending domain becomes unverified.
	PendingDomainUnverifiedAt NullableTime `json:"pending_domain_unverified_at,omitempty"`
	// Whether the Page has a custom 404 page.
	Custom404 bool `json:"custom_404"`
	// The web address the Page can be accessed from.
	HtmlUrl *string `json:"html_url,omitempty"`
	// The process in which the Page will be built.
	BuildType NullableString `json:"build_type,omitempty"`
	Source *PagesSourceHash `json:"source,omitempty"`
	// Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
	Public bool `json:"public"`
	HttpsCertificate *PagesHttpsCertificate `json:"https_certificate,omitempty"`
	// Whether https is enabled on the domain
	HttpsEnforced *bool `json:"https_enforced,omitempty"`
}

// NewPage instantiates a new Page object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPage(url string, status NullableString, cname NullableString, custom404 bool, public bool) *Page {
	this := Page{}
	this.Url = url
	this.Status = status
	this.Cname = cname
	this.Custom404 = custom404
	this.Public = public
	return &this
}

// NewPageWithDefaults instantiates a new Page object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageWithDefaults() *Page {
	this := Page{}
	var custom404 bool = false
	this.Custom404 = custom404
	return &this
}

// GetUrl returns the Url field value
func (o *Page) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Page) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Page) SetUrl(v string) {
	o.Url = v
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Page) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Page) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *Page) SetStatus(v string) {
	o.Status.Set(&v)
}

// GetCname returns the Cname field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Page) GetCname() string {
	if o == nil || o.Cname.Get() == nil {
		var ret string
		return ret
	}

	return *o.Cname.Get()
}

// GetCnameOk returns a tuple with the Cname field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Page) GetCnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cname.Get(), o.Cname.IsSet()
}

// SetCname sets field value
func (o *Page) SetCname(v string) {
	o.Cname.Set(&v)
}

// GetProtectedDomainState returns the ProtectedDomainState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Page) GetProtectedDomainState() string {
	if o == nil || o.ProtectedDomainState.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectedDomainState.Get()
}

// GetProtectedDomainStateOk returns a tuple with the ProtectedDomainState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Page) GetProtectedDomainStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectedDomainState.Get(), o.ProtectedDomainState.IsSet()
}

// HasProtectedDomainState returns a boolean if a field has been set.
func (o *Page) HasProtectedDomainState() bool {
	if o != nil && o.ProtectedDomainState.IsSet() {
		return true
	}

	return false
}

// SetProtectedDomainState gets a reference to the given NullableString and assigns it to the ProtectedDomainState field.
func (o *Page) SetProtectedDomainState(v string) {
	o.ProtectedDomainState.Set(&v)
}
// SetProtectedDomainStateNil sets the value for ProtectedDomainState to be an explicit nil
func (o *Page) SetProtectedDomainStateNil() {
	o.ProtectedDomainState.Set(nil)
}

// UnsetProtectedDomainState ensures that no value is present for ProtectedDomainState, not even an explicit nil
func (o *Page) UnsetProtectedDomainState() {
	o.ProtectedDomainState.Unset()
}

// GetPendingDomainUnverifiedAt returns the PendingDomainUnverifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Page) GetPendingDomainUnverifiedAt() time.Time {
	if o == nil || o.PendingDomainUnverifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.PendingDomainUnverifiedAt.Get()
}

// GetPendingDomainUnverifiedAtOk returns a tuple with the PendingDomainUnverifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Page) GetPendingDomainUnverifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PendingDomainUnverifiedAt.Get(), o.PendingDomainUnverifiedAt.IsSet()
}

// HasPendingDomainUnverifiedAt returns a boolean if a field has been set.
func (o *Page) HasPendingDomainUnverifiedAt() bool {
	if o != nil && o.PendingDomainUnverifiedAt.IsSet() {
		return true
	}

	return false
}

// SetPendingDomainUnverifiedAt gets a reference to the given NullableTime and assigns it to the PendingDomainUnverifiedAt field.
func (o *Page) SetPendingDomainUnverifiedAt(v time.Time) {
	o.PendingDomainUnverifiedAt.Set(&v)
}
// SetPendingDomainUnverifiedAtNil sets the value for PendingDomainUnverifiedAt to be an explicit nil
func (o *Page) SetPendingDomainUnverifiedAtNil() {
	o.PendingDomainUnverifiedAt.Set(nil)
}

// UnsetPendingDomainUnverifiedAt ensures that no value is present for PendingDomainUnverifiedAt, not even an explicit nil
func (o *Page) UnsetPendingDomainUnverifiedAt() {
	o.PendingDomainUnverifiedAt.Unset()
}

// GetCustom404 returns the Custom404 field value
func (o *Page) GetCustom404() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Custom404
}

// GetCustom404Ok returns a tuple with the Custom404 field value
// and a boolean to check if the value has been set.
func (o *Page) GetCustom404Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Custom404, true
}

// SetCustom404 sets field value
func (o *Page) SetCustom404(v bool) {
	o.Custom404 = v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise.
func (o *Page) GetHtmlUrl() string {
	if o == nil || o.HtmlUrl == nil {
		var ret string
		return ret
	}
	return *o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetHtmlUrlOk() (*string, bool) {
	if o == nil || o.HtmlUrl == nil {
		return nil, false
	}
	return o.HtmlUrl, true
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *Page) HasHtmlUrl() bool {
	if o != nil && o.HtmlUrl != nil {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given string and assigns it to the HtmlUrl field.
func (o *Page) SetHtmlUrl(v string) {
	o.HtmlUrl = &v
}

// GetBuildType returns the BuildType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Page) GetBuildType() string {
	if o == nil || o.BuildType.Get() == nil {
		var ret string
		return ret
	}
	return *o.BuildType.Get()
}

// GetBuildTypeOk returns a tuple with the BuildType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Page) GetBuildTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildType.Get(), o.BuildType.IsSet()
}

// HasBuildType returns a boolean if a field has been set.
func (o *Page) HasBuildType() bool {
	if o != nil && o.BuildType.IsSet() {
		return true
	}

	return false
}

// SetBuildType gets a reference to the given NullableString and assigns it to the BuildType field.
func (o *Page) SetBuildType(v string) {
	o.BuildType.Set(&v)
}
// SetBuildTypeNil sets the value for BuildType to be an explicit nil
func (o *Page) SetBuildTypeNil() {
	o.BuildType.Set(nil)
}

// UnsetBuildType ensures that no value is present for BuildType, not even an explicit nil
func (o *Page) UnsetBuildType() {
	o.BuildType.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Page) GetSource() PagesSourceHash {
	if o == nil || o.Source == nil {
		var ret PagesSourceHash
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetSourceOk() (*PagesSourceHash, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Page) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given PagesSourceHash and assigns it to the Source field.
func (o *Page) SetSource(v PagesSourceHash) {
	o.Source = &v
}

// GetPublic returns the Public field value
func (o *Page) GetPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *Page) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Public, true
}

// SetPublic sets field value
func (o *Page) SetPublic(v bool) {
	o.Public = v
}

// GetHttpsCertificate returns the HttpsCertificate field value if set, zero value otherwise.
func (o *Page) GetHttpsCertificate() PagesHttpsCertificate {
	if o == nil || o.HttpsCertificate == nil {
		var ret PagesHttpsCertificate
		return ret
	}
	return *o.HttpsCertificate
}

// GetHttpsCertificateOk returns a tuple with the HttpsCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetHttpsCertificateOk() (*PagesHttpsCertificate, bool) {
	if o == nil || o.HttpsCertificate == nil {
		return nil, false
	}
	return o.HttpsCertificate, true
}

// HasHttpsCertificate returns a boolean if a field has been set.
func (o *Page) HasHttpsCertificate() bool {
	if o != nil && o.HttpsCertificate != nil {
		return true
	}

	return false
}

// SetHttpsCertificate gets a reference to the given PagesHttpsCertificate and assigns it to the HttpsCertificate field.
func (o *Page) SetHttpsCertificate(v PagesHttpsCertificate) {
	o.HttpsCertificate = &v
}

// GetHttpsEnforced returns the HttpsEnforced field value if set, zero value otherwise.
func (o *Page) GetHttpsEnforced() bool {
	if o == nil || o.HttpsEnforced == nil {
		var ret bool
		return ret
	}
	return *o.HttpsEnforced
}

// GetHttpsEnforcedOk returns a tuple with the HttpsEnforced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetHttpsEnforcedOk() (*bool, bool) {
	if o == nil || o.HttpsEnforced == nil {
		return nil, false
	}
	return o.HttpsEnforced, true
}

// HasHttpsEnforced returns a boolean if a field has been set.
func (o *Page) HasHttpsEnforced() bool {
	if o != nil && o.HttpsEnforced != nil {
		return true
	}

	return false
}

// SetHttpsEnforced gets a reference to the given bool and assigns it to the HttpsEnforced field.
func (o *Page) SetHttpsEnforced(v bool) {
	o.HttpsEnforced = &v
}

func (o Page) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["status"] = o.Status.Get()
	}
	if true {
		toSerialize["cname"] = o.Cname.Get()
	}
	if o.ProtectedDomainState.IsSet() {
		toSerialize["protected_domain_state"] = o.ProtectedDomainState.Get()
	}
	if o.PendingDomainUnverifiedAt.IsSet() {
		toSerialize["pending_domain_unverified_at"] = o.PendingDomainUnverifiedAt.Get()
	}
	if true {
		toSerialize["custom_404"] = o.Custom404
	}
	if o.HtmlUrl != nil {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if o.BuildType.IsSet() {
		toSerialize["build_type"] = o.BuildType.Get()
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["public"] = o.Public
	}
	if o.HttpsCertificate != nil {
		toSerialize["https_certificate"] = o.HttpsCertificate
	}
	if o.HttpsEnforced != nil {
		toSerialize["https_enforced"] = o.HttpsEnforced
	}
	return json.Marshal(toSerialize)
}

type NullablePage struct {
	value *Page
	isSet bool
}

func (v NullablePage) Get() *Page {
	return v.value
}

func (v *NullablePage) Set(val *Page) {
	v.value = val
	v.isSet = true
}

func (v NullablePage) IsSet() bool {
	return v.isSet
}

func (v *NullablePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePage(val *Page) *NullablePage {
	return &NullablePage{value: val, isSet: true}
}

func (v NullablePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


