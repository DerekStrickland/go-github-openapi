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

// FileCommitCommitTree struct for FileCommitCommitTree
type FileCommitCommitTree struct {
	Url *string `json:"url,omitempty"`
	Sha *string `json:"sha,omitempty"`
}

// NewFileCommitCommitTree instantiates a new FileCommitCommitTree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileCommitCommitTree() *FileCommitCommitTree {
	this := FileCommitCommitTree{}
	return &this
}

// NewFileCommitCommitTreeWithDefaults instantiates a new FileCommitCommitTree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileCommitCommitTreeWithDefaults() *FileCommitCommitTree {
	this := FileCommitCommitTree{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *FileCommitCommitTree) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitCommitTree) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *FileCommitCommitTree) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *FileCommitCommitTree) SetUrl(v string) {
	o.Url = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *FileCommitCommitTree) GetSha() string {
	if o == nil || o.Sha == nil {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitCommitTree) GetShaOk() (*string, bool) {
	if o == nil || o.Sha == nil {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *FileCommitCommitTree) HasSha() bool {
	if o != nil && o.Sha != nil {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *FileCommitCommitTree) SetSha(v string) {
	o.Sha = &v
}

func (o FileCommitCommitTree) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Sha != nil {
		toSerialize["sha"] = o.Sha
	}
	return json.Marshal(toSerialize)
}

type NullableFileCommitCommitTree struct {
	value *FileCommitCommitTree
	isSet bool
}

func (v NullableFileCommitCommitTree) Get() *FileCommitCommitTree {
	return v.value
}

func (v *NullableFileCommitCommitTree) Set(val *FileCommitCommitTree) {
	v.value = val
	v.isSet = true
}

func (v NullableFileCommitCommitTree) IsSet() bool {
	return v.isSet
}

func (v *NullableFileCommitCommitTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileCommitCommitTree(val *FileCommitCommitTree) *NullableFileCommitCommitTree {
	return &NullableFileCommitCommitTree{value: val, isSet: true}
}

func (v NullableFileCommitCommitTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileCommitCommitTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


