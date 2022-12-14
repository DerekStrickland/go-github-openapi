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

// PullRequestLabelsInner struct for PullRequestLabelsInner
type PullRequestLabelsInner struct {
	Id int64 `json:"id"`
	NodeId string `json:"node_id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
	Color string `json:"color"`
	Default bool `json:"default"`
}

// NewPullRequestLabelsInner instantiates a new PullRequestLabelsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequestLabelsInner(id int64, nodeId string, url string, name string, description NullableString, color string, default_ bool) *PullRequestLabelsInner {
	this := PullRequestLabelsInner{}
	this.Id = id
	this.NodeId = nodeId
	this.Url = url
	this.Name = name
	this.Description = description
	this.Color = color
	this.Default = default_
	return &this
}

// NewPullRequestLabelsInnerWithDefaults instantiates a new PullRequestLabelsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestLabelsInnerWithDefaults() *PullRequestLabelsInner {
	this := PullRequestLabelsInner{}
	return &this
}

// GetId returns the Id field value
func (o *PullRequestLabelsInner) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PullRequestLabelsInner) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PullRequestLabelsInner) SetId(v int64) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *PullRequestLabelsInner) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *PullRequestLabelsInner) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *PullRequestLabelsInner) SetNodeId(v string) {
	o.NodeId = v
}

// GetUrl returns the Url field value
func (o *PullRequestLabelsInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PullRequestLabelsInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PullRequestLabelsInner) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *PullRequestLabelsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PullRequestLabelsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PullRequestLabelsInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PullRequestLabelsInner) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestLabelsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *PullRequestLabelsInner) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetColor returns the Color field value
func (o *PullRequestLabelsInner) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *PullRequestLabelsInner) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *PullRequestLabelsInner) SetColor(v string) {
	o.Color = v
}

// GetDefault returns the Default field value
func (o *PullRequestLabelsInner) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *PullRequestLabelsInner) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *PullRequestLabelsInner) SetDefault(v bool) {
	o.Default = v
}

func (o PullRequestLabelsInner) MarshalJSON() ([]byte, error) {
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
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["color"] = o.Color
	}
	if true {
		toSerialize["default"] = o.Default
	}
	return json.Marshal(toSerialize)
}

type NullablePullRequestLabelsInner struct {
	value *PullRequestLabelsInner
	isSet bool
}

func (v NullablePullRequestLabelsInner) Get() *PullRequestLabelsInner {
	return v.value
}

func (v *NullablePullRequestLabelsInner) Set(val *PullRequestLabelsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequestLabelsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequestLabelsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequestLabelsInner(val *PullRequestLabelsInner) *NullablePullRequestLabelsInner {
	return &NullablePullRequestLabelsInner{value: val, isSet: true}
}

func (v NullablePullRequestLabelsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequestLabelsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


