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

// RateLimit struct for RateLimit
type RateLimit struct {
	Limit int32 `json:"limit"`
	Remaining int32 `json:"remaining"`
	Reset int32 `json:"reset"`
	Used int32 `json:"used"`
}

// NewRateLimit instantiates a new RateLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimit(limit int32, remaining int32, reset int32, used int32) *RateLimit {
	this := RateLimit{}
	this.Limit = limit
	this.Remaining = remaining
	this.Reset = reset
	this.Used = used
	return &this
}

// NewRateLimitWithDefaults instantiates a new RateLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitWithDefaults() *RateLimit {
	this := RateLimit{}
	return &this
}

// GetLimit returns the Limit field value
func (o *RateLimit) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *RateLimit) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *RateLimit) SetLimit(v int32) {
	o.Limit = v
}

// GetRemaining returns the Remaining field value
func (o *RateLimit) GetRemaining() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value
// and a boolean to check if the value has been set.
func (o *RateLimit) GetRemainingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remaining, true
}

// SetRemaining sets field value
func (o *RateLimit) SetRemaining(v int32) {
	o.Remaining = v
}

// GetReset returns the Reset field value
func (o *RateLimit) GetReset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Reset
}

// GetResetOk returns a tuple with the Reset field value
// and a boolean to check if the value has been set.
func (o *RateLimit) GetResetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reset, true
}

// SetReset sets field value
func (o *RateLimit) SetReset(v int32) {
	o.Reset = v
}

// GetUsed returns the Used field value
func (o *RateLimit) GetUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *RateLimit) GetUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *RateLimit) SetUsed(v int32) {
	o.Used = v
}

func (o RateLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["remaining"] = o.Remaining
	}
	if true {
		toSerialize["reset"] = o.Reset
	}
	if true {
		toSerialize["used"] = o.Used
	}
	return json.Marshal(toSerialize)
}

type NullableRateLimit struct {
	value *RateLimit
	isSet bool
}

func (v NullableRateLimit) Get() *RateLimit {
	return v.value
}

func (v *NullableRateLimit) Set(val *RateLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimit(val *RateLimit) *NullableRateLimit {
	return &NullableRateLimit{value: val, isSet: true}
}

func (v NullableRateLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


