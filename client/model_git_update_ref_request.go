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

// GitUpdateRefRequest struct for GitUpdateRefRequest
type GitUpdateRefRequest struct {
	// The SHA1 value to set this reference to
	Sha string `json:"sha"`
	// Indicates whether to force the update or to make sure the update is a fast-forward update. Leaving this out or setting it to `false` will make sure you're not overwriting work.
	Force *bool `json:"force,omitempty"`
}

// NewGitUpdateRefRequest instantiates a new GitUpdateRefRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitUpdateRefRequest(sha string) *GitUpdateRefRequest {
	this := GitUpdateRefRequest{}
	this.Sha = sha
	var force bool = false
	this.Force = &force
	return &this
}

// NewGitUpdateRefRequestWithDefaults instantiates a new GitUpdateRefRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitUpdateRefRequestWithDefaults() *GitUpdateRefRequest {
	this := GitUpdateRefRequest{}
	var force bool = false
	this.Force = &force
	return &this
}

// GetSha returns the Sha field value
func (o *GitUpdateRefRequest) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *GitUpdateRefRequest) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *GitUpdateRefRequest) SetSha(v string) {
	o.Sha = v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *GitUpdateRefRequest) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitUpdateRefRequest) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *GitUpdateRefRequest) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *GitUpdateRefRequest) SetForce(v bool) {
	o.Force = &v
}

func (o GitUpdateRefRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sha"] = o.Sha
	}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableGitUpdateRefRequest struct {
	value *GitUpdateRefRequest
	isSet bool
}

func (v NullableGitUpdateRefRequest) Get() *GitUpdateRefRequest {
	return v.value
}

func (v *NullableGitUpdateRefRequest) Set(val *GitUpdateRefRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGitUpdateRefRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGitUpdateRefRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitUpdateRefRequest(val *GitUpdateRefRequest) *NullableGitUpdateRefRequest {
	return &NullableGitUpdateRefRequest{value: val, isSet: true}
}

func (v NullableGitUpdateRefRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitUpdateRefRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

