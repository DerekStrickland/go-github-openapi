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

// CheckRunOutput struct for CheckRunOutput
type CheckRunOutput struct {
	Title NullableString `json:"title"`
	Summary NullableString `json:"summary"`
	Text NullableString `json:"text"`
	AnnotationsCount int32 `json:"annotations_count"`
	AnnotationsUrl string `json:"annotations_url"`
}

// NewCheckRunOutput instantiates a new CheckRunOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckRunOutput(title NullableString, summary NullableString, text NullableString, annotationsCount int32, annotationsUrl string) *CheckRunOutput {
	this := CheckRunOutput{}
	this.Title = title
	this.Summary = summary
	this.Text = text
	this.AnnotationsCount = annotationsCount
	this.AnnotationsUrl = annotationsUrl
	return &this
}

// NewCheckRunOutputWithDefaults instantiates a new CheckRunOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckRunOutputWithDefaults() *CheckRunOutput {
	this := CheckRunOutput{}
	return &this
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckRunOutput) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckRunOutput) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CheckRunOutput) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetSummary returns the Summary field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckRunOutput) GetSummary() string {
	if o == nil || o.Summary.Get() == nil {
		var ret string
		return ret
	}

	return *o.Summary.Get()
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckRunOutput) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary.Get(), o.Summary.IsSet()
}

// SetSummary sets field value
func (o *CheckRunOutput) SetSummary(v string) {
	o.Summary.Set(&v)
}

// GetText returns the Text field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckRunOutput) GetText() string {
	if o == nil || o.Text.Get() == nil {
		var ret string
		return ret
	}

	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckRunOutput) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// SetText sets field value
func (o *CheckRunOutput) SetText(v string) {
	o.Text.Set(&v)
}

// GetAnnotationsCount returns the AnnotationsCount field value
func (o *CheckRunOutput) GetAnnotationsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AnnotationsCount
}

// GetAnnotationsCountOk returns a tuple with the AnnotationsCount field value
// and a boolean to check if the value has been set.
func (o *CheckRunOutput) GetAnnotationsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnnotationsCount, true
}

// SetAnnotationsCount sets field value
func (o *CheckRunOutput) SetAnnotationsCount(v int32) {
	o.AnnotationsCount = v
}

// GetAnnotationsUrl returns the AnnotationsUrl field value
func (o *CheckRunOutput) GetAnnotationsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnnotationsUrl
}

// GetAnnotationsUrlOk returns a tuple with the AnnotationsUrl field value
// and a boolean to check if the value has been set.
func (o *CheckRunOutput) GetAnnotationsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnnotationsUrl, true
}

// SetAnnotationsUrl sets field value
func (o *CheckRunOutput) SetAnnotationsUrl(v string) {
	o.AnnotationsUrl = v
}

func (o CheckRunOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title.Get()
	}
	if true {
		toSerialize["summary"] = o.Summary.Get()
	}
	if true {
		toSerialize["text"] = o.Text.Get()
	}
	if true {
		toSerialize["annotations_count"] = o.AnnotationsCount
	}
	if true {
		toSerialize["annotations_url"] = o.AnnotationsUrl
	}
	return json.Marshal(toSerialize)
}

type NullableCheckRunOutput struct {
	value *CheckRunOutput
	isSet bool
}

func (v NullableCheckRunOutput) Get() *CheckRunOutput {
	return v.value
}

func (v *NullableCheckRunOutput) Set(val *CheckRunOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckRunOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckRunOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckRunOutput(val *CheckRunOutput) *NullableCheckRunOutput {
	return &NullableCheckRunOutput{value: val, isSet: true}
}

func (v NullableCheckRunOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckRunOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

