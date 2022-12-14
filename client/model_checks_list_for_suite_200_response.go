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

// ChecksListForSuite200Response struct for ChecksListForSuite200Response
type ChecksListForSuite200Response struct {
	TotalCount int32 `json:"total_count"`
	CheckRuns []CheckRun `json:"check_runs"`
}

// NewChecksListForSuite200Response instantiates a new ChecksListForSuite200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChecksListForSuite200Response(totalCount int32, checkRuns []CheckRun) *ChecksListForSuite200Response {
	this := ChecksListForSuite200Response{}
	this.TotalCount = totalCount
	this.CheckRuns = checkRuns
	return &this
}

// NewChecksListForSuite200ResponseWithDefaults instantiates a new ChecksListForSuite200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChecksListForSuite200ResponseWithDefaults() *ChecksListForSuite200Response {
	this := ChecksListForSuite200Response{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *ChecksListForSuite200Response) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ChecksListForSuite200Response) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ChecksListForSuite200Response) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetCheckRuns returns the CheckRuns field value
func (o *ChecksListForSuite200Response) GetCheckRuns() []CheckRun {
	if o == nil {
		var ret []CheckRun
		return ret
	}

	return o.CheckRuns
}

// GetCheckRunsOk returns a tuple with the CheckRuns field value
// and a boolean to check if the value has been set.
func (o *ChecksListForSuite200Response) GetCheckRunsOk() ([]CheckRun, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckRuns, true
}

// SetCheckRuns sets field value
func (o *ChecksListForSuite200Response) SetCheckRuns(v []CheckRun) {
	o.CheckRuns = v
}

func (o ChecksListForSuite200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_count"] = o.TotalCount
	}
	if true {
		toSerialize["check_runs"] = o.CheckRuns
	}
	return json.Marshal(toSerialize)
}

type NullableChecksListForSuite200Response struct {
	value *ChecksListForSuite200Response
	isSet bool
}

func (v NullableChecksListForSuite200Response) Get() *ChecksListForSuite200Response {
	return v.value
}

func (v *NullableChecksListForSuite200Response) Set(val *ChecksListForSuite200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableChecksListForSuite200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableChecksListForSuite200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChecksListForSuite200Response(val *ChecksListForSuite200Response) *NullableChecksListForSuite200Response {
	return &NullableChecksListForSuite200Response{value: val, isSet: true}
}

func (v NullableChecksListForSuite200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChecksListForSuite200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


