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

// ProtectedBranchRequiredLinearHistory struct for ProtectedBranchRequiredLinearHistory
type ProtectedBranchRequiredLinearHistory struct {
	Enabled bool `json:"enabled"`
}

// NewProtectedBranchRequiredLinearHistory instantiates a new ProtectedBranchRequiredLinearHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectedBranchRequiredLinearHistory(enabled bool) *ProtectedBranchRequiredLinearHistory {
	this := ProtectedBranchRequiredLinearHistory{}
	this.Enabled = enabled
	return &this
}

// NewProtectedBranchRequiredLinearHistoryWithDefaults instantiates a new ProtectedBranchRequiredLinearHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectedBranchRequiredLinearHistoryWithDefaults() *ProtectedBranchRequiredLinearHistory {
	this := ProtectedBranchRequiredLinearHistory{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ProtectedBranchRequiredLinearHistory) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ProtectedBranchRequiredLinearHistory) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ProtectedBranchRequiredLinearHistory) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ProtectedBranchRequiredLinearHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableProtectedBranchRequiredLinearHistory struct {
	value *ProtectedBranchRequiredLinearHistory
	isSet bool
}

func (v NullableProtectedBranchRequiredLinearHistory) Get() *ProtectedBranchRequiredLinearHistory {
	return v.value
}

func (v *NullableProtectedBranchRequiredLinearHistory) Set(val *ProtectedBranchRequiredLinearHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectedBranchRequiredLinearHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectedBranchRequiredLinearHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectedBranchRequiredLinearHistory(val *ProtectedBranchRequiredLinearHistory) *NullableProtectedBranchRequiredLinearHistory {
	return &NullableProtectedBranchRequiredLinearHistory{value: val, isSet: true}
}

func (v NullableProtectedBranchRequiredLinearHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectedBranchRequiredLinearHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


