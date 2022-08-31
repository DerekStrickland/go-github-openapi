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

// WebhookMergeGroupChecksRequestedMergeGroup struct for WebhookMergeGroupChecksRequestedMergeGroup
type WebhookMergeGroupChecksRequestedMergeGroup struct {
	BaseRef string `json:"base_ref"`
	HeadRef string `json:"head_ref"`
	HeadSha string `json:"head_sha"`
}

// NewWebhookMergeGroupChecksRequestedMergeGroup instantiates a new WebhookMergeGroupChecksRequestedMergeGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookMergeGroupChecksRequestedMergeGroup(baseRef string, headRef string, headSha string) *WebhookMergeGroupChecksRequestedMergeGroup {
	this := WebhookMergeGroupChecksRequestedMergeGroup{}
	this.BaseRef = baseRef
	this.HeadRef = headRef
	this.HeadSha = headSha
	return &this
}

// NewWebhookMergeGroupChecksRequestedMergeGroupWithDefaults instantiates a new WebhookMergeGroupChecksRequestedMergeGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookMergeGroupChecksRequestedMergeGroupWithDefaults() *WebhookMergeGroupChecksRequestedMergeGroup {
	this := WebhookMergeGroupChecksRequestedMergeGroup{}
	return &this
}

// GetBaseRef returns the BaseRef field value
func (o *WebhookMergeGroupChecksRequestedMergeGroup) GetBaseRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseRef
}

// GetBaseRefOk returns a tuple with the BaseRef field value
// and a boolean to check if the value has been set.
func (o *WebhookMergeGroupChecksRequestedMergeGroup) GetBaseRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseRef, true
}

// SetBaseRef sets field value
func (o *WebhookMergeGroupChecksRequestedMergeGroup) SetBaseRef(v string) {
	o.BaseRef = v
}

// GetHeadRef returns the HeadRef field value
func (o *WebhookMergeGroupChecksRequestedMergeGroup) GetHeadRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeadRef
}

// GetHeadRefOk returns a tuple with the HeadRef field value
// and a boolean to check if the value has been set.
func (o *WebhookMergeGroupChecksRequestedMergeGroup) GetHeadRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeadRef, true
}

// SetHeadRef sets field value
func (o *WebhookMergeGroupChecksRequestedMergeGroup) SetHeadRef(v string) {
	o.HeadRef = v
}

// GetHeadSha returns the HeadSha field value
func (o *WebhookMergeGroupChecksRequestedMergeGroup) GetHeadSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeadSha
}

// GetHeadShaOk returns a tuple with the HeadSha field value
// and a boolean to check if the value has been set.
func (o *WebhookMergeGroupChecksRequestedMergeGroup) GetHeadShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeadSha, true
}

// SetHeadSha sets field value
func (o *WebhookMergeGroupChecksRequestedMergeGroup) SetHeadSha(v string) {
	o.HeadSha = v
}

func (o WebhookMergeGroupChecksRequestedMergeGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["base_ref"] = o.BaseRef
	}
	if true {
		toSerialize["head_ref"] = o.HeadRef
	}
	if true {
		toSerialize["head_sha"] = o.HeadSha
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookMergeGroupChecksRequestedMergeGroup struct {
	value *WebhookMergeGroupChecksRequestedMergeGroup
	isSet bool
}

func (v NullableWebhookMergeGroupChecksRequestedMergeGroup) Get() *WebhookMergeGroupChecksRequestedMergeGroup {
	return v.value
}

func (v *NullableWebhookMergeGroupChecksRequestedMergeGroup) Set(val *WebhookMergeGroupChecksRequestedMergeGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookMergeGroupChecksRequestedMergeGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookMergeGroupChecksRequestedMergeGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookMergeGroupChecksRequestedMergeGroup(val *WebhookMergeGroupChecksRequestedMergeGroup) *NullableWebhookMergeGroupChecksRequestedMergeGroup {
	return &NullableWebhookMergeGroupChecksRequestedMergeGroup{value: val, isSet: true}
}

func (v NullableWebhookMergeGroupChecksRequestedMergeGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookMergeGroupChecksRequestedMergeGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

