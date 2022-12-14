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

// GitCommitVerification struct for GitCommitVerification
type GitCommitVerification struct {
	Verified bool `json:"verified"`
	Reason string `json:"reason"`
	Signature NullableString `json:"signature"`
	Payload NullableString `json:"payload"`
}

// NewGitCommitVerification instantiates a new GitCommitVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitCommitVerification(verified bool, reason string, signature NullableString, payload NullableString) *GitCommitVerification {
	this := GitCommitVerification{}
	this.Verified = verified
	this.Reason = reason
	this.Signature = signature
	this.Payload = payload
	return &this
}

// NewGitCommitVerificationWithDefaults instantiates a new GitCommitVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitCommitVerificationWithDefaults() *GitCommitVerification {
	this := GitCommitVerification{}
	return &this
}

// GetVerified returns the Verified field value
func (o *GitCommitVerification) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *GitCommitVerification) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *GitCommitVerification) SetVerified(v bool) {
	o.Verified = v
}

// GetReason returns the Reason field value
func (o *GitCommitVerification) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *GitCommitVerification) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *GitCommitVerification) SetReason(v string) {
	o.Reason = v
}

// GetSignature returns the Signature field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GitCommitVerification) GetSignature() string {
	if o == nil || o.Signature.Get() == nil {
		var ret string
		return ret
	}

	return *o.Signature.Get()
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitCommitVerification) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signature.Get(), o.Signature.IsSet()
}

// SetSignature sets field value
func (o *GitCommitVerification) SetSignature(v string) {
	o.Signature.Set(&v)
}

// GetPayload returns the Payload field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GitCommitVerification) GetPayload() string {
	if o == nil || o.Payload.Get() == nil {
		var ret string
		return ret
	}

	return *o.Payload.Get()
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitCommitVerification) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payload.Get(), o.Payload.IsSet()
}

// SetPayload sets field value
func (o *GitCommitVerification) SetPayload(v string) {
	o.Payload.Set(&v)
}

func (o GitCommitVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["verified"] = o.Verified
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["signature"] = o.Signature.Get()
	}
	if true {
		toSerialize["payload"] = o.Payload.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGitCommitVerification struct {
	value *GitCommitVerification
	isSet bool
}

func (v NullableGitCommitVerification) Get() *GitCommitVerification {
	return v.value
}

func (v *NullableGitCommitVerification) Set(val *GitCommitVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableGitCommitVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableGitCommitVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitCommitVerification(val *GitCommitVerification) *NullableGitCommitVerification {
	return &NullableGitCommitVerification{value: val, isSet: true}
}

func (v NullableGitCommitVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitCommitVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


