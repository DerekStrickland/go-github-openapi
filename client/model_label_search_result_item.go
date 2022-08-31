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

// LabelSearchResultItem Label Search Result Item
type LabelSearchResultItem struct {
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Color string `json:"color"`
	Default bool `json:"default"`
	Description NullableString `json:"description"`
	Score float32 `json:"score"`
	TextMatches []SearchResultTextMatchesInner `json:"text_matches,omitempty"`
}

// NewLabelSearchResultItem instantiates a new LabelSearchResultItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelSearchResultItem(id int32, nodeId string, url string, name string, color string, default_ bool, description NullableString, score float32) *LabelSearchResultItem {
	this := LabelSearchResultItem{}
	this.Id = id
	this.NodeId = nodeId
	this.Url = url
	this.Name = name
	this.Color = color
	this.Default = default_
	this.Description = description
	this.Score = score
	return &this
}

// NewLabelSearchResultItemWithDefaults instantiates a new LabelSearchResultItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelSearchResultItemWithDefaults() *LabelSearchResultItem {
	this := LabelSearchResultItem{}
	return &this
}

// GetId returns the Id field value
func (o *LabelSearchResultItem) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LabelSearchResultItem) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LabelSearchResultItem) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *LabelSearchResultItem) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *LabelSearchResultItem) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *LabelSearchResultItem) SetNodeId(v string) {
	o.NodeId = v
}

// GetUrl returns the Url field value
func (o *LabelSearchResultItem) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *LabelSearchResultItem) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *LabelSearchResultItem) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *LabelSearchResultItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LabelSearchResultItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LabelSearchResultItem) SetName(v string) {
	o.Name = v
}

// GetColor returns the Color field value
func (o *LabelSearchResultItem) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *LabelSearchResultItem) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *LabelSearchResultItem) SetColor(v string) {
	o.Color = v
}

// GetDefault returns the Default field value
func (o *LabelSearchResultItem) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *LabelSearchResultItem) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *LabelSearchResultItem) SetDefault(v bool) {
	o.Default = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LabelSearchResultItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabelSearchResultItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *LabelSearchResultItem) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetScore returns the Score field value
func (o *LabelSearchResultItem) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *LabelSearchResultItem) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *LabelSearchResultItem) SetScore(v float32) {
	o.Score = v
}

// GetTextMatches returns the TextMatches field value if set, zero value otherwise.
func (o *LabelSearchResultItem) GetTextMatches() []SearchResultTextMatchesInner {
	if o == nil || o.TextMatches == nil {
		var ret []SearchResultTextMatchesInner
		return ret
	}
	return o.TextMatches
}

// GetTextMatchesOk returns a tuple with the TextMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelSearchResultItem) GetTextMatchesOk() ([]SearchResultTextMatchesInner, bool) {
	if o == nil || o.TextMatches == nil {
		return nil, false
	}
	return o.TextMatches, true
}

// HasTextMatches returns a boolean if a field has been set.
func (o *LabelSearchResultItem) HasTextMatches() bool {
	if o != nil && o.TextMatches != nil {
		return true
	}

	return false
}

// SetTextMatches gets a reference to the given []SearchResultTextMatchesInner and assigns it to the TextMatches field.
func (o *LabelSearchResultItem) SetTextMatches(v []SearchResultTextMatchesInner) {
	o.TextMatches = v
}

func (o LabelSearchResultItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["color"] = o.Color
	}
	if true {
		toSerialize["default"] = o.Default
	}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["score"] = o.Score
	}
	if o.TextMatches != nil {
		toSerialize["text_matches"] = o.TextMatches
	}
	return json.Marshal(toSerialize)
}

type NullableLabelSearchResultItem struct {
	value *LabelSearchResultItem
	isSet bool
}

func (v NullableLabelSearchResultItem) Get() *LabelSearchResultItem {
	return v.value
}

func (v *NullableLabelSearchResultItem) Set(val *LabelSearchResultItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelSearchResultItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelSearchResultItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelSearchResultItem(val *LabelSearchResultItem) *NullableLabelSearchResultItem {
	return &NullableLabelSearchResultItem{value: val, isSet: true}
}

func (v NullableLabelSearchResultItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelSearchResultItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


