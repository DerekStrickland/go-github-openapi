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

// ScimUserListEnterprise struct for ScimUserListEnterprise
type ScimUserListEnterprise struct {
	Schemas []string `json:"schemas"`
	TotalResults float32 `json:"totalResults"`
	ItemsPerPage float32 `json:"itemsPerPage"`
	StartIndex float32 `json:"startIndex"`
	Resources []ScimUserListEnterpriseResourcesInner `json:"Resources"`
}

// NewScimUserListEnterprise instantiates a new ScimUserListEnterprise object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimUserListEnterprise(schemas []string, totalResults float32, itemsPerPage float32, startIndex float32, resources []ScimUserListEnterpriseResourcesInner) *ScimUserListEnterprise {
	this := ScimUserListEnterprise{}
	this.Schemas = schemas
	this.TotalResults = totalResults
	this.ItemsPerPage = itemsPerPage
	this.StartIndex = startIndex
	this.Resources = resources
	return &this
}

// NewScimUserListEnterpriseWithDefaults instantiates a new ScimUserListEnterprise object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimUserListEnterpriseWithDefaults() *ScimUserListEnterprise {
	this := ScimUserListEnterprise{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ScimUserListEnterprise) GetSchemas() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ScimUserListEnterprise) GetSchemasOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ScimUserListEnterprise) SetSchemas(v []string) {
	o.Schemas = v
}

// GetTotalResults returns the TotalResults field value
func (o *ScimUserListEnterprise) GetTotalResults() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value
// and a boolean to check if the value has been set.
func (o *ScimUserListEnterprise) GetTotalResultsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalResults, true
}

// SetTotalResults sets field value
func (o *ScimUserListEnterprise) SetTotalResults(v float32) {
	o.TotalResults = v
}

// GetItemsPerPage returns the ItemsPerPage field value
func (o *ScimUserListEnterprise) GetItemsPerPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value
// and a boolean to check if the value has been set.
func (o *ScimUserListEnterprise) GetItemsPerPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsPerPage, true
}

// SetItemsPerPage sets field value
func (o *ScimUserListEnterprise) SetItemsPerPage(v float32) {
	o.ItemsPerPage = v
}

// GetStartIndex returns the StartIndex field value
func (o *ScimUserListEnterprise) GetStartIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value
// and a boolean to check if the value has been set.
func (o *ScimUserListEnterprise) GetStartIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartIndex, true
}

// SetStartIndex sets field value
func (o *ScimUserListEnterprise) SetStartIndex(v float32) {
	o.StartIndex = v
}

// GetResources returns the Resources field value
func (o *ScimUserListEnterprise) GetResources() []ScimUserListEnterpriseResourcesInner {
	if o == nil {
		var ret []ScimUserListEnterpriseResourcesInner
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *ScimUserListEnterprise) GetResourcesOk() ([]ScimUserListEnterpriseResourcesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *ScimUserListEnterprise) SetResources(v []ScimUserListEnterpriseResourcesInner) {
	o.Resources = v
}

func (o ScimUserListEnterprise) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["totalResults"] = o.TotalResults
	}
	if true {
		toSerialize["itemsPerPage"] = o.ItemsPerPage
	}
	if true {
		toSerialize["startIndex"] = o.StartIndex
	}
	if true {
		toSerialize["Resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableScimUserListEnterprise struct {
	value *ScimUserListEnterprise
	isSet bool
}

func (v NullableScimUserListEnterprise) Get() *ScimUserListEnterprise {
	return v.value
}

func (v *NullableScimUserListEnterprise) Set(val *ScimUserListEnterprise) {
	v.value = val
	v.isSet = true
}

func (v NullableScimUserListEnterprise) IsSet() bool {
	return v.isSet
}

func (v *NullableScimUserListEnterprise) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimUserListEnterprise(val *ScimUserListEnterprise) *NullableScimUserListEnterprise {
	return &NullableScimUserListEnterprise{value: val, isSet: true}
}

func (v NullableScimUserListEnterprise) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimUserListEnterprise) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


