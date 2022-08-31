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

// PullsUpdateBranchRequest struct for PullsUpdateBranchRequest
type PullsUpdateBranchRequest struct {
	// The expected SHA of the pull request's HEAD ref. This is the most recent commit on the pull request's branch. If the expected SHA does not match the pull request's HEAD, you will receive a `422 Unprocessable Entity` status. You can use the \"[List commits](https://docs.github.com/rest/reference/repos#list-commits)\" endpoint to find the most recent commit SHA. Default: SHA of the pull request's current HEAD ref.
	ExpectedHeadSha *string `json:"expected_head_sha,omitempty"`
}

// NewPullsUpdateBranchRequest instantiates a new PullsUpdateBranchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullsUpdateBranchRequest() *PullsUpdateBranchRequest {
	this := PullsUpdateBranchRequest{}
	return &this
}

// NewPullsUpdateBranchRequestWithDefaults instantiates a new PullsUpdateBranchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullsUpdateBranchRequestWithDefaults() *PullsUpdateBranchRequest {
	this := PullsUpdateBranchRequest{}
	return &this
}

// GetExpectedHeadSha returns the ExpectedHeadSha field value if set, zero value otherwise.
func (o *PullsUpdateBranchRequest) GetExpectedHeadSha() string {
	if o == nil || o.ExpectedHeadSha == nil {
		var ret string
		return ret
	}
	return *o.ExpectedHeadSha
}

// GetExpectedHeadShaOk returns a tuple with the ExpectedHeadSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullsUpdateBranchRequest) GetExpectedHeadShaOk() (*string, bool) {
	if o == nil || o.ExpectedHeadSha == nil {
		return nil, false
	}
	return o.ExpectedHeadSha, true
}

// HasExpectedHeadSha returns a boolean if a field has been set.
func (o *PullsUpdateBranchRequest) HasExpectedHeadSha() bool {
	if o != nil && o.ExpectedHeadSha != nil {
		return true
	}

	return false
}

// SetExpectedHeadSha gets a reference to the given string and assigns it to the ExpectedHeadSha field.
func (o *PullsUpdateBranchRequest) SetExpectedHeadSha(v string) {
	o.ExpectedHeadSha = &v
}

func (o PullsUpdateBranchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpectedHeadSha != nil {
		toSerialize["expected_head_sha"] = o.ExpectedHeadSha
	}
	return json.Marshal(toSerialize)
}

type NullablePullsUpdateBranchRequest struct {
	value *PullsUpdateBranchRequest
	isSet bool
}

func (v NullablePullsUpdateBranchRequest) Get() *PullsUpdateBranchRequest {
	return v.value
}

func (v *NullablePullsUpdateBranchRequest) Set(val *PullsUpdateBranchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePullsUpdateBranchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePullsUpdateBranchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullsUpdateBranchRequest(val *PullsUpdateBranchRequest) *NullablePullsUpdateBranchRequest {
	return &NullablePullsUpdateBranchRequest{value: val, isSet: true}
}

func (v NullablePullsUpdateBranchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullsUpdateBranchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

