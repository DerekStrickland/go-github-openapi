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

// ActivitySetRepoSubscriptionRequest struct for ActivitySetRepoSubscriptionRequest
type ActivitySetRepoSubscriptionRequest struct {
	// Determines if notifications should be received from this repository.
	Subscribed *bool `json:"subscribed,omitempty"`
	// Determines if all notifications should be blocked from this repository.
	Ignored *bool `json:"ignored,omitempty"`
}

// NewActivitySetRepoSubscriptionRequest instantiates a new ActivitySetRepoSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivitySetRepoSubscriptionRequest() *ActivitySetRepoSubscriptionRequest {
	this := ActivitySetRepoSubscriptionRequest{}
	return &this
}

// NewActivitySetRepoSubscriptionRequestWithDefaults instantiates a new ActivitySetRepoSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivitySetRepoSubscriptionRequestWithDefaults() *ActivitySetRepoSubscriptionRequest {
	this := ActivitySetRepoSubscriptionRequest{}
	return &this
}

// GetSubscribed returns the Subscribed field value if set, zero value otherwise.
func (o *ActivitySetRepoSubscriptionRequest) GetSubscribed() bool {
	if o == nil || o.Subscribed == nil {
		var ret bool
		return ret
	}
	return *o.Subscribed
}

// GetSubscribedOk returns a tuple with the Subscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivitySetRepoSubscriptionRequest) GetSubscribedOk() (*bool, bool) {
	if o == nil || o.Subscribed == nil {
		return nil, false
	}
	return o.Subscribed, true
}

// HasSubscribed returns a boolean if a field has been set.
func (o *ActivitySetRepoSubscriptionRequest) HasSubscribed() bool {
	if o != nil && o.Subscribed != nil {
		return true
	}

	return false
}

// SetSubscribed gets a reference to the given bool and assigns it to the Subscribed field.
func (o *ActivitySetRepoSubscriptionRequest) SetSubscribed(v bool) {
	o.Subscribed = &v
}

// GetIgnored returns the Ignored field value if set, zero value otherwise.
func (o *ActivitySetRepoSubscriptionRequest) GetIgnored() bool {
	if o == nil || o.Ignored == nil {
		var ret bool
		return ret
	}
	return *o.Ignored
}

// GetIgnoredOk returns a tuple with the Ignored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivitySetRepoSubscriptionRequest) GetIgnoredOk() (*bool, bool) {
	if o == nil || o.Ignored == nil {
		return nil, false
	}
	return o.Ignored, true
}

// HasIgnored returns a boolean if a field has been set.
func (o *ActivitySetRepoSubscriptionRequest) HasIgnored() bool {
	if o != nil && o.Ignored != nil {
		return true
	}

	return false
}

// SetIgnored gets a reference to the given bool and assigns it to the Ignored field.
func (o *ActivitySetRepoSubscriptionRequest) SetIgnored(v bool) {
	o.Ignored = &v
}

func (o ActivitySetRepoSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subscribed != nil {
		toSerialize["subscribed"] = o.Subscribed
	}
	if o.Ignored != nil {
		toSerialize["ignored"] = o.Ignored
	}
	return json.Marshal(toSerialize)
}

type NullableActivitySetRepoSubscriptionRequest struct {
	value *ActivitySetRepoSubscriptionRequest
	isSet bool
}

func (v NullableActivitySetRepoSubscriptionRequest) Get() *ActivitySetRepoSubscriptionRequest {
	return v.value
}

func (v *NullableActivitySetRepoSubscriptionRequest) Set(val *ActivitySetRepoSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActivitySetRepoSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActivitySetRepoSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivitySetRepoSubscriptionRequest(val *ActivitySetRepoSubscriptionRequest) *NullableActivitySetRepoSubscriptionRequest {
	return &NullableActivitySetRepoSubscriptionRequest{value: val, isSet: true}
}

func (v NullableActivitySetRepoSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivitySetRepoSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

