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

// ShortBranchCommit struct for ShortBranchCommit
type ShortBranchCommit struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// NewShortBranchCommit instantiates a new ShortBranchCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShortBranchCommit(sha string, url string) *ShortBranchCommit {
	this := ShortBranchCommit{}
	this.Sha = sha
	this.Url = url
	return &this
}

// NewShortBranchCommitWithDefaults instantiates a new ShortBranchCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShortBranchCommitWithDefaults() *ShortBranchCommit {
	this := ShortBranchCommit{}
	return &this
}

// GetSha returns the Sha field value
func (o *ShortBranchCommit) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *ShortBranchCommit) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *ShortBranchCommit) SetSha(v string) {
	o.Sha = v
}

// GetUrl returns the Url field value
func (o *ShortBranchCommit) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ShortBranchCommit) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ShortBranchCommit) SetUrl(v string) {
	o.Url = v
}

func (o ShortBranchCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sha"] = o.Sha
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableShortBranchCommit struct {
	value *ShortBranchCommit
	isSet bool
}

func (v NullableShortBranchCommit) Get() *ShortBranchCommit {
	return v.value
}

func (v *NullableShortBranchCommit) Set(val *ShortBranchCommit) {
	v.value = val
	v.isSet = true
}

func (v NullableShortBranchCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableShortBranchCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShortBranchCommit(val *ShortBranchCommit) *NullableShortBranchCommit {
	return &NullableShortBranchCommit{value: val, isSet: true}
}

func (v NullableShortBranchCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShortBranchCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


