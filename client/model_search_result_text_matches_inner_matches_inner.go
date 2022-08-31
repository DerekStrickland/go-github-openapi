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

// SearchResultTextMatchesInnerMatchesInner struct for SearchResultTextMatchesInnerMatchesInner
type SearchResultTextMatchesInnerMatchesInner struct {
	Text *string `json:"text,omitempty"`
	Indices []int32 `json:"indices,omitempty"`
}

// NewSearchResultTextMatchesInnerMatchesInner instantiates a new SearchResultTextMatchesInnerMatchesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResultTextMatchesInnerMatchesInner() *SearchResultTextMatchesInnerMatchesInner {
	this := SearchResultTextMatchesInnerMatchesInner{}
	return &this
}

// NewSearchResultTextMatchesInnerMatchesInnerWithDefaults instantiates a new SearchResultTextMatchesInnerMatchesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultTextMatchesInnerMatchesInnerWithDefaults() *SearchResultTextMatchesInnerMatchesInner {
	this := SearchResultTextMatchesInnerMatchesInner{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *SearchResultTextMatchesInnerMatchesInner) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultTextMatchesInnerMatchesInner) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *SearchResultTextMatchesInnerMatchesInner) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *SearchResultTextMatchesInnerMatchesInner) SetText(v string) {
	o.Text = &v
}

// GetIndices returns the Indices field value if set, zero value otherwise.
func (o *SearchResultTextMatchesInnerMatchesInner) GetIndices() []int32 {
	if o == nil || o.Indices == nil {
		var ret []int32
		return ret
	}
	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultTextMatchesInnerMatchesInner) GetIndicesOk() ([]int32, bool) {
	if o == nil || o.Indices == nil {
		return nil, false
	}
	return o.Indices, true
}

// HasIndices returns a boolean if a field has been set.
func (o *SearchResultTextMatchesInnerMatchesInner) HasIndices() bool {
	if o != nil && o.Indices != nil {
		return true
	}

	return false
}

// SetIndices gets a reference to the given []int32 and assigns it to the Indices field.
func (o *SearchResultTextMatchesInnerMatchesInner) SetIndices(v []int32) {
	o.Indices = v
}

func (o SearchResultTextMatchesInnerMatchesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Indices != nil {
		toSerialize["indices"] = o.Indices
	}
	return json.Marshal(toSerialize)
}

type NullableSearchResultTextMatchesInnerMatchesInner struct {
	value *SearchResultTextMatchesInnerMatchesInner
	isSet bool
}

func (v NullableSearchResultTextMatchesInnerMatchesInner) Get() *SearchResultTextMatchesInnerMatchesInner {
	return v.value
}

func (v *NullableSearchResultTextMatchesInnerMatchesInner) Set(val *SearchResultTextMatchesInnerMatchesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResultTextMatchesInnerMatchesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResultTextMatchesInnerMatchesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResultTextMatchesInnerMatchesInner(val *SearchResultTextMatchesInnerMatchesInner) *NullableSearchResultTextMatchesInnerMatchesInner {
	return &NullableSearchResultTextMatchesInnerMatchesInner{value: val, isSet: true}
}

func (v NullableSearchResultTextMatchesInnerMatchesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResultTextMatchesInnerMatchesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

