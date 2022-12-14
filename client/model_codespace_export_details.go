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

// CodespaceExportDetails An export of a codespace. Also, latest export details for a codespace can be fetched with id = latest
type CodespaceExportDetails struct {
	// State of the latest export
	State NullableString `json:"state,omitempty"`
	// Completion time of the last export operation
	CompletedAt NullableTime `json:"completed_at,omitempty"`
	// Name of the exported branch
	Branch NullableString `json:"branch,omitempty"`
	// Git commit SHA of the exported branch
	Sha NullableString `json:"sha,omitempty"`
	// Id for the export details
	Id *string `json:"id,omitempty"`
	// Url for fetching export details
	ExportUrl *string `json:"export_url,omitempty"`
	// Web url for the exported branch
	HtmlUrl NullableString `json:"html_url,omitempty"`
}

// NewCodespaceExportDetails instantiates a new CodespaceExportDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodespaceExportDetails() *CodespaceExportDetails {
	this := CodespaceExportDetails{}
	return &this
}

// NewCodespaceExportDetailsWithDefaults instantiates a new CodespaceExportDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodespaceExportDetailsWithDefaults() *CodespaceExportDetails {
	this := CodespaceExportDetails{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodespaceExportDetails) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodespaceExportDetails) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *CodespaceExportDetails) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *CodespaceExportDetails) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *CodespaceExportDetails) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *CodespaceExportDetails) UnsetState() {
	o.State.Unset()
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodespaceExportDetails) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodespaceExportDetails) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *CodespaceExportDetails) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableTime and assigns it to the CompletedAt field.
func (o *CodespaceExportDetails) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *CodespaceExportDetails) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *CodespaceExportDetails) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetBranch returns the Branch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodespaceExportDetails) GetBranch() string {
	if o == nil || o.Branch.Get() == nil {
		var ret string
		return ret
	}
	return *o.Branch.Get()
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodespaceExportDetails) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Branch.Get(), o.Branch.IsSet()
}

// HasBranch returns a boolean if a field has been set.
func (o *CodespaceExportDetails) HasBranch() bool {
	if o != nil && o.Branch.IsSet() {
		return true
	}

	return false
}

// SetBranch gets a reference to the given NullableString and assigns it to the Branch field.
func (o *CodespaceExportDetails) SetBranch(v string) {
	o.Branch.Set(&v)
}
// SetBranchNil sets the value for Branch to be an explicit nil
func (o *CodespaceExportDetails) SetBranchNil() {
	o.Branch.Set(nil)
}

// UnsetBranch ensures that no value is present for Branch, not even an explicit nil
func (o *CodespaceExportDetails) UnsetBranch() {
	o.Branch.Unset()
}

// GetSha returns the Sha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodespaceExportDetails) GetSha() string {
	if o == nil || o.Sha.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sha.Get()
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodespaceExportDetails) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha.Get(), o.Sha.IsSet()
}

// HasSha returns a boolean if a field has been set.
func (o *CodespaceExportDetails) HasSha() bool {
	if o != nil && o.Sha.IsSet() {
		return true
	}

	return false
}

// SetSha gets a reference to the given NullableString and assigns it to the Sha field.
func (o *CodespaceExportDetails) SetSha(v string) {
	o.Sha.Set(&v)
}
// SetShaNil sets the value for Sha to be an explicit nil
func (o *CodespaceExportDetails) SetShaNil() {
	o.Sha.Set(nil)
}

// UnsetSha ensures that no value is present for Sha, not even an explicit nil
func (o *CodespaceExportDetails) UnsetSha() {
	o.Sha.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CodespaceExportDetails) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodespaceExportDetails) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CodespaceExportDetails) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CodespaceExportDetails) SetId(v string) {
	o.Id = &v
}

// GetExportUrl returns the ExportUrl field value if set, zero value otherwise.
func (o *CodespaceExportDetails) GetExportUrl() string {
	if o == nil || o.ExportUrl == nil {
		var ret string
		return ret
	}
	return *o.ExportUrl
}

// GetExportUrlOk returns a tuple with the ExportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodespaceExportDetails) GetExportUrlOk() (*string, bool) {
	if o == nil || o.ExportUrl == nil {
		return nil, false
	}
	return o.ExportUrl, true
}

// HasExportUrl returns a boolean if a field has been set.
func (o *CodespaceExportDetails) HasExportUrl() bool {
	if o != nil && o.ExportUrl != nil {
		return true
	}

	return false
}

// SetExportUrl gets a reference to the given string and assigns it to the ExportUrl field.
func (o *CodespaceExportDetails) SetExportUrl(v string) {
	o.ExportUrl = &v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodespaceExportDetails) GetHtmlUrl() string {
	if o == nil || o.HtmlUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.HtmlUrl.Get()
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodespaceExportDetails) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HtmlUrl.Get(), o.HtmlUrl.IsSet()
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *CodespaceExportDetails) HasHtmlUrl() bool {
	if o != nil && o.HtmlUrl.IsSet() {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given NullableString and assigns it to the HtmlUrl field.
func (o *CodespaceExportDetails) SetHtmlUrl(v string) {
	o.HtmlUrl.Set(&v)
}
// SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil
func (o *CodespaceExportDetails) SetHtmlUrlNil() {
	o.HtmlUrl.Set(nil)
}

// UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil
func (o *CodespaceExportDetails) UnsetHtmlUrl() {
	o.HtmlUrl.Unset()
}

func (o CodespaceExportDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if o.Branch.IsSet() {
		toSerialize["branch"] = o.Branch.Get()
	}
	if o.Sha.IsSet() {
		toSerialize["sha"] = o.Sha.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ExportUrl != nil {
		toSerialize["export_url"] = o.ExportUrl
	}
	if o.HtmlUrl.IsSet() {
		toSerialize["html_url"] = o.HtmlUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCodespaceExportDetails struct {
	value *CodespaceExportDetails
	isSet bool
}

func (v NullableCodespaceExportDetails) Get() *CodespaceExportDetails {
	return v.value
}

func (v *NullableCodespaceExportDetails) Set(val *CodespaceExportDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCodespaceExportDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCodespaceExportDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodespaceExportDetails(val *CodespaceExportDetails) *NullableCodespaceExportDetails {
	return &NullableCodespaceExportDetails{value: val, isSet: true}
}

func (v NullableCodespaceExportDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodespaceExportDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


