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

// GitTagObject struct for GitTagObject
type GitTagObject struct {
	Sha string `json:"sha"`
	Type string `json:"type"`
	Url string `json:"url"`
}

// NewGitTagObject instantiates a new GitTagObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitTagObject(sha string, type_ string, url string) *GitTagObject {
	this := GitTagObject{}
	this.Sha = sha
	this.Type = type_
	this.Url = url
	return &this
}

// NewGitTagObjectWithDefaults instantiates a new GitTagObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitTagObjectWithDefaults() *GitTagObject {
	this := GitTagObject{}
	return &this
}

// GetSha returns the Sha field value
func (o *GitTagObject) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *GitTagObject) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *GitTagObject) SetSha(v string) {
	o.Sha = v
}

// GetType returns the Type field value
func (o *GitTagObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GitTagObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GitTagObject) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *GitTagObject) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GitTagObject) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GitTagObject) SetUrl(v string) {
	o.Url = v
}

func (o GitTagObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sha"] = o.Sha
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableGitTagObject struct {
	value *GitTagObject
	isSet bool
}

func (v NullableGitTagObject) Get() *GitTagObject {
	return v.value
}

func (v *NullableGitTagObject) Set(val *GitTagObject) {
	v.value = val
	v.isSet = true
}

func (v NullableGitTagObject) IsSet() bool {
	return v.isSet
}

func (v *NullableGitTagObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitTagObject(val *GitTagObject) *NullableGitTagObject {
	return &NullableGitTagObject{value: val, isSet: true}
}

func (v NullableGitTagObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitTagObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


