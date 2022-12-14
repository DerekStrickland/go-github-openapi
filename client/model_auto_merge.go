/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
)

// AutoMerge The status of auto merging a pull request.
type AutoMerge struct {
	EnabledBy SimpleUser `json:"enabled_by"`
	// The merge method to use.
	MergeMethod string `json:"merge_method"`
	// Title for the merge commit message.
	CommitTitle string `json:"commit_title"`
	// Commit message for the merge commit.
	CommitMessage string `json:"commit_message"`
}

// NewAutoMerge instantiates a new AutoMerge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoMerge(enabledBy SimpleUser, mergeMethod string, commitTitle string, commitMessage string) *AutoMerge {
	this := AutoMerge{}
	this.EnabledBy = enabledBy
	this.MergeMethod = mergeMethod
	this.CommitTitle = commitTitle
	this.CommitMessage = commitMessage
	return &this
}

// NewAutoMergeWithDefaults instantiates a new AutoMerge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoMergeWithDefaults() *AutoMerge {
	this := AutoMerge{}
	return &this
}

// GetEnabledBy returns the EnabledBy field value
func (o *AutoMerge) GetEnabledBy() SimpleUser {
	if o == nil {
		var ret SimpleUser
		return ret
	}

	return o.EnabledBy
}

// GetEnabledByOk returns a tuple with the EnabledBy field value
// and a boolean to check if the value has been set.
func (o *AutoMerge) GetEnabledByOk() (*SimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnabledBy, true
}

// SetEnabledBy sets field value
func (o *AutoMerge) SetEnabledBy(v SimpleUser) {
	o.EnabledBy = v
}

// GetMergeMethod returns the MergeMethod field value
func (o *AutoMerge) GetMergeMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MergeMethod
}

// GetMergeMethodOk returns a tuple with the MergeMethod field value
// and a boolean to check if the value has been set.
func (o *AutoMerge) GetMergeMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MergeMethod, true
}

// SetMergeMethod sets field value
func (o *AutoMerge) SetMergeMethod(v string) {
	o.MergeMethod = v
}

// GetCommitTitle returns the CommitTitle field value
func (o *AutoMerge) GetCommitTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitTitle
}

// GetCommitTitleOk returns a tuple with the CommitTitle field value
// and a boolean to check if the value has been set.
func (o *AutoMerge) GetCommitTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitTitle, true
}

// SetCommitTitle sets field value
func (o *AutoMerge) SetCommitTitle(v string) {
	o.CommitTitle = v
}

// GetCommitMessage returns the CommitMessage field value
func (o *AutoMerge) GetCommitMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitMessage
}

// GetCommitMessageOk returns a tuple with the CommitMessage field value
// and a boolean to check if the value has been set.
func (o *AutoMerge) GetCommitMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitMessage, true
}

// SetCommitMessage sets field value
func (o *AutoMerge) SetCommitMessage(v string) {
	o.CommitMessage = v
}

func (o AutoMerge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled_by"] = o.EnabledBy
	}
	if true {
		toSerialize["merge_method"] = o.MergeMethod
	}
	if true {
		toSerialize["commit_title"] = o.CommitTitle
	}
	if true {
		toSerialize["commit_message"] = o.CommitMessage
	}
	return json.Marshal(toSerialize)
}

type NullableAutoMerge struct {
	value *AutoMerge
	isSet bool
}

func (v NullableAutoMerge) Get() *AutoMerge {
	return v.value
}

func (v *NullableAutoMerge) Set(val *AutoMerge) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoMerge) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoMerge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoMerge(val *AutoMerge) *NullableAutoMerge {
	return &NullableAutoMerge{value: val, isSet: true}
}

func (v NullableAutoMerge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoMerge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


