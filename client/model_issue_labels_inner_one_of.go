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

// IssueLabelsInnerOneOf struct for IssueLabelsInnerOneOf
type IssueLabelsInnerOneOf struct {
	Id *int64 `json:"id,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	Url *string `json:"url,omitempty"`
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Color NullableString `json:"color,omitempty"`
	Default *bool `json:"default,omitempty"`
}

// NewIssueLabelsInnerOneOf instantiates a new IssueLabelsInnerOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueLabelsInnerOneOf() *IssueLabelsInnerOneOf {
	this := IssueLabelsInnerOneOf{}
	return &this
}

// NewIssueLabelsInnerOneOfWithDefaults instantiates a new IssueLabelsInnerOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueLabelsInnerOneOfWithDefaults() *IssueLabelsInnerOneOf {
	this := IssueLabelsInnerOneOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IssueLabelsInnerOneOf) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueLabelsInnerOneOf) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IssueLabelsInnerOneOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IssueLabelsInnerOneOf) SetId(v int64) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *IssueLabelsInnerOneOf) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueLabelsInnerOneOf) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *IssueLabelsInnerOneOf) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *IssueLabelsInnerOneOf) SetNodeId(v string) {
	o.NodeId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *IssueLabelsInnerOneOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueLabelsInnerOneOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *IssueLabelsInnerOneOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *IssueLabelsInnerOneOf) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IssueLabelsInnerOneOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueLabelsInnerOneOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IssueLabelsInnerOneOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IssueLabelsInnerOneOf) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueLabelsInnerOneOf) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueLabelsInnerOneOf) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IssueLabelsInnerOneOf) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *IssueLabelsInnerOneOf) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *IssueLabelsInnerOneOf) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *IssueLabelsInnerOneOf) UnsetDescription() {
	o.Description.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueLabelsInnerOneOf) GetColor() string {
	if o == nil || o.Color.Get() == nil {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueLabelsInnerOneOf) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *IssueLabelsInnerOneOf) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *IssueLabelsInnerOneOf) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *IssueLabelsInnerOneOf) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *IssueLabelsInnerOneOf) UnsetColor() {
	o.Color.Unset()
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *IssueLabelsInnerOneOf) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueLabelsInnerOneOf) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *IssueLabelsInnerOneOf) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *IssueLabelsInnerOneOf) SetDefault(v bool) {
	o.Default = &v
}

func (o IssueLabelsInnerOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NodeId != nil {
		toSerialize["node_id"] = o.NodeId
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	return json.Marshal(toSerialize)
}

type NullableIssueLabelsInnerOneOf struct {
	value *IssueLabelsInnerOneOf
	isSet bool
}

func (v NullableIssueLabelsInnerOneOf) Get() *IssueLabelsInnerOneOf {
	return v.value
}

func (v *NullableIssueLabelsInnerOneOf) Set(val *IssueLabelsInnerOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueLabelsInnerOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueLabelsInnerOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueLabelsInnerOneOf(val *IssueLabelsInnerOneOf) *NullableIssueLabelsInnerOneOf {
	return &NullableIssueLabelsInnerOneOf{value: val, isSet: true}
}

func (v NullableIssueLabelsInnerOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueLabelsInnerOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


