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

// ReactionsCreateForIssueRequest struct for ReactionsCreateForIssueRequest
type ReactionsCreateForIssueRequest struct {
	// The [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types) to add to the issue.
	Content string `json:"content"`
}

// NewReactionsCreateForIssueRequest instantiates a new ReactionsCreateForIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionsCreateForIssueRequest(content string) *ReactionsCreateForIssueRequest {
	this := ReactionsCreateForIssueRequest{}
	this.Content = content
	return &this
}

// NewReactionsCreateForIssueRequestWithDefaults instantiates a new ReactionsCreateForIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionsCreateForIssueRequestWithDefaults() *ReactionsCreateForIssueRequest {
	this := ReactionsCreateForIssueRequest{}
	return &this
}

// GetContent returns the Content field value
func (o *ReactionsCreateForIssueRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ReactionsCreateForIssueRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ReactionsCreateForIssueRequest) SetContent(v string) {
	o.Content = v
}

func (o ReactionsCreateForIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableReactionsCreateForIssueRequest struct {
	value *ReactionsCreateForIssueRequest
	isSet bool
}

func (v NullableReactionsCreateForIssueRequest) Get() *ReactionsCreateForIssueRequest {
	return v.value
}

func (v *NullableReactionsCreateForIssueRequest) Set(val *ReactionsCreateForIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionsCreateForIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionsCreateForIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionsCreateForIssueRequest(val *ReactionsCreateForIssueRequest) *NullableReactionsCreateForIssueRequest {
	return &NullableReactionsCreateForIssueRequest{value: val, isSet: true}
}

func (v NullableReactionsCreateForIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionsCreateForIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


